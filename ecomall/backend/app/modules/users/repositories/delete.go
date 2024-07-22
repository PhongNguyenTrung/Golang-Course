package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r *UserRepo) DeleteUser(user *models.User) error {
	result := r.DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
