package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *BookCategoryHandler) CreateBookCategory(c *gin.Context) {
	var bookCategory models.BookCategoryParams

	err := c.ShouldBind(&bookCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err == nil {
		uploadedImageUrl, uploadErr := middleware.HandleFileUploadS3(file)
		if uploadErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uploadErr.Error()})
			return
		}
		// Update the payload with the path or reference to the saved file
		bookCategory.Image = uploadedImageUrl
	} else if err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBookCategory, err := h.bookCategoryUsecase.CreateBookCategory(&bookCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.NewBookCategoryResponse(createdBookCategory))
}
