package server

import (
	"github.com/1rhino/clean_architecture/config"
	"github.com/1rhino/clean_architecture/db"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// definition Server struct
type Server struct {
	Gin    *gin.Engine
	DB     *gorm.DB
	Config *config.Config
	S3     *s3.Client
}

// New Server function
func NewServer(cfg *config.Config) *Server {
	return &Server{
		Gin:    gin.Default(),
		DB:     db.Init(cfg),
		Config: cfg,
		S3:     config.InitS3(),
	}
}

func (server *Server) Start() error {
	// Configure CORS for Gin
	server.Gin.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS")
		c.Next()
	})

	SetupRoutes(server)
	return server.Gin.Run()
}
