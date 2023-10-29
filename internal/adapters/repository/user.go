package repository

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"server/internal/core/domain"
)

type RepoUser struct {
	db *bun.DB
}

var repoUser *RepoUser

func GetRepoUser(db *bun.DB) *RepoUser {
	if repoUser != nil {
		return repoUser
	}

	repoUser = &RepoUser{
		db: db,
	}

	return repoUser
}

func (r *RepoUser) SaveUser(ctx context.Context, user *domain.User) error {
	r.slice = append(r.slice, *user)

	return nil
}

func (r *RepoUser) GetUser(ctx context.Context, id string) (*domain.User, error) {
	for i := range r.slice {
		if r.slice[i].Id == id {
			return &r.slice[i], nil
		}
	}

	return nil, echo.ErrNotFound
}

func (r *RepoUser) GetUsers(ctx context.Context) ([]*domain.User, error) {
	users := []*domain.User{}

	for i := range r.slice {
		users = append(users, &r.slice[i])
	}

	return users, nil
}
