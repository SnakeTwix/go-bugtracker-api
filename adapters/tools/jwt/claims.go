package jwt

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Id        uint64 `json:"id"`
	Privilege string `json:"privilege"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	Id uint64 `json:"id"`
	jwt.RegisteredClaims
}
