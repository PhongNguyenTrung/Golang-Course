package usecase

import "github.com/1rhino/clean_architecture/app/models"

func (h *BookCategoryUsecase) GetBookCategories() ([]*models.BookCategory, error) {
	bookCategories, err := h.bookCategoryRepo.GetBookCategories()
	if err != nil {
		return nil, err
	}

	return bookCategories, nil
}
