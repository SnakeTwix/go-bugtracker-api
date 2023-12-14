package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ParseUserClaims(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, echo.ErrBadRequest
	}

	return userClaims, nil
}
