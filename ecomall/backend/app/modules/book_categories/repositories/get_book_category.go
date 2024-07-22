package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r *BookCategoryRepo) GetBookCategory(id string) (*models.BookCategory, error) {
	var bookCategory models.BookCategory
	result := r.DB.First(&bookCategory, id)
	return &bookCategory, result.Error
}
