package ports

import (
	"context"
	"github.com/google/uuid"
	"server/core/domain"
)

type ServiceSession interface {
	GetSession(ctx context.Context, name string) (*domain.Session, error)
	SaveSession(ctx context.Context, session *domain.IssueSession) (*domain.Session, error)

	// NewSession Creates a new session without saving it into the store
	NewSession(ip *string, user *domain.User) *domain.IssueSession

	// CreateSession Creates a session and saves it into the store
	CreateSession(ctx context.Context, ip *string, user *domain.User) (*domain.Session, error)
}

type RepositorySession interface {
	GetSession(ctx context.Context, name string) (*domain.Session, error)
	SaveSession(ctx context.Context, session *domain.Session) (*domain.Session, error)
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
}
