package ports

import (
	"context"
	"server/core/domain"
)

type ServiceUser interface {
	// RegisterUser TODO: Return domain.GetUser instead of domain.User

	RegisterUser(ctx context.Context, user *domain.CreateUser) (*domain.User, error)
	// LoginUser TODO: Return domain.GetUser instead of domain.User

	LoginUser(ctx context.Context, user *domain.LoginUser) (*domain.User, error)
	GetUser(ctx context.Context, id uint64) (*domain.GetUser, error)
	GetUsers(ctx context.Context) ([]*domain.GetUser, error)
	LogoutUser(ctx context.Context, id uint64) error
	UpdateSession(ctx context.Context, refreshToken *string) (*domain.Token, error)
}

type RepositoryUser interface {
	SaveUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id uint64) (*domain.GetUser, error)
	GetUsers(ctx context.Context) ([]*domain.GetUser, error)
	GetUserByUsername(ctx context.Context, username *string) (*domain.User, error)
	UpdateUserRefreshTokenById(ctx context.Context, id uint64, refreshToken *string) error
	GetUserByRefreshToken(ctx context.Context, refreshToken *string) (*domain.User, error)
}
