package models

import (
	"time"

	"github.com/rosberry/go-pagination"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title          string       `gorm:"not null" json:"title"`
	Author         string       `gorm:"not null" json:"author"`
	Image          string       `json:"image"`
	PublishDate    time.Time    `json:"publish_date"`
	Description    string       `json:"description"`
	BookCategoryID uint         `json:"book_category_id"`
	BookCategory   BookCategory `gorm:"constrant:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book_category"`
	UserID         uint         `json:"user_id"`
	User           User         `gorm:"constrant:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}

type BookParams struct {
	Title          string    `gorm:"not null" form:"title" binding:"required"`
	Author         string    `gorm:"not null" form:"author" binding:"required"`
	Image          string    `file:"image" json:"image"`
	PublishDate    time.Time `form:"publish_date" time_format:"2006-01-02"`
	Description    string    `form:"description"`
	BookCategoryID uint      `form:"book_category_id" binding:"required"`
}

type BookQueryParams struct {
	Title          string `form:"title" json:"title,omitempty"`
	Author         string `form:"author" json:"author,omitempty"`
	BookCategoryID uint   `form:"book_category_id" json:"book_category_id,omitempty"`
}

type BookResponse struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Image            string `json:"image"`
	PublishDate      string `json:"publish_date"`
	Description      string `json:"description"`
	BookCategoryName string `json:"book_category_name"`
}

type BookDetailResponse struct {
	BookResponse
	BookCategory BookCategoryResponse `json:"book_category"`
}

type BooksListResponse struct {
	Data       []*BookResponse      `json:"data"`
	Pagination *pagination.PageInfo `json:"pagination"`
}

func NewBookResponse(book *Book) *BookResponse {
	return &BookResponse{
		ID:               book.ID,
		Title:            book.Title,
		Author:           book.Author,
		Image:            book.Image,
		PublishDate:      book.PublishDate.Format("2006-01-02"),
		Description:      book.Description,
		BookCategoryName: book.BookCategory.Name,
	}
}

func NewBookDetailResponse(book *Book) *BookDetailResponse {
	return &BookDetailResponse{
		BookResponse: *NewBookResponse(book),
		BookCategory: *NewBookCategoryResponse(&book.BookCategory),
	}
}

func NewBooksReponse(books []*Book, paginator *pagination.Paginator) BooksListResponse {
	booksResponse := make([]*BookResponse, 0, len(books))

	for _, book := range books {
		booksResponse = append(booksResponse, NewBookResponse(book))
	}
	return BooksListResponse{
		Data:       booksResponse,
		Pagination: paginator.PageInfo,
	}
}
