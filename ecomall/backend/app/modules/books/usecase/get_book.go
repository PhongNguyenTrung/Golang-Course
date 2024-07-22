package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (u *BookUsecase) GetBook(c *gin.Context) (*models.Book, error) {
	id := c.Param("id")
	user := c.MustGet("user").(*models.User)
	return u.bookRepo.GetBook(user, id)
}
