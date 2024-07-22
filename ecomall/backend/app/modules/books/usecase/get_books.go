package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
	"github.com/rosberry/go-pagination"
)

func (u *BookUsecase) GetBooks(c *gin.Context) ([]*models.Book, *pagination.Paginator, error) {
	user := c.MustGet("user").(*models.User)
	bookParams := models.BookQueryParams{}
	err := c.ShouldBindQuery(&bookParams)
	if err != nil {
		return nil, nil, err
	}

	return u.bookRepo.GetBooks(c, user, bookParams)
}
