package usecase

import (
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
)

func (u *UserUsecase) DeleteProfile(c *gin.Context) error {
	user := c.MustGet("user").(*models.User)
	err := u.userRepo.DeleteUser(user)
	if err != nil {
		return err
	}

	// Set the authentication cookie's MaxAge to -1 to delete it
	c.SetCookie("Authorization", "", -1, "/", "", true, true)
	return nil
}
