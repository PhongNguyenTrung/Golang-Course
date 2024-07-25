package usecase_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	mocks "github.com/1rhino/clean_architecture/app/mocks/users"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/1rhino/clean_architecture/app/modules/users/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(ctrl)
	userUsecase := usecase.NewUserUseCase(mockUserRepo)

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)

	t.Run("token expired", func(t *testing.T) {
		expiredToken := createToken(time.Now().Add(-time.Hour), 1)
		err := userUsecase.Authenticate(c, expiredToken)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "token expired")
	})

	t.Run("parse token claims failed", func(t *testing.T) {
		invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{})

		err := userUsecase.Authenticate(c, invalidToken)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "parse token claims failed")
	})

	t.Run("user not found", func(t *testing.T) {
		validToken := createToken(time.Now().Add(time.Hour), 1)

		mockUserRepo.EXPECT().FindUserById(uint(1)).Return(nil, errors.New("user not found"))

		err := userUsecase.Authenticate(c, validToken)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "user not found")
	})

	t.Run("success", func(t *testing.T) {
		validToken := createToken(time.Now().Add(time.Hour), 1)

		user := &models.User{
			Name: "John Doe",
		}

		mockUserRepo.EXPECT().FindUserById(uint(1)).Return(user, nil)

		err := userUsecase.Authenticate(c, validToken)

		assert.Nil(t, err)
		value, exists := c.Get("user")
		assert.True(t, exists)
		assert.Equal(t, user, value)
	})
}

func TestDeleteProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(ctrl)
	userUsecase := usecase.NewUserUseCase(mockUserRepo)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			Name: "John Doe",
		}

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Set("user", user)

		mockUserRepo.EXPECT().DeleteUser(user).Return(nil)

		err := userUsecase.DeleteProfile(c)

		assert.Nil(t, err)

		cookies := c.Writer.Header().Values("Set-Cookie")
		assert.NotEmpty(t, cookies)
		assert.True(t, strings.Contains(cookies[0], "Authorization=;"))
		assert.True(t, strings.Contains(cookies[0], "Max-Age=0"))
		assert.True(t, strings.Contains(cookies[0], "Path=/"))
		assert.True(t, strings.Contains(cookies[0], "HttpOnly"))
		assert.True(t, strings.Contains(cookies[0], "Secure"))
	})

	t.Run("failure", func(t *testing.T) {
		user := &models.User{
			Name: "John Doe",
		}

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Set("user", user)

		mockUserRepo.EXPECT().DeleteUser(user).Return(errors.New("delete failed"))

		err := userUsecase.DeleteProfile(c)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "delete failed")

		cookies := c.Writer.Header().Values("Set-Cookie")
		assert.Empty(t, cookies)
	})
}

func TestGetProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			Name: "John Doe",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user", user)

		userUsecase := usecase.UserUsecase{}
		result := userUsecase.GetProfile(c)

		assert.NotNil(t, result)
		assert.Equal(t, user, result)
	})

	t.Run("user not set in context", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		userUsecase := usecase.UserUsecase{}
		assert.Panics(t, func() { userUsecase.GetProfile(c) })
	})
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(ctrl)
	userUsecase := usecase.NewUserUseCase(mockUserRepo)

	gin.SetMode(gin.TestMode)

	t.Run("invalid email", func(t *testing.T) {
		payload := &models.SignInInput{
			Email:    "invalid@example.com",
			Password: "password123",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockUserRepo.EXPECT().FindUserByEmail(payload.Email).Return(nil, errors.New("invalid Email"))

		err := userUsecase.SignIn(c, payload)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "invalid Email")
	})

	t.Run("invalid password", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := &models.User{
			Email:    "test@example.com",
			Password: string(hashedPassword),
		}

		payload := &models.SignInInput{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockUserRepo.EXPECT().FindUserByEmail(payload.Email).Return(user, nil)

		err := userUsecase.SignIn(c, payload)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "invalid password")
	})

	t.Run("success", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := &models.User{
			Email:    "test@example.com",
			Password: string(hashedPassword),
		}

		payload := &models.SignInInput{
			Email:    "test@example.com",
			Password: "password123",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockUserRepo.EXPECT().FindUserByEmail(payload.Email).Return(user, nil)

		err := userUsecase.SignIn(c, payload)

		assert.Nil(t, err)

		cookies := w.Result().Cookies()
		assert.NotEmpty(t, cookies)

		var authCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "Authorization" {
				authCookie = cookie
				break
			}
		}

		assert.NotNil(t, authCookie)
		assert.Equal(t, "/", authCookie.Path)
		assert.True(t, authCookie.HttpOnly)
		assert.True(t, authCookie.Secure)

		// Verify the token
		token, err := jwt.Parse(authCookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		assert.Nil(t, err)
		claims, ok := token.Claims.(jwt.MapClaims)
		assert.True(t, ok)
		assert.Equal(t, float64(user.ID), claims["user_id"].(float64))
	})
}

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(ctrl)
	userUsecase := usecase.NewUserUseCase(mockUserRepo)

	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(nil)

	t.Run("passwords do not match", func(t *testing.T) {
		payload := &models.SignUpInput{
			Email:           "test@example.com",
			Password:        "password123",
			PasswordConfirm: "password321",
		}

		result, err := userUsecase.SignUp(ctx, payload)

		assert.Nil(t, result)
		assert.EqualError(t, err, "passwords do not match")
	})

	t.Run("email already existing", func(t *testing.T) {
		payload := &models.SignUpInput{
			Email:           "test@example.com",
			Password:        "password123",
			PasswordConfirm: "password123",
		}

		mockUserRepo.EXPECT().CheckEmailExisting(payload.Email).Return(true)

		result, err := userUsecase.SignUp(ctx, payload)

		assert.Nil(t, result)
		assert.EqualError(t, err, "email existing, please choose another email")
	})

	t.Run("create user failed", func(t *testing.T) {
		payload := &models.SignUpInput{
			Email:           "test@example.com",
			Password:        "password123",
			PasswordConfirm: "password123",
		}

		mockUserRepo.EXPECT().CheckEmailExisting(payload.Email).Return(false)
		mockUserRepo.EXPECT().CreateUser(payload).Return(nil, errors.New("create user error"))

		result, err := userUsecase.SignUp(ctx, payload)

		assert.Nil(t, result)
		assert.EqualError(t, err, "create user error")
	})

	t.Run("success", func(t *testing.T) {
		payload := &models.SignUpInput{
			Email:           "test@example.com",
			Password:        "password123",
			PasswordConfirm: "password123",
		}

		createdUser := &models.User{
			Email: payload.Email,
		}

		mockUserRepo.EXPECT().CheckEmailExisting(payload.Email).Return(false)
		mockUserRepo.EXPECT().CreateUser(payload).Return(createdUser, nil)

		result, err := userUsecase.SignUp(ctx, payload)

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, result.Email, createdUser.Email)
	})
}

func TestUpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepo(ctrl)
	userUsecase := usecase.NewUserUseCase(mockUserRepo)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		payload := &models.UserParams{
			Email: "new@example.com",
			Name:  "Jane Doe",
		}

		updatedUser := &models.User{
			Email: payload.Email,
			Name:  payload.Name,
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user", user)

		mockUserRepo.EXPECT().UpdateUser(user, payload).Return(updatedUser, nil)

		result, err := userUsecase.UpdateProfile(c, payload)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedUser.Email, result.Email)
		assert.Equal(t, updatedUser.Name, result.Name)
	})

	t.Run("user not set in context", func(t *testing.T) {
		payload := &models.UserParams{
			Email: "new@example.com",
			Name:  "Jane Doe",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		assert.Panics(t, func() { _, _ = userUsecase.UpdateProfile(c, payload) })
	})

	t.Run("update user failed", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		payload := &models.UserParams{
			Email: "new@example.com",
			Name:  "Jane Doe",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user", user)

		mockUserRepo.EXPECT().UpdateUser(user, payload).Return(nil, errors.New("update failed"))

		result, err := userUsecase.UpdateProfile(c, payload)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "update failed")
	})
}

func createToken(expirationTime time.Time, userID uint) *jwt.Token {
	claims := jwt.MapClaims{
		"exp":     expirationTime.Unix(),
		"user_id": float64(userID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	parsedToken, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return parsedToken
}
