package service

import (
	"errors"
	"known-anchors/model"
	"log"
)

func (s *ServiceContext) AuthConfirmPost(req *model.AuthConfirmPostReq) (*model.AuthConfirmPostResp, error) {
	redis := *s.Redis
	uq := s.DBQuery.User
	Token, err := redis.Get(s.Ctx, req.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("无法获取 Token")
	}
	if Token != req.ConfirmToken {
		log.Println("token error")
		return nil, errors.New("验证码错误")
	}
	err = uq.UpdateActivatedByEmail(req.Email, true)
	if err != nil {
		log.Println(err)
		return nil, errors.New("激活失败")
	}
	redis.Del(s.Ctx, req.Email)
	resp := model.AuthConfirmPostResp{}
	return &resp, nil
}
