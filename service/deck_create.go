package service

import (
	"errors"
	dbmodel "known-anchors/dal/db/model"
	"known-anchors/model"
	"log"

	"gorm.io/gorm"
)

func (s *ServiceContext) DeckCreate(uid uint64, req *model.DeckCreateReq) (*model.DeckCreateResp, error) {
	ud := s.DBQuery.Deck
	deck := dbmodel.Deck{
		UserID: uint(uid),
		Name:   req.Name,
		Tags:   req.Tags,
		Ispub:  req.Ispub,
	}
	err := ud.UnderlyingDB().Transaction(func(tx *gorm.DB) error {
		// 创建卡组
		err := ud.Create(&deck)
		if err != nil {
			log.Println(err)
			return errors.New("创建卡组失败")
		}
		// 成批创建卡片
		ucd := s.DBQuery.Card
		cards := make([]*dbmodel.Card, len(req.Cards))
		for i, v := range req.Cards {
			cards[i] = &dbmodel.Card{
				DeckID:  uint(deck.ID),
				OwnerID: uint(uid),
				Front:   v.Front,
				Back:    v.Back,
			}
		}
		err = ucd.CreateInBatches(cards, len(cards))
		if err != nil {
			log.Println(err)
			return errors.New("创建卡片失败")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp := model.DeckCreateResp{
		Id: uint(deck.ID),
	}
	return &resp, nil
}
