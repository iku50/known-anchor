package model

import "gorm.io/gen"

type Method interface {
	//Where("id=@id")
	FindByID(id uint) (gen.T, error)
}

type UserMethod interface {
	//Where("email=@email")
	FindByEmail(email string) (gen.T, error)
	//delete from users where email=@email
	DeleteByEmail(email string) error
	//update users set username=@username,email=@email,password_hash=@password where id=@id
	UpdateByID(id uint, email string, username string, password string) error
	//update users set activated=@activated where email=@email
	UpdateActivatedByEmail(email string, activated bool) error
}

type DeckMethod interface {
	// select * from decks where user_id=@user_id limit @limit offset @offset
	ListByUserID(user_id uint, limit int, offset int) ([]gen.T, error)
	// select count(*) from decks where user_id=@user_id
	CountByUserID(user_id uint) (int64, error)
}
