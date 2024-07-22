package models

import "gorm.io/gorm"

type BookCategory struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100); not null" json:"name" binding:"required,min=2,max=100"`
	Image       string `gorm:"type:text" json:"image" binding:"url"`
	Description string `gorm:"type:text" json:"description"`
	Books       []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BookCategoryParams struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Image       string `file:"image" json:"image"`
}

type BookCategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func NewBookCategoryResponse(bookCategory *BookCategory) *BookCategoryResponse {
	return &BookCategoryResponse{
		ID:          bookCategory.ID,
		Name:        bookCategory.Name,
		Image:       bookCategory.Image,
		Description: bookCategory.Description,
	}
}

func NewBookCategoriesResponse(bookCategories []*BookCategory) []*BookCategoryResponse {
	var bookCategoriesResponse []*BookCategoryResponse
	for _, bookCategory := range bookCategories {
		bookCategoriesResponse = append(bookCategoriesResponse, NewBookCategoryResponse(bookCategory))
	}
	return bookCategoriesResponse
}

func (BookCategory) TableName() string {
	return "book_categories"
}
