package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"column:username;not null;type:varchar(20)"`
	Email        string `gorm:"column:email;unique;not null;type:varchar(50)"`
	PasswordHash string `gorm:"column:password_hash;not null;type:varchar(67)"`
	Activated    bool   `gorm:"column:activated;default:false;type:boolean"`
	// ActivationToken string `gorm:"column:activation_token:type:varchar(67)"`
	// LastLogin    time.Time `gorm:"column:last_login;default:current_timestamp"`
}

type Post struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id;not null;type:bigint"`
	Title   string `gorm:"column:title;not null;type:varchar(50)"`
	Content string `gorm:"column:content;not null;type:varchar(1000)"`
}

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id;not null;type:bigint"`
	PostID  uint   `gorm:"column:post_id;not null;type:bigint"`
	Content string `gorm:"column:contect;not null;type:varchar(1000)"`
}

type Deck struct {
	gorm.Model
	UserID uint   `gorm:"column:user_id;not null;type:bigint"`
	Name   string `gorm:"column:name;not null;type:varchar(50)"`
	Tags   string `gorm:"column:tags;type:varchar(100)"`
	Ispub  bool   `gorm:"column:ispub;default:false;type:boolean"`
}

type Card struct {
	gorm.Model
	DeckID  uint   `gorm:"column:deck_id;not null;type:bigint"`
	OwnerID uint   `gorm:"column:owner_id;not null;type:bigint"`
	Front   string `gorm:"column:front;not null;type:varchar(100)"`
	Back    string `gorm:"column:back;not null;type:varchar(100)"`
}
