package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"server/core/domain"
	"time"
)

type TokenGenerator struct {
	User *domain.User
}

func (g TokenGenerator) Token() (*domain.Token, error) {
	day := time.Hour * 24

	// Currently the token is held for 7 days
	// Might consider making it less when I figure out how to implement refresh tokens properly
	userClaims := &UserClaims{
		Id:        g.User.ID,
		Privilege: *g.User.Privilege,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(day * 7)),
		},
	}

	userToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims).SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshClaims := &RefreshClaims{
		Id: g.User.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	//refreshToken := "IMPLEMENT"

	token := &domain.Token{
		Jwt:          userToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
