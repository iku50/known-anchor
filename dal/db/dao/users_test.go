package dao_test

import (
	"known-anchors/dal"
	"known-anchors/dal/db/dao"
	"known-anchors/dal/db/model"
	"testing"
)

func TestUserDao_FindByEmail(t *testing.T) {
	// set excution path to the root of the project
	// so that the config file can be found
	db := dao.Use(dal.DB.Debug())
	// insert a user named wiz
	user := model.User{
		Username:     "wiz",
		Email:        "wiz",
		PasswordHash: "wiz",
	}
	err := db.User.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	user, err = db.User.FindByEmail("wiz")
	if err != nil {
		t.Fatal(err)
	}
	if user.Username != "wiz" {
		t.Fatal("username not match")
	}
	// delete the user named wiz
	err = db.User.DeleteByEmail("wiz")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDao_FindByID(t *testing.T) {
	// set excution path to the root of the project
	// so that the config file can be found
	db := dao.Use(dal.DB.Debug())
	// insert a user named wiz
	user := model.User{
		Username:     "wiz",
		Email:        "wiz",
		PasswordHash: "wiz",
	}
	err := db.User.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	user, err = db.User.FindByID(user.ID)
	if err != nil {
		t.Fatal(err)
	}
	if user.Username != "wiz" {
		t.Fatal("username not match")
	}
	// delete the user named wiz
	err = db.User.DeleteByEmail("wiz")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDao_DeleteByEmail(t *testing.T) {
	// set excution path to the root of the project
	// so that the config file can be found
	db := dao.Use(dal.DB.Debug())
	// insert a user named wiz
	user := model.User{
		Username:     "wiz",
		Email:        "wiz",
		PasswordHash: "wiz",
	}
	err := db.User.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	// delete the user named wiz
	err = db.User.DeleteByEmail("wiz")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDao_UpdateByID(t *testing.T) {
	// set excution path to the root of the project
	// so that the config file can be found
	db := dao.Use(dal.DB.Debug())
	// insert a user named wiz
	user := model.User{
		Username:     "wiz",
		Email:        "wiz",
		PasswordHash: "wiz",
	}
	err := db.User.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	u, err := db.User.FindByEmail("wiz")
	if err != nil {
		t.Fatal(err)
	}
	// update the user named wiz
	err = db.User.UpdateByID(u.ID, "wizzz", "wiz", "wiz")
	if err != nil {
		t.Fatal(err)
	}
	u, err = db.User.FindByID(u.ID)
	if err != nil {
		t.Fatal(err)
	}
	if u.Email != "wizzz" {
		t.Fatal("username not match")
	}
	// delete the user named wiz
	err = db.User.DeleteByEmail("wiz")
	if err != nil {
		t.Fatal(err)
	}
}
