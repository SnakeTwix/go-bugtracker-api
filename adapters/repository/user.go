package repository

import (
	"context"
	"gorm.io/gorm"
	"server/core/domain"
)

type RepoUser struct {
	db *gorm.DB
}

var repoUser *RepoUser

func GetRepoUser(db *gorm.DB) *RepoUser {
	if repoUser != nil {
		return repoUser
	}

	repoUser = &RepoUser{
		db: db,
	}

	return repoUser
}

func (r *RepoUser) SaveUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error

	return user, err
}

func (r *RepoUser) GetUser(ctx context.Context, id uint64) (*domain.GetUser, error) {
	var user *domain.GetUser

	result := r.db.Model(&domain.User{}).Take(&user, id)

	return user, result.Error
}

func (r *RepoUser) GetUsers(ctx context.Context) ([]*domain.GetUser, error) {
	var users []*domain.GetUser

	result := r.db.Model(&domain.User{}).Find(&users)

	return users, result.Error
}
