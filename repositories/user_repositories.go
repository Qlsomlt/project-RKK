package repositories

import (
	"errors"
	"kode/domain"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
	FindAll() ([]domain.User, error)

	FindByTokenID(tokenID string) (*domain.Session, error)
	FindActiveByUserID(userID uint) ([]domain.Session, error)
	CountActiveByUserID(userID uint) (int64, error)

	Revoke(tokenID string) error
	RevokeAll(userID uint) error

	DeleteExpired() error
	DeleteByUserID(userID uint) error
	Exists(tokenID string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByTokenID(tokenID string) (*domain.Session, error) {

	var session domain.Session

	err := r.db.
		Where("token_id = ?", tokenID).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *userRepository) FindActiveByUserID(userID uint) ([]domain.Session, error) {

	var sessions []domain.Session

	err := r.db.
		Where("user_id = ? AND revoked = false AND expires_at > ?", userID, time.Now()).
		Find(&sessions).Error

	return sessions, err
}

func (r *userRepository) CountActiveByUserID(userID uint) (int64, error) {

	var count int64

	err := r.db.
		Model(&domain.Session{}).
		Where("user_id = ? AND revoked = false AND expires_at > ?", userID, time.Now()).
		Count(&count).Error

	return count, err
}

func (r *userRepository) Revoke(tokenID string) error {

	result := r.db.
		Model(&domain.Session{}).
		Where("token_id = ?", tokenID).
		Update("revoked", true)

	if result.RowsAffected == 0 {
		return errors.New("session not found")
	}

	return result.Error
}

func (r *userRepository) RevokeAll(userID uint) error {

	return r.db.
		Model(&domain.Session{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}

func (r *userRepository) DeleteExpired() error {

	return r.db.
		Where("expires_at < ?", time.Now()).
		Delete(&domain.Session{}).Error
}

func (r *userRepository) DeleteByUserID(userID uint) error {

	return r.db.
		Where("user_id = ?", userID).
		Delete(&domain.Session{}).Error
}

func (r *userRepository) Exists(tokenID string) (bool, error) {

	var count int64

	err := r.db.
		Model(&domain.Session{}).
		Where("token_id = ? AND revoked = false AND expires_at > ?", tokenID, time.Now()).
		Count(&count).Error

	return count > 0, err
}
