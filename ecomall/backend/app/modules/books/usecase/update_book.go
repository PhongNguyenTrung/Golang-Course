package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (u *BookUsecase) UpdateBook(c *gin.Context, payload models.BookParams) (*models.Book, error) {
	id := c.Param("id")
	user := c.MustGet("user").(*models.User)
	return u.bookRepo.UpdateBook(user, id, payload)
}
