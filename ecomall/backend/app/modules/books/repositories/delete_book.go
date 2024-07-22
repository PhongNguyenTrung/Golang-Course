package repository

import (
	"errors"

	"github.com/1rhino/clean_architecture/app/models"
)

func (r *BookRepo) DeleteBook(id string) error {
	result := r.DB.Unscoped().Delete(&models.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}
