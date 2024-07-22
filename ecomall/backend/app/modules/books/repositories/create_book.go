package repository

import "github.com/1rhino/clean_architecture/app/models"

func (r *BookRepo) CreateBook(user *models.User, payload models.BookParams) (*models.Book, error) {
	book := models.Book{
		UserID:         user.ID,
		BookCategoryID: payload.BookCategoryID,
		Title:          payload.Title,
		Author:         payload.Author,
		Image:          payload.Image,
		PublishDate:    payload.PublishDate,
		Description:    payload.Description,
	}

	if err := r.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("BookCategory").First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
