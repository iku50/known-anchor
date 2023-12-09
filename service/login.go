package service

import (
	"context"
	"errors"
	"known-anchors/config"
	"known-anchors/model"
	"known-anchors/util/pwd"
	"log"
	"time"
)

func (s *ServiceContext) AuthLoginPost(req *model.AuthLoginPostReq) (*model.AuthLoginPostResp, error) {
	uq := s.DBQuery.User
	user, err := uq.FindByEmail(req.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("邮箱不存在")
	}
	if !pwd.CheckPasswordHash(req.Password, user.PasswordHash) {
		log.Println("password error")
		return nil, errors.New("密码错误")
	}
	if !user.Activated {
		log.Println("user not activated")
		return nil, errors.New("用户未激活")
	}
	token, err := pwd.GenToken(user.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("生成 token 失败")
	}
	resp := model.AuthLoginPostResp{
		Token: token,
	}
	// save token to redis
	err = (*s.Redis).Set(context.Background(), token, user.ID, time.Duration(config.Conf.JWT.Expiredtime)*time.Second)
	if err != nil {
		log.Println(err)
		return nil, errors.New("无法存储 Token 到 Redis")
	}
	return &resp, nil
}
