package service

import (
	"context"
	scrypt "github.com/elithrar/simple-scrypt"
	"server/adapters/tools/jwt"
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

func (s *User) RegisterUser(ctx context.Context, APIUser *domain.CreateUser) (*domain.User, error) {
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

	return createdUser, err
}

func (s *User) GetUser(ctx context.Context, id uint64) (*domain.GetUser, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *User) GetUsers(ctx context.Context) ([]*domain.GetUser, error) {
	return s.userRepo.GetUsers(ctx)
}

func (s *User) LoginUser(ctx context.Context, loginUser *domain.LoginUser) (*domain.User, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, loginUser.Username)
	if err != nil {
		return nil, err
	}
	if err := scrypt.CompareHashAndPassword([]byte(*user.Password), []byte(*loginUser.Password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) LogoutUser(ctx context.Context, id uint64) error {
	tokenGenerator := jwt.TokenGenerator{
		User: &domain.User{},
	}

	token, err := tokenGenerator.Token()
	if err != nil {
		return err
	}

	err = s.userRepo.UpdateUserRefreshTokenById(ctx, id, &token.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (s *User) UpdateSession(ctx context.Context, refreshToken *string) (*domain.Token, error) {
	user, err := s.userRepo.GetUserByRefreshToken(ctx, refreshToken)
	if err != nil {
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
