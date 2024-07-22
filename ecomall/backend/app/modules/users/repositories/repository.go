package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	CheckEmailExisting(email string) bool
	CreateUser(data *models.SignUpInput) (*models.User, error)
	UpdateUser(data *models.User, payload *models.UserParams) (*models.User, error)
	DeleteUser(data *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{DB: db}
}
