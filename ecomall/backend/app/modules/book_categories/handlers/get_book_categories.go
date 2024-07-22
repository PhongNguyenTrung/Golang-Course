package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *BookCategoryHandler) GetBookCategories(c *gin.Context) {
	bookCategories, err := h.bookCategoryUsecase.GetBookCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.NewBookCategoriesResponse(bookCategories))
}
