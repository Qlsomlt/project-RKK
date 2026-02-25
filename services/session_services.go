package services

import (
	"errors"
	"kode/repositories"
)

type SessionService interface {
	Logout(userID uint, tokenID string) error
	LogoutAll(userID uint) error
}

type sessionService struct {
	repo repositories.SessionRepository
}

func NewSessionService(r repositories.SessionRepository) SessionService {
	return &sessionService{r}
}

// Logout implements [SessionService].
func (s *sessionService) Logout(userID uint, tokenID string) error {
	session, err := s.repo.FindByTokenID(tokenID)
	if err != nil {
		return err
	}

	// Pastikan session milik user
	if session.UserID != userID {
		return errors.New("unauthorized logout attempt")
	}

	return s.repo.Revoke(tokenID)
}

// LogoutAll implements [SessionService].
func (s *sessionService) LogoutAll(userID uint) error {
	return s.repo.RevokeAll(userID)
}
