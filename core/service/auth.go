package service

import (
	"context"
	scrypt "github.com/elithrar/simple-scrypt"
	"github.com/google/uuid"
	"server/core/domain"
	"server/core/ports"
)

type Auth struct {
	repoUser    ports.RepositoryUser
	repoSession ports.RepositorySession
}

var serviceAuth *Auth

func GetServiceAuth(repoUser ports.RepositoryUser, repoSession ports.RepositorySession) *Auth {
	if serviceAuth != nil {
		return serviceAuth
	}

	serviceAuth = &Auth{
		repoUser,
		repoSession,
	}

	return serviceAuth
}

func (s *Auth) RegisterUser(ctx context.Context, APIUser *domain.CreateUser) (*domain.User, error) {
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

	createdUser, err := s.repoUser.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, err
}

func (s *Auth) LoginUser(ctx context.Context, loginUser *domain.LoginUser) (*domain.User, error) {
	user, err := s.repoUser.GetUserByUsername(ctx, loginUser.Username)
	if err != nil {
		return nil, err
	}
	if err := scrypt.CompareHashAndPassword([]byte(*user.Password), []byte(*loginUser.Password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Auth) LogoutUser(ctx context.Context, sessionID uuid.UUID) error {
	return s.repoSession.DeleteSession(ctx, sessionID)
}
