package pwd_test

import (
	"known-anchors/util/pwd"
	"testing"
)

func TestHashword(t *testing.T) {
	password1, err := pwd.HashPassword("123456")
	if err != nil {
		t.Error(err)
	}
	if !pwd.CheckPasswordHash("123456", password1) {
		t.Error("wrong")
	}
	password2, err := pwd.HashPassword("1234567")
	if err != nil {
		t.Error(err)
	}
	if pwd.CheckPasswordHash("123456", password2) {
		t.Error("wrong")
	}
	if pwd.CheckPasswordHash("password1", "password1") {
		t.Error("wrong")
	}
}

func TestToken(t *testing.T) {
	Token, err := pwd.GenToken("email")
	if err != nil {
		t.Error(err)
	}
	token, claims, err := pwd.ParseToken(Token)
	if err != nil || !token.Valid {
		t.Error("wrong")
	}
	if claims.Email != "email" {
		t.Error("wrong")
	}
}
