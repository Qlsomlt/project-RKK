package repositories

import (
	"errors"
	"time"

	"kode/domain"

	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(session *domain.Session) error
	FindByTokenID(tokenID string) (*domain.Session, error)
	FindActiveByUserID(userID uint) ([]domain.Session, error)
	CountActiveByUserID(userID uint) (int64, error)

	Revoke(tokenID string) error
	RevokeAll(userID uint) error

	DeleteExpired() error
	DeleteByUserID(userID uint) error
	Exists(tokenID string) (bool, error)
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(session *domain.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepository) FindByTokenID(tokenID string) (*domain.Session, error) {

	var session domain.Session

	err := r.db.
		Where("token_id = ?", tokenID).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *sessionRepository) FindActiveByUserID(userID uint) ([]domain.Session, error) {

	var sessions []domain.Session

	err := r.db.
		Where("user_id = ? AND revoked = false AND expires_at > ?", userID, time.Now()).
		Find(&sessions).Error

	return sessions, err
}

func (r *sessionRepository) CountActiveByUserID(userID uint) (int64, error) {

	var count int64

	err := r.db.
		Model(&domain.Session{}).
		Where("user_id = ? AND revoked = false AND expires_at > ?", userID, time.Now()).
		Count(&count).Error

	return count, err
}

func (r *sessionRepository) Revoke(tokenID string) error {

	result := r.db.
		Model(&domain.Session{}).
		Where("token_id = ?", tokenID).
		Update("revoked", true)

	if result.RowsAffected == 0 {
		return errors.New("session not found")
	}

	return result.Error
}

func (r *sessionRepository) RevokeAll(userID uint) error {

	return r.db.
		Model(&domain.Session{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}

func (r *sessionRepository) DeleteExpired() error {

	return r.db.
		Where("expires_at < ?", time.Now()).
		Delete(&domain.Session{}).Error
}

func (r *sessionRepository) DeleteByUserID(userID uint) error {

	return r.db.
		Where("user_id = ?", userID).
		Delete(&domain.Session{}).Error
}

func (r *sessionRepository) Exists(tokenID string) (bool, error) {

	var count int64

	err := r.db.
		Model(&domain.Session{}).
		Where("token_id = ? AND revoked = false AND expires_at > ?", tokenID, time.Now()).
		Count(&count).Error

	return count > 0, err
}
