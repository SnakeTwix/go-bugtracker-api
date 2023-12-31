package repository

import (
	"context"
	"gorm.io/gorm"
	"server/core/domain"
)

type User struct {
	db *gorm.DB
}

var repoUser *User

func GetRepoUser(db *gorm.DB) *User {
	if repoUser != nil {
		return repoUser
	}

	repoUser = &User{
		db: db,
	}

	return repoUser
}

func (r *User) SaveUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error

	return user, err
}

func (r *User) GetUser(ctx context.Context, id uint64) (*domain.GetUser, error) {
	var user *domain.GetUser

	result := r.db.Model(&domain.User{}).Take(&user, id)

	return user, result.Error
}

func (r *User) GetUsers(ctx context.Context) ([]*domain.GetUser, error) {
	var users []*domain.GetUser

	result := r.db.Model(&domain.User{}).Find(&users)

	return users, result.Error
}

func (r *User) GetUserByUsername(ctx context.Context, username *string) (*domain.User, error) {
	var user *domain.User

	result := r.db.Model(&domain.User{}).Where("username = ?", username).Take(&user)

	return user, result.Error
}

func (r *User) UpdateUserRefreshTokenById(ctx context.Context, id uint64, refreshToken *string) error {
	result := r.db.Model(&domain.User{}).Where("id = ?", id).Update("refresh_token", refreshToken)

	return result.Error
}

func (r *User) GetUserByRefreshToken(ctx context.Context, refreshToken *string) (*domain.User, error) {
	var user *domain.User

	result := r.db.Model(&domain.User{}).Where("refresh_token = ?", refreshToken).Take(&user)

	return user, result.Error
}
