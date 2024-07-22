package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	repository "github.com/1rhino/clean_architecture/app/modules/books/repositories"
	"github.com/gin-gonic/gin"
	"github.com/rosberry/go-pagination"
)

type IBookUsecase interface {
	CreateBook(c *gin.Context, payload models.BookParams) (*models.Book, error)
	GetBooks(c *gin.Context) ([]*models.Book, *pagination.Paginator, error)
	GetBook(c *gin.Context) (*models.Book, error)
	UpdateBook(c *gin.Context, payload models.BookParams) (*models.Book, error)
	DeleteBook(id string) error
}

type BookUsecase struct {
	bookRepo repository.IBookRepo
}

func NewBookUsecase(bookRepo repository.IBookRepo) IBookUsecase {
	return &BookUsecase{bookRepo: bookRepo}
}
