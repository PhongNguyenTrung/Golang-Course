package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
	"golang.org/x/crypto/bcrypt"
)

func (r UserRepo) CreateUser(data *models.SignUpInput) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	var user = &models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hashedPassword),
	}

	result := r.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
