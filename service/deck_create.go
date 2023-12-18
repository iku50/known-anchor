package service

import (
	"errors"
	dbmodel "known-anchors/dal/db/model"
	"known-anchors/model"
	"log"
)

func (s *ServiceContext) DeckCreate(uid uint64, req *model.DeckCreateReq) (*model.DeckCreateResp, error) {
	ud := s.DBQuery.Deck
	err := ud.Create(&dbmodel.Deck{
		UserID: uint(uid),
		Name:   req.Name,
		Tags:   req.Tags,
		Ispub:  req.Ispub,
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("创建卡组失败")
	}
	resp := model.DeckCreateResp{}
	return &resp, nil
}
