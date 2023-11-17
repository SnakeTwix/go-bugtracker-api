package services

import (
	"context"
	scrypt "github.com/elithrar/simple-scrypt"
	"server/core/domain"
	"server/core/ports"
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

func (s *ServiceUser) SaveUser(ctx context.Context, APIUser *domain.CreateUser) error {
	hash, err := scrypt.GenerateFromPassword([]byte(*APIUser.Password), scrypt.DefaultParams)

	if err != nil {
		return err
	}

	hashStr := string(hash)
	status := "active"
	privilege := "user"

	user := &domain.User{
		Username:  APIUser.Username,
		Password:  &hashStr,
		Status:    &status,
		Privilege: &privilege,
	}
	return s.userRepo.SaveUser(ctx, user)
}

func (s *ServiceUser) GetUser(ctx context.Context, id uint64) (*domain.GetUser, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *ServiceUser) GetUsers(ctx context.Context) ([]*domain.GetUser, error) {
	return s.userRepo.GetUsers(ctx)
}
