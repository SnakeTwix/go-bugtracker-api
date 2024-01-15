package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"server/core/domain"
)

type Session struct {
	db *gorm.DB
}

var repoSession *Session

func GetRepoSession(db *gorm.DB) *Session {
	if repoSession != nil {
		return repoSession
	}

	repoSession = &Session{
		db: db,
	}

	return repoSession
}

func (r *Session) GetSession(ctx context.Context, name string) (*domain.Session, error) {
	var session *domain.Session
	result := r.db.Where("id = ?", name).Take(&session)
	return session, result.Error
}

func (r *Session) SaveSession(ctx context.Context, session *domain.Session) (*domain.Session, error) {
	result := r.db.Create(session)
	return session, result.Error
}

func (r *Session) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	result := r.db.Delete(&domain.Session{}, sessionID)
	return result.Error
}
