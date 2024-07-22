package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := h.bookUsecase.DeleteBook(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
