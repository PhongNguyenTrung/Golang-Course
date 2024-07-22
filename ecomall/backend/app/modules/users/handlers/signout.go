package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) SignOut(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Sign out successful"})
}
