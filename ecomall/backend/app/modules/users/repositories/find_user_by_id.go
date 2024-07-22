package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
)

func (r UserRepo) FindUserById(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
