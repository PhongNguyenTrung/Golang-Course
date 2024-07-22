package usecase

import "github.com/1rhino/clean_architecture/app/models"

func (u *BookCategoryUsecase) GetBookCategory(id string) (*models.BookCategory, error) {
	return u.bookCategoryRepo.GetBookCategory(id)
}
