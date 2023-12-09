package service

import (
	"errors"
	"known-anchors/model"
	"known-anchors/util/pwd"
	"log"
)

func (s *ServiceContext) UserUpdate(uid uint64, req *model.UserUpdateReq) (*model.UserUpdateResp, error) {
	// switch uid to uint
	passwordHash, err := pwd.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return nil, errors.New("哈希密码错误")
	}
	err = s.DBQuery.User.UpdateByID((uint)(uid), req.Email, req.Username, passwordHash)
	if err != nil {
		log.Println(err)
		return nil, errors.New("更新用户失败")
	}
	resp := model.UserUpdateResp{}
	return &resp, nil
}
