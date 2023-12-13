package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string `gorm:"column:username;unique;not null"`
	Email           string `gorm:"column:email;unique;not null"`
	PasswordHash    string `gorm:"column:password_hash;not null"`
	Activated       bool   `gorm:"column:activated;default:false"`
	ActivationToken string `gorm:"column:activation_token"`
	// LastLogin    time.Time `gorm:"column:last_login;default:current_timestamp"`
	// IsActive     bool      `gorm:"column:is_active;default:true"`
}

type Post struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id;not null"`
	Title   string `gorm:"column:title;not null"`
	Content string `gorm:"column:content;not null"`
}

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id;not null"`
	PostID  uint   `gorm:"column:post_id;not null"`
	Content string `gorm:"column:contect;not null"`
}

type Deck struct {
	gorm.Model
	UserID uint   `gorm:"column:user_id;not null"`
	Name   string `gorm:"column:name;not null"`
	Tags   string `gorm:"column:tags"`
	Ispub  bool   `gorm:"column:ispub;default:false"`
}

type Card struct {
	gorm.Model
	DeckID      uint      `gorm:"column:deck_id;not null"`
	UserID      uint      `gorm:"column:user_id;not null"`
	Front       string    `gorm:"column:front;not null"`
	Back        string    `gorm:"column:back;not null"`
	Status      string    `gorm:"column:status;type:ENUM('new', 'review', 'mastered');default:'new'"`
	LastReview  time.Time `gorm:"column:last_review"`
	NextReview  time.Time `gorm:"column:next_review"`
	Interval    int       `gorm:"column:interval"`
	Repetitions int       `gorm:"column:repetitions"`
	Ease        float64   `gorm:"column:ease"`
}
