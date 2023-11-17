package repository

import (
	"context"
	"github.com/labstack/echo/v4"
	"server/core/domain"
)

type RepoUserLocal struct {
	slice []domain.User
}

var repoUserLocal *RepoUserLocal

func GetRepoUserLocal() *RepoUserLocal {
	if repoUserLocal != nil {
		return repoUserLocal
	}

	repoUserLocal = &RepoUserLocal{
		slice: []domain.User{},
	}

	return repoUserLocal
}

func (r *RepoUserLocal) SaveUser(ctx context.Context, user *domain.User) error {
	r.slice = append(r.slice, *user)

	return nil
}

func (r *RepoUserLocal) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	for i := range r.slice {
		if r.slice[i].ID == id {
			return &r.slice[i], nil
		}
	}

	return nil, echo.ErrNotFound
}

func (r *RepoUserLocal) GetUsers(ctx context.Context) ([]*domain.User, error) {
	users := []*domain.User{}

	for i := range r.slice {
		users = append(users, &r.slice[i])
	}

	return users, nil
}
