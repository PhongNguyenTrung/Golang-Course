package handlers

import "github.com/1rhino/clean_architecture/app/modules/book_categories/usecase"

type BookCategoryHandler struct {
	bookCategoryUsecase usecase.IBookCategoryUsecase
}

func NewBookCategoryHandler(bookCategoryUsecase usecase.IBookCategoryUsecase) *BookCategoryHandler {
	return &BookCategoryHandler{bookCategoryUsecase: bookCategoryUsecase}
}
