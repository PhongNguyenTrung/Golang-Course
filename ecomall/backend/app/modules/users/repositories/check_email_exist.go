package repository

import (
	"fmt"

	"github.com/1rhino/clean_architecture/app/models"
)

func (r UserRepo) CheckEmailExisting(email string) bool {
	var user *models.User

	result := r.DB.Table(models.User{}.TableName()).Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println("err: ", result.Error)

		return false
	}

	return result.RowsAffected > 0
}
