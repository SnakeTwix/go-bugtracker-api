package service

import (
	"context"
	"github.com/google/uuid"
	"server/core/domain"
	"server/core/ports"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	repoSession ports.RepositorySession
}

var serviceSession *Session

func GetServiceSession(repo ports.RepositorySession) *Session {
	if serviceSession != nil {
		return serviceSession
	}

	serviceSession = &Session{
		repoSession: repo,
	}

	return serviceSession
}

func (s *Session) GetSession(ctx context.Context, name string) (*domain.Session, error) {
	return s.repoSession.GetSession(ctx, name)
}

func (s *Session) SaveSession(ctx context.Context, issueSession *domain.IssueSession) (*domain.Session, error) {
	session := &domain.Session{
		ID:     uuid.New(),
		UserID: issueSession.UserID,
		IP:     issueSession.IP,
		Expiry: issueSession.Expiry,
	}

	return s.repoSession.SaveSession(ctx, session)
}

func (s *Session) NewSession(IP *string, user *domain.User) *domain.IssueSession {
	splitIP := strings.Split(*IP, ".")
	var numIP uint32 = 0
	for index, IPPart := range splitIP {
		// 8 bits because just one ip part
		numIPPart, err := strconv.ParseUint(IPPart, 10, 8)
		if err != nil {
			numIP = 0
			break
		}

		// One by one this will set all the ip parts based on the index
		// E.g. numIP = 010000100 00000000 00000000 00000000 numIPPart = 10101010 index = 1
		// numIP | 10101010 << 16 == 010000100 10101010 00000000 00000000
		numIP = numIP | uint32(numIPPart)<<((3-index)*8)
	}

	issueSession := &domain.IssueSession{
		UserID: user.ID,
		IP:     numIP,
		Expiry: time.Now().Add(time.Hour * 24 * 7),
	}

	return issueSession
}

func (s *Session) CreateSession(ctx context.Context, IP *string, user *domain.User) (*domain.Session, error) {
	issueSession := s.NewSession(IP, user)
	return s.SaveSession(ctx, issueSession)
}
