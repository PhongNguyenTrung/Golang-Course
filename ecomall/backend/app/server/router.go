package server

import (
	userHandler "github.com/1rhino/clean_architecture/app/modules/users/handlers"
	userRepository "github.com/1rhino/clean_architecture/app/modules/users/repositories"
	userUseCase "github.com/1rhino/clean_architecture/app/modules/users/usecase"
	"github.com/gin-gonic/gin"

	bookCategoryHandler "github.com/1rhino/clean_architecture/app/modules/book_categories/handlers"
	bookCategoryRepository "github.com/1rhino/clean_architecture/app/modules/book_categories/repositories"
	bookCategoryUsecase "github.com/1rhino/clean_architecture/app/modules/book_categories/usecase"

	bookHandler "github.com/1rhino/clean_architecture/app/modules/books/handlers"
	bookRepository "github.com/1rhino/clean_architecture/app/modules/books/repositories"
	bookUseCase "github.com/1rhino/clean_architecture/app/modules/books/usecase"
)

func SetupRoutes(server *Server) {
	api := server.Gin.Group("/api/v1")
	api.Static("/uploads", "./uploads")
	SetUserRoutes(server, api)
	SetBookCategoryRoutes(server, api)
	SetBookRoutes(server, api)
}

func SetUserRoutes(server *Server, api *gin.RouterGroup) {
	userRepo := userRepository.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := userHandler.NewUserHandler(userUseCase)

	api.POST("/signup", userHandler.SignUp)
	api.POST("/signin", userHandler.SignIn)

	user := api.Group("/user")
	user.Use(userHandler.Authenticate)
	user.GET("/profile", userHandler.GetProfile)
	user.PUT("/profile", userHandler.UpdateProfile)
	user.DELETE("/profile", userHandler.DeleteProfile)
	user.DELETE("/signout", userHandler.SignOut)
}

func SetBookCategoryRoutes(server *Server, api *gin.RouterGroup) {
	userRepo := userRepository.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := userHandler.NewUserHandler(userUseCase)

	bookCategoryRepo := bookCategoryRepository.NewBookCategoryRepo(server.DB)
	bookCategoryUsecase := bookCategoryUsecase.NewBookCategoryUsecase(bookCategoryRepo)
	bookCategoryHandler := bookCategoryHandler.NewBookCategoryHandler(bookCategoryUsecase)

	bookCategory := api.Group("/book_categories")
	bookCategory.Use(userHandler.Authenticate)

	bookCategory.GET("", bookCategoryHandler.GetBookCategories)
	bookCategory.POST("", bookCategoryHandler.CreateBookCategory)
	bookCategory.GET("/:id", bookCategoryHandler.GetBookCategory)
	bookCategory.PUT("/:id", bookCategoryHandler.UpdateBookCategory)
	bookCategory.DELETE("/:id", bookCategoryHandler.DeleteBookCategory)
}

func SetBookRoutes(server *Server, api *gin.RouterGroup) {
	userRepo := userRepository.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := userHandler.NewUserHandler(userUseCase)

	bookRepo := bookRepository.NewBookRepo(server.DB)
	bookUseCase := bookUseCase.NewBookUsecase(bookRepo)
	bookHandler := bookHandler.NewBookHandler(bookUseCase)

	book := api.Group("/books")
	book.Use(userHandler.Authenticate)

	book.GET("", bookHandler.GetBooks)
	book.POST("", bookHandler.CreateBook)
	book.GET("/:id", bookHandler.GetBook)
	book.PUT("/:id", bookHandler.UpdateBook)
	book.DELETE("/:id", bookHandler.DeleteBook)
}
