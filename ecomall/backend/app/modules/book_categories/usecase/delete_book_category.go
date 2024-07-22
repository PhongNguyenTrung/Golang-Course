package usecase

func (u *BookCategoryUsecase) DeleteBookCategory(bookCategoryID string) error {
	return u.bookCategoryRepo.DeleteBookCategory(bookCategoryID)
}
