package pwd

import (
	"known-anchors/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - check password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenToken(email string) (string, error) {
	//expiredtime = 86400
	Expiredtime := time.Now().Add(time.Duration(config.Conf.JWT.Expiredtime) * time.Second)
	c := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Expiredtime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	t, err := token.SignedString([]byte(config.Conf.JWT.Salt))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ParseToken(tokenstring string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.JWT.Salt), nil
	})
	return token, claims, err
}

func StrToUint64(s string) uint64 {
	var res uint64
	for _, c := range s {
		res = res*10 + uint64(c-'0')
	}
	return res
}
