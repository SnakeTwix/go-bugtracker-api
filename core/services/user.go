package services

import (
	"context"
	scrypt "github.com/elithrar/simple-scrypt"
	"server/adapters/tools/jwt"
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

func (s *ServiceUser) RegisterUser(ctx context.Context, APIUser *domain.CreateUser) (*domain.Token, error) {
	hash, err := scrypt.GenerateFromPassword([]byte(*APIUser.Password), scrypt.DefaultParams)

	if err != nil {
		return nil, err
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

	createdUser, err := s.userRepo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	tokener := &jwt.TokenGenerator{
		User: createdUser,
	}

	token, err := tokener.Token()
	if err != nil {
		return nil, err
	}

	err = s.userRepo.UpdateUserRefreshTokenById(ctx, createdUser.ID, &token.RefreshToken)
	if err != nil {
		return nil, err
	}

	return token, err
}

func (s *ServiceUser) GetUser(ctx context.Context, id uint64) (*domain.GetUser, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *ServiceUser) GetUsers(ctx context.Context) ([]*domain.GetUser, error) {
	return s.userRepo.GetUsers(ctx)
}

func (s *ServiceUser) LoginUser(ctx context.Context, loginUser *domain.LoginUser) (*domain.Token, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, loginUser.Username)
	if err != nil {
		return nil, err
	}
	if err := scrypt.CompareHashAndPassword([]byte(*user.Password), []byte(*loginUser.Password)); err != nil {
		return nil, err
	}

	tokenGenerator := jwt.TokenGenerator{
		User: user,
	}

	token, err := tokenGenerator.Token()
	if err != nil {
		return nil, err
	}

	err = s.userRepo.UpdateUserRefreshTokenById(ctx, user.ID, &token.RefreshToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}
