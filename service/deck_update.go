package service

import (
	"errors"
	dbmodel "known-anchors/dal/db/model"
	"known-anchors/model"
	"log"
)

func (s *ServiceContext) DeckUpdate(uid uint64, req *model.DeckUpdateReq) (*model.DeckUpdateResp, error) {
	ud := s.DBQuery.Deck

	deck, err := ud.FindByID(req.Id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取卡组失败")
	}
	if deck.UserID != uint(uid) {
		return nil, errors.New("无权修改卡组")
	}

	_, err = ud.Where(ud.ID.Eq(req.Id)).Updates(&dbmodel.Deck{
		Name:  req.Name,
		Tags:  req.Tags,
		Ispub: req.Ispub,
	})

	if err != nil {
		log.Println(err)
		return nil, errors.New("更新卡组失败")
	}
	resp := model.DeckUpdateResp{}
	return &resp, nil
}
