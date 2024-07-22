package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (u *BookUsecase) CreateBook(c *gin.Context, payload models.BookParams) (*models.Book, error) {
	user := c.MustGet("user").(*models.User)
	return u.bookRepo.CreateBook(user, payload)
}
