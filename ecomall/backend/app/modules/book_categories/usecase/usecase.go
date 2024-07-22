package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	repository "github.com/1rhino/clean_architecture/app/modules/book_categories/repositories"
)

type IBookCategoryUsecase interface {
	CreateBookCategory(*models.BookCategoryParams) (*models.BookCategory, error)
	GetBookCategories() ([]*models.BookCategory, error)
	GetBookCategory(id string) (*models.BookCategory, error)
	UpdateBookCategory(id string, params *models.BookCategoryParams) (*models.BookCategory, error)
	DeleteBookCategory(bookCategoryID string) error
}

type BookCategoryUsecase struct {
	bookCategoryRepo repository.IBookCategoryRepo
}

func NewBookCategoryUsecase(bookCategoryRepo repository.IBookCategoryRepo) IBookCategoryUsecase {
	return &BookCategoryUsecase{bookCategoryRepo: bookCategoryRepo}
}
