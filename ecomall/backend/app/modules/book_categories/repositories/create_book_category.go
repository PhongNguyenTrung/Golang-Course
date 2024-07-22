package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
)

func (r *BookCategoryRepo) CreateBookCategory(payload *models.BookCategoryParams) (*models.BookCategory, error) {
	var bookCategory = &models.BookCategory{
		Name:        payload.Name,
		Image:       payload.Image,
		Description: payload.Description,
	}

	result := r.DB.Create(bookCategory)
	return bookCategory, result.Error
}
