package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255)" json:"email" validate:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar" validate:"url"`

	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One to Many relationship
}

func (User) TableName() string {
	return "users"
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UserParams struct {
	Name     string `form:"name" json:"name"`
	Avatar   string `file:"avatar" json:"avatar"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
