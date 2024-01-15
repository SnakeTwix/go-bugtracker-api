package ports

import (
	"context"
	"github.com/google/uuid"
	"server/core/domain"
)

type ServiceAuth interface {
	// RegisterUser TODO: Return domain.GetUser instead of domain.User
	RegisterUser(ctx context.Context, user *domain.CreateUser) (*domain.User, error)

	// LoginUser TODO: Return domain.GetUser instead of domain.User
	LoginUser(ctx context.Context, user *domain.LoginUser) (*domain.User, error)

	LogoutUser(ctx context.Context, sessionId uuid.UUID) error
}
