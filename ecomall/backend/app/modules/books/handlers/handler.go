package handlers

import "github.com/1rhino/clean_architecture/app/modules/books/usecase"

type BookHandler struct {
	bookUsecase usecase.IBookUsecase
}

func NewBookHandler(bookUsecase usecase.IBookUsecase) *BookHandler {
	return &BookHandler{bookUsecase: bookUsecase}
}
