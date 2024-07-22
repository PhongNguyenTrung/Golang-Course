package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" binding:"required,email" json:"email"`
	Password string `json:"password,omitempty" binding:"required"`
}
