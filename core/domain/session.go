package domain

import (
	"github.com/google/uuid"
	"net/http"
	"server/core/enums/cookies"
	"time"
)

type Session struct {
	ID        uuid.UUID `json:"id"`
	UserID    uint64    `json:"userId"`
	CreatedAt time.Time `json:"loginTime"`
	IP        uint32    `json:"IP"`
	Expiry    time.Time
}

type IssueSession struct {
	UserID uint64
	IP     uint32
	Expiry time.Time
}

func (s *Session) Cookie() *http.Cookie {
	return &http.Cookie{
		Name:     cookies.Session,
		Value:    s.ID.String(),
		Path:     "/",
		Expires:  s.Expiry,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}
