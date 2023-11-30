package ports

import (
	"server/core/domain"
)

// Tokener Interface for generating JWT tokens and a refresh token
// As well as getting the cookie to be sent
type Tokener interface {
	Token() (*domain.Token, error)
}
