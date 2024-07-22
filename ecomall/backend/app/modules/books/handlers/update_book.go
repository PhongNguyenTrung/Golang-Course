package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var bookParams models.BookParams
	err := c.ShouldBind(&bookParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err == nil {
		uploadedImageUrl, uploadErr := middleware.HandleFileUploadS3(file)
		if uploadErr != nil {
			c.JSON(http.StatusInternalServerError, uploadErr.Error())
			return
		}
		bookParams.Image = uploadedImageUrl
	} else if err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookUsecase.UpdateBook(c, bookParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.NewBookDetailResponse(book))
}
