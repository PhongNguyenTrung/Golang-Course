package repository

import (
	"errors"

	"github.com/1rhino/clean_architecture/app/models"
)

func (r *BookCategoryRepo) DeleteBookCategory(bookCategoryID string) error {
	result := r.DB.Delete(&models.BookCategory{}, bookCategoryID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("book category not found")
	}
	return nil
}
