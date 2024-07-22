package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BookCategoryHandler) DeleteBookCategory(c *gin.Context) {
	bookCategoryID := c.Param("id")
	err := h.bookCategoryUsecase.DeleteBookCategory(bookCategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully!"})
}
