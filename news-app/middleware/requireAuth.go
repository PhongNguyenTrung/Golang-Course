package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PhongVoyager/news-project/initializers"
	"github.com/PhongVoyager/news-project/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		var user models.User
		initializers.DB.First(&user, claims["user_id"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User Not found"})
		}

		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println(err)
	}
}
