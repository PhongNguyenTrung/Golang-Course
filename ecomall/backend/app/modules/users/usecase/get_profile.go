package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *UserUsecase) GetProfile(c *gin.Context) *models.User {
	return c.MustGet("user").(*models.User)
}
