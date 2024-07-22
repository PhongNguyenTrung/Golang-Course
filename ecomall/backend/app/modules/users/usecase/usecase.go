package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	users "github.com/1rhino/clean_architecture/app/modules/users/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type IUserUsecase interface {
	SignUp(ctx *gin.Context, payload *models.SignUpInput) (*models.UserResponse, error)
	SignIn(ctx *gin.Context, payload *models.SignInInput) error
	Authenticate(ctx *gin.Context, token *jwt.Token) error
	GetProfile(ctx *gin.Context) *models.User
	UpdateProfile(ctx *gin.Context, payload *models.UserParams) (*models.UserResponse, error)
	DeleteProfile(ctx *gin.Context) error
}

type UserUsecase struct {
	userRepo users.IUserRepo
}

func NewUserUseCase(userRepo users.IUserRepo) IUserUsecase {
	return &UserUsecase{userRepo: userRepo}
}
