package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *BookCategoryHandler) UpdateBookCategory(c *gin.Context) {
	bookCategoryID := c.Param("id")
	payload := &models.BookCategoryParams{}

	if err := c.ShouldBind(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		if err != http.ErrMissingFile {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		uploadedImageUrl, uploadErr := middleware.HandleFileUploadS3(file)
		if uploadErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uploadErr.Error()})
			return
		}
		payload.Image = uploadedImageUrl
	}

	updatedBookCategory, err := h.bookCategoryUsecase.UpdateBookCategory(bookCategoryID, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.NewBookCategoryResponse(updatedBookCategory))
}
