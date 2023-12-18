package model

type DeckCreateReq struct {
	// name
	Name string `json:"name"`
	// tags
	Tags string `json:"tags"`
	// ispub
	Ispub bool `json:"ispub"`
	// cards
	Cards []Cards `json:"cards"`
}

type DeckCreateResp struct {
	// id
	Id uint `json:"id"`
}

type DeckGetReq struct {
	// id
	Id uint `json:"id"`
}

type DeckGetResp struct {
	// id
	Id uint `json:"id"`
	// name
	Name string `json:"name"`
	// tags
	Tags string `json:"tags"`
	// ispub
	Ispub bool `json:"ispub"`
	// cards
	Cards []Cards `json:"cards"`
}

type DeckUpdateReq struct {
	// id
	Id uint `json:"id"`
	// name
	Name string `json:"name"`
	// tags
	Tags string `json:"tags"`
	// ispub
	Ispub bool `json:"ispub"`
	// cards
	Cards []Cards `json:"cards"`
}

type DeckUpdateResp struct {
}

type DeckDeleteReq struct {
	// id
	Id uint `json:"id"`
}

type DeckDeleteResp struct {
}

type DeckListReq struct {
	// limit
	Limit int `json:"limit"`
	// offset
	Offset int `json:"offset"`
}

type DeckListResp struct {
	// total
	Total int `json:"total"`
	// decks
	Decks []Decks `json:"decks"`
}

type Decks struct {
	// id
	Id uint `json:"id"`
	// name
	Name string `json:"name"`
	// tags
	Tags string `json:"tags"`
	// ispub
	Ispub bool `json:"ispub"`
}

type Cards struct {
	// front
	Front string `json:"front"`
	// back
	Back string `json:"back"`
}
