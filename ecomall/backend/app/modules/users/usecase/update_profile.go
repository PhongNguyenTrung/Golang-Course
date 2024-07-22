package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *UserUsecase) UpdateProfile(c *gin.Context, payload *models.UserParams) (*models.UserResponse, error) {
	user := c.MustGet("user").(*models.User)
	updatedUser, err := h.userRepo.UpdateUser(user, payload)
	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(updatedUser), nil
}
