package domain

import "gorm.io/gorm"

type Role string

const (
	SUPERADMIN Role = "superadmin"
	ADMIN      Role = "admin"
	USER       Role = "user"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role     Role   `json:"role" gorm:"type:varchar(20);default:'user'"`
}
