package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
	"gorm.io/gorm"
)

type IBookCategoryRepo interface {
	CreateBookCategory(payload *models.BookCategoryParams) (*models.BookCategory, error)
	GetBookCategories() ([]*models.BookCategory, error)
	GetBookCategory(id string) (*models.BookCategory, error)
	UpdateBookCategory(bookCategory *models.BookCategory, payload *models.BookCategoryParams) (*models.BookCategory, error)
	DeleteBookCategory(bookCategoryID string) error
}

type BookCategoryRepo struct {
	DB *gorm.DB
}

func NewBookCategoryRepo(DB *gorm.DB) IBookCategoryRepo {
	return &BookCategoryRepo{DB: DB}
}
