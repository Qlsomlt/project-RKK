package domain

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID    uint
	TokenID   string `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	Revoked   bool
}
