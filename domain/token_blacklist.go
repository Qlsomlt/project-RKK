package domain

import (
	"time"

	"gorm.io/gorm"
)

type BlacklistedToken struct {
	gorm.Model
	Token     string `gorm:"type:text;uniqueIndex"`
	ExpiresAt time.Time
}
