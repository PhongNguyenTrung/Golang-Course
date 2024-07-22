package repository

import (
	"github.com/1rhino/clean_architecture/app/middleware"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/gin-gonic/gin"
	"github.com/rosberry/go-pagination"
	"gorm.io/gorm"
)

func (r *BookRepo) GetBooks(c *gin.Context, user *models.User, query models.BookQueryParams) ([]*models.Book, *pagination.Paginator, error) {
	paginator := middleware.InitPaginator(c, r.DB, &models.Book{}, 10)

	books := []*models.Book{}
	result := r.DB.Model(&models.Book{}).Preload("BookCategory").Where("user_id = ?", user.ID)
	result = filter_by_title(result, query.Title)
	result = filter_by_author(result, query.Author)
	result = filter_by_category(result, query.BookCategoryID)
	result = result.Order("created_at desc")
	err := paginator.Find(result, &books)

	if err != nil {
		return nil, nil, err
	}

	return books, paginator, nil
}

func filter_by_title(tx *gorm.DB, title string) *gorm.DB {
	if title == "" {
		return tx
	}

	return tx.Where("title LIKE ?", "%"+title+"%")
}

func filter_by_author(tx *gorm.DB, author string) *gorm.DB {
	if author == "" {
		return tx
	}

	return tx.Where("author LIKE ?", "%"+author+"%")
}

func filter_by_category(tx *gorm.DB, bookCategoryId uint) *gorm.DB {
	if bookCategoryId == 0 {
		return tx
	}

	return tx.Where("book_category_id = ?", bookCategoryId)
}
