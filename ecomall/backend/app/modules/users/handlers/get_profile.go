package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetProfile(c *gin.Context) {
	userProfile := h.userUsecase.GetProfile(c)
	c.JSON(http.StatusOK, models.FilterUserRecord(userProfile))
}
