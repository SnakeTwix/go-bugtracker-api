package ports

import (
	"context"
	"server/internal/core/domain"
)

type ServiceUser interface {
	SaveUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]*domain.User, error)
}

type RepositoryUser interface {
	SaveUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]*domain.User, error)
}
