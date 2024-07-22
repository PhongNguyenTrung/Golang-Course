package usecase

func (u *BookUsecase) DeleteBook(id string) error {
	return u.bookRepo.DeleteBook(id)
}
