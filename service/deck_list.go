package service

import (
	"errors"
	"known-anchors/model"
	"log"
)

func (s *ServiceContext) DeckList(uid uint64, req *model.DeckListReq) (*model.DeckListResp, error) {
	ud := s.DBQuery.Deck
	deck, err := ud.ListByUserID((uint)(uid), req.Limit, req.Offset)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取卡组失败")
	}
	total := len(deck)
	resp := model.DeckListResp{
		Decks: make([]model.Decks, total),
	}
	for i := 0; i < total; i++ {
		resp.Decks[i].Id = deck[i].ID
		resp.Decks[i].Name = deck[i].Name
		resp.Decks[i].Tags = deck[i].Tags
		resp.Decks[i].Ispub = deck[i].Ispub
	}
	resp.Total = total
	return &resp, nil
}
