package service

import (
	"errors"
	"known-anchors/model"
	"log"
)

func (s *ServiceContext) DeckDelete(uid uint64, req *model.DeckDeleteReq) (*model.DeckDeleteResp, error) {
	ud := s.DBQuery.Deck

	deck, err := ud.FindByID(req.Id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取卡组失败")
	}
	if deck.UserID != uint(uid) {
		return nil, errors.New("无权删除卡组")
	}

	_, err = ud.Where(ud.ID.Eq(req.Id)).Delete()

	if err != nil {
		log.Println(err)
		return nil, errors.New("删除卡组失败")
	}
	resp := model.DeckDeleteResp{}
	return &resp, nil
}
