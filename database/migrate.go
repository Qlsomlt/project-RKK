package database

import (
	"kode/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
	)
}
