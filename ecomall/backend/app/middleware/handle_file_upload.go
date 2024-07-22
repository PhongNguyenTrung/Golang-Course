package middleware

import (
	"context"
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/1rhino/clean_architecture/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

// Upload image on disk
func HandleFileUploadDisk(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// Specify the directory to save the avatar
	// Ensure this directory exists or create it during application startup
	saveDir := "./uploads/avatars/"
	fileName := filepath.Base(file.Filename)
	savePath := filepath.Join(saveDir, fileName)

	// Save the uploaded file
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return "", errors.New("failed to save avatar")
	}

	return savePath, nil
}

// Upload image to S3
func HandleFileUploadS3(file *multipart.FileHeader) (string, error) {
	f, openErr := file.Open()
	if openErr != nil {
		return "", errors.New("failed to open file")
	}
	defer f.Close()

	client := config.InitS3()

	uploader := manager.NewUploader(client)
	result, s3Err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("backend-go-s3"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if s3Err != nil {
		return "", s3Err
	}
	return result.Location, nil
}
