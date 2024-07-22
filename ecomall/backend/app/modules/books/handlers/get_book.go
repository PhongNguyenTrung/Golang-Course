package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *BookHandler) GetBook(c *gin.Context) {
	book, err := h.bookUsecase.GetBook(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.NewBookDetailResponse(book))
}
