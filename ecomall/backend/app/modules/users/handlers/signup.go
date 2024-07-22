package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) SignUp(c *gin.Context) {
	payload := models.SignUpInput{}

	if c.Bind(&payload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	user, err := h.userUsecase.SignUp(c, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
