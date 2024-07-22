package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *BookCategoryHandler) GetBookCategory(c *gin.Context) {
	id := c.Param("id")
	bookCategory, err := h.bookCategoryUsecase.GetBookCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.NewBookCategoryResponse(bookCategory))
}
