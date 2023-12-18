package service

import (
	"context"
	"errors"
	dbmodel "known-anchors/dal/db/model"
	"known-anchors/model"
	"known-anchors/util/mail"
	"known-anchors/util/pwd"
	"known-anchors/util/strings"
	"log"
	"time"
)

func (s *ServiceContext) AuthRegisterPost(c context.Context, req *model.AuthRegisterPostReq) (*model.AuthRegisterPostResp, error) {
	// check if email exists, if so and not activated, delete it, else return error
	uq := s.DBQuery.User
	if user, err := uq.FindByEmail(req.Email); err == nil {
		if user.Activated {
			log.Println(err)
			return nil, errors.New("邮箱已存在")
		} else {
			redis := *s.Redis
			err = redis.Del(c, user.Email)
			if err != nil {
				log.Println(err)
				return nil, errors.New("删除旧有验证码失败")
			}
			// delete user
			err = uq.DeleteByEmail(user.Email)
			if err != nil {
				log.Println(err)
				return nil, errors.New("删除旧有用户失败")
			}
		}
	}
	hashpassword, err := pwd.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return nil, errors.New("哈希密码错误")
	}
	err = uq.Create(&dbmodel.User{
		Email:        req.Email,
		PasswordHash: hashpassword,
		Username:     req.Username,
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("创建用户失败")
	}
	// select code
	user, err := uq.FindByEmail(req.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取用户失败")
	}
	// send mail
	AcToken := strings.RandStringBytes(7)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		err := mail.SendMailCode(req.Username, AcToken, req.Email, 5)
		if err != nil {
			log.Println(err)
		}
		var redisClient = *s.Redis
		// 这里的 context 不能用 c，c 是该请求的 context，但这里是异步的，可能在返回请求后才执行到要 set 的时候，所以要新建一个 context
		err = redisClient.Set(context.Background(), user.Email, AcToken, 5*60*time.Second)
		if err != nil {
			log.Println(err)
		}
	}()
	// mail.SendMailCode()
	resp := model.AuthRegisterPostResp{}
	return &resp, nil
}
