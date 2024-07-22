package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r *BookCategoryRepo) GetBookCategories() ([]*models.BookCategory, error) {
	var bookCategories []*models.BookCategory
	result := r.DB.Find(&bookCategories)
	return bookCategories, result.Error
}
