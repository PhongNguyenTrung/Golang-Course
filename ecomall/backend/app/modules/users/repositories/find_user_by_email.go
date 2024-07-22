package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r UserRepo) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
