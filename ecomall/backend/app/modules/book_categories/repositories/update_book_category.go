package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r *BookCategoryRepo) UpdateBookCategory(bookCategory *models.BookCategory, payload *models.BookCategoryParams) (*models.BookCategory, error) {
	result := r.DB.Model(&bookCategory).Updates(payload)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookCategory, nil
}
