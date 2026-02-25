package repositories

import (
	"kode/domain"
	"time"

	"gorm.io/gorm"
)

type TokenRepository interface {
	Blacklist(token string, expiresAt int64) error
	IsBlacklisted(token string) (bool, error)
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db}
}

func (r *tokenRepository) Blacklist(token string, expiresAt int64) error {
	return r.db.Create(&domain.BlacklistedToken{
		Token:     token,
		ExpiresAt: time.Unix(expiresAt, 0),
	}).Error
}

func (r *tokenRepository) IsBlacklisted(token string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.BlacklistedToken{}).
		Where("token = ?", token).
		Count(&count).Error

	return count > 0, err
}
