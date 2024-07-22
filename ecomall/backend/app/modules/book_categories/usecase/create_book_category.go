package usecase

import "github.com/1rhino/clean_architecture/app/models"

func (u *BookCategoryUsecase) CreateBookCategory(payload *models.BookCategoryParams) (*models.BookCategory, error) {
	return u.bookCategoryRepo.CreateBookCategory(payload)
}
