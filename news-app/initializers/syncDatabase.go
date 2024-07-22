package initializers

import (
	"github.com/PhongVoyager/news-project/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
