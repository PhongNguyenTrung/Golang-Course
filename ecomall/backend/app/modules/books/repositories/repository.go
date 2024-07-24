package repository

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
	"github.com/rosberry/go-pagination"
	"gorm.io/gorm"
)

type IBookRepo interface {
	CreateBook(user *models.User, payload models.BookParams) (*models.Book, error)
	GetBooks(c *gin.Context, user *models.User, query models.BookQueryParams) ([]*models.Book, *pagination.Paginator, error)
	GetBook(user *models.User, id string) (*models.Book, error)
	UpdateBook(user *models.User, id string, payload models.BookParams) (*models.Book, error)
	DeleteBook(id string) error
}

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) IBookRepo {
	return &BookRepo{DB: db}
}
