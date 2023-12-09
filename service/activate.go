package service

import (
	"errors"
	"known-anchors/model"
	"known-anchors/util/mail"
	"known-anchors/util/pwd"
	"known-anchors/util/strings"
	"log"
	"time"
)

func (s *ServiceContext) AuthActivatePost(req *model.AuthActivatePostReq) (*model.AuthActivatePostResp, error) {
	uq := s.DBQuery.User
	user, err := uq.FindByEmail(req.Email)
	if err != nil {
		log.Println()
		return nil, errors.New("邮箱不存在")
	}
	if !pwd.CheckPasswordHash(req.Password, user.PasswordHash) {
		log.Println("password error")
		return nil, errors.New("密码错误")
	}
	AcToken := strings.RandStringBytes(7)
	mail.SendMailCode(req.Username, AcToken, req.Email, 5)
	var redis = *s.Redis
	err = redis.Del(s.Ctx, user.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("删除旧有验证码失败")
	}
	err = redis.Set(s.Ctx, req.Email, AcToken, 5*60*time.Second)
	if err != nil {
		log.Println(err)
		return nil, errors.New("设置 Redis 失败")
	}
	resp := model.AuthActivatePostResp{}
	return &resp, nil
}
