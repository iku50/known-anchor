package service

import (
	"errors"
	"known-anchors/model"
	"log"
)

func (s *ServiceContext) DeckGet(uid uint64, req *model.DeckGetReq) (*model.DeckGetResp, error) {
	ud := s.DBQuery.Deck
	deck, err := ud.FindByID(req.Id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取卡组失败")
	}
	if deck.UserID != uint(uid) && !deck.Ispub {
		return nil, errors.New("无权查看卡组")
	}
	resp := model.DeckGetResp{
		Id:    deck.ID,
		Name:  deck.Name,
		Tags:  deck.Tags,
		Ispub: deck.Ispub,
	}
	uc := s.DBQuery.Card
	cards, err := uc.ListByDeckID(deck.ID, 5, 0)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取卡片失败")
	}
	resp.Cards = make([]model.Cards, len(cards))
	for i, v := range cards {
		resp.Cards[i] = model.Cards{
			Front: v.Front,
			Back:  v.Back,
		}
	}

	return &resp, nil
}
