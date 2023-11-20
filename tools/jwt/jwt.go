package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	Id        uint64 `json:"id"`
	Privilege string `json:"privilege"`
	jwt.RegisteredClaims
}

func CreateToken(id uint64, privilege string) (string, error) {
	day := time.Hour * 24

	claims := &UserClaims{
		Id:        id,
		Privilege: privilege,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(day * 7)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}
