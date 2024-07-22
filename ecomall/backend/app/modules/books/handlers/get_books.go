package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, paginator, err := h.bookUsecase.GetBooks(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.NewBooksReponse(books, paginator))
}
