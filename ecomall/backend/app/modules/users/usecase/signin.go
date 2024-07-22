package usecase

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u UserUsecase) SignIn(c *gin.Context, payload *models.SignInInput) error {
	user, err := u.userRepo.FindUserByEmail(payload.Email)

	if err != nil {
		return errors.New("invalid Email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return errors.New("failed to generate token")
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", tokenString, 60*60*24*30, "/", "", true, true)

	return nil
}
