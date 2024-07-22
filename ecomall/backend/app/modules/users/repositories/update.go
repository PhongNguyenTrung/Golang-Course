package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
)

func (r UserRepo) UpdateUser(user *models.User, payload *models.UserParams) (*models.User, error) {
	result := r.DB.Model(&user).Updates(payload)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
