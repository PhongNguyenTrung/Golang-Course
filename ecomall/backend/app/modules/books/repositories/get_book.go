package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r *BookRepo) GetBook(user *models.User, id string) (*models.Book, error) {
	book := &models.Book{}
	result := r.DB.Preload("BookCategory").Where("id = ? AND user_id = ?", id, user.ID).First(book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}
