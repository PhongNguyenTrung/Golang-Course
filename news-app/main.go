package main

import (
	"github.com/PhongVoyager/news-project/controllers"
	"github.com/PhongVoyager/news-project/initializers"
	"github.com/PhongVoyager/news-project/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDataBase()
	initializers.SyncDatabase()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/get-profile", middleware.RequireAuth, controllers.GetCurrentUser)

	return r
}

func main() {
	r := SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
