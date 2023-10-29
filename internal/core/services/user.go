package services

import (
	"context"
	"github.com/google/uuid"
	"server/internal/core/domain"
	"server/internal/core/ports"
)

type ServiceUser struct {
	userRepo ports.RepositoryUser
}

var userService *ServiceUser

func GetServiceUser(repo ports.RepositoryUser) *ServiceUser {
	if userService != nil {
		return userService
	}

	userService = &ServiceUser{
		userRepo: repo,
	}

	return userService
}

func (s *ServiceUser) SaveUser(ctx context.Context, user *domain.User) error {
	user.Id = uuid.NewString()

	return s.userRepo.SaveUser(ctx, user)
}

func (s *ServiceUser) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *ServiceUser) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return s.userRepo.GetUsers(ctx)
}
