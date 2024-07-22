package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r *BookRepo) UpdateBook(user *models.User, id string, payload models.BookParams) (*models.Book, error) {
	book := &models.Book{}
	if err := r.DB.Where("user_id = ? AND id = ?", user.ID, id).First(&book).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Model(&book).Updates(payload).Error; err != nil {
		return nil, err
	}

	// After updating, preload the BookCategory to include it in the returned book
	if err := r.DB.Preload("BookCategory").First(&book, book.ID).Error; err != nil {
		return nil, err
	}
	return book, nil
}
