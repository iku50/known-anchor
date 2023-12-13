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
	uq := s.DBQuery.User
	_, err := uq.FindByEmail(req.Email)
	if err == nil {
		log.Println(err)
		return nil, errors.New("邮箱已存在")
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
		err = redisClient.Set(c, user.Email, AcToken, 5*60*time.Second)
		if err != nil {
			log.Println(err)
		}
	}()
	// mail.SendMailCode()
	resp := model.AuthRegisterPostResp{}
	return &resp, nil
}
