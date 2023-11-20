package ports

import (
	"context"
	"server/core/domain"
)

type ServiceUser interface {
	SaveUser(ctx context.Context, user *domain.CreateUser) (uint64, error)
	GetUser(ctx context.Context, id uint64) (*domain.GetUser, error)
	GetUsers(ctx context.Context) ([]*domain.GetUser, error)
}

type RepositoryUser interface {
	SaveUser(ctx context.Context, user *domain.User) (uint64, error)
	GetUser(ctx context.Context, id uint64) (*domain.GetUser, error)
	GetUsers(ctx context.Context) ([]*domain.GetUser, error)
}
