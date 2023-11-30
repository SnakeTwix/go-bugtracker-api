package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"server/core/domain"
	"time"
)

type UserClaims struct {
	Id        uint64 `json:"id"`
	Privilege string `json:"privilege"`
	jwt.RegisteredClaims
}

type TokenGenerator struct {
	User *domain.User
}

func (g TokenGenerator) Token() (*domain.Token, error) {
	day := time.Hour * 24

	// Currently the token is held for 7 days
	// Might consider making it less when I figure out how to implement refresh tokens properly
	claims := &UserClaims{
		Id:        g.User.ID,
		Privilege: *g.User.Privilege,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(day * 7)),
		},
	}

	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))

	if err != nil {
		return nil, err
	}

	token := &domain.Token{
		Jwt:          jwtToken,
		RefreshToken: "TODO: IMPLEMENT",
	}

	return token, nil
}
