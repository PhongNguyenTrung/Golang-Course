package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rosberry/go-pagination"
	"gorm.io/gorm"
)

func InitPaginator(c *gin.Context, db *gorm.DB, model interface{}, limit uint) *pagination.Paginator {
	if limit == 0 {
		limit = 5
	}

	paginator, err := pagination.New(pagination.Options{
		GinContext: c,
		DB:         db,
		Model:      model,
		Limit:      limit,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return nil
	}

	return paginator
}
