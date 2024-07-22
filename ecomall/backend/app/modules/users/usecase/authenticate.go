package usecase

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (h *UserUsecase) Authenticate(c *gin.Context, token *jwt.Token) error {
	// Check if token is valid
	// Get user ID from token
	// Get user from database
	// Set user in context
	// Return nil
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return errors.New("token expired")
		}

		userID := uint(claims["user_id"].(float64))
		user, err := h.userRepo.FindUserById(userID)

		if err != nil {
			return errors.New("user not found")
		}

		c.Set("user", user)
		return nil
	} else {
		return errors.New("parse token claims failed")
	}
}
