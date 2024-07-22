package usecase

import (
	"errors"

	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (u UserUsecase) SignUp(ctx *gin.Context, payload *models.SignUpInput) (*models.UserResponse, error) {
	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	// check existing email
	existing := u.userRepo.CheckEmailExisting(payload.Email)

	if existing {
		return nil, errors.New("email existing, please choose another email")
	}

	createdUser, err := u.userRepo.CreateUser(payload)

	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(createdUser), nil
}
