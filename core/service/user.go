package service

import (
	"context"
	"server/core/domain"
	"server/core/ports"
)

type User struct {
	userRepo ports.RepositoryUser
}

var userService *User

func GetServiceUser(repo ports.RepositoryUser) *User {
	if userService != nil {
		return userService
	}

	userService = &User{
		userRepo: repo,
	}

	return userService
}

func (s *User) GetUser(ctx context.Context, id uint64) (*domain.GetUser, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *User) GetUsers(ctx context.Context) ([]*domain.GetUser, error) {
	return s.userRepo.GetUsers(ctx)
}
