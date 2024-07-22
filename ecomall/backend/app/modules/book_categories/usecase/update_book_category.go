package usecase

import "github.com/1rhino/clean_architecture/app/models"

func (u *BookCategoryUsecase) UpdateBookCategory(id string, params *models.BookCategoryParams) (*models.BookCategory, error) {
	bookCategory, err := u.bookCategoryRepo.GetBookCategory(id)
	if err != nil {
		return nil, err
	}

	return u.bookCategoryRepo.UpdateBookCategory(bookCategory, params)
}
