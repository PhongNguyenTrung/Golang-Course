package handlers

import (
	user "github.com/1rhino/clean_architecture/app/modules/users/usecase"
)

type UserHandler struct {
	userUsecase user.IUserUsecase
}

func NewUserHandler(userUsecase user.IUserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}
