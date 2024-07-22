package handlers

import (
	"net/http"

	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	payload := models.UserParams{}
	if c.ShouldBind(&payload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	// Handle file upload for the avatar
	file, err := c.FormFile("avatar")
	if err == nil {
		// uploadedImageUrl, uploadErr := middleware.HandleFileUploadDisk(c, file)
		uploadedImageUrl, uploadErr := middleware.HandleFileUploadS3(file)
		if uploadErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uploadErr.Error()})
			return
		}
		// Update the payload with the path or reference to the saved file
		payload.Avatar = uploadedImageUrl
	} else if err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userProfile, err := h.userUsecase.UpdateProfile(c, &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userProfile)
}
