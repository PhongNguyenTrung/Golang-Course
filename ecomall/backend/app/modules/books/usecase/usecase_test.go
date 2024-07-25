package usecase_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mocks "github.com/1rhino/clean_architecture/app/mocks/books"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/1rhino/clean_architecture/app/modules/books/usecase"
	"github.com/gin-gonic/gin"
	"github.com/rosberry/go-pagination"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookRepo := mocks.NewMockIBookRepo(ctrl)
	bookUsecase := usecase.NewBookUsecase(mockBookRepo)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		payload := models.BookParams{
			Title:  "Book Title",
			Author: "Author Name",
		}

		createdBook := &models.Book{
			Title:  payload.Title,
			Author: payload.Author,
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user", user)

		mockBookRepo.EXPECT().CreateBook(user, payload).Return(createdBook, nil)

		result, err := bookUsecase.CreateBook(c, payload)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, createdBook.Title, result.Title)
		assert.Equal(t, createdBook.Author, result.Author)
	})

	t.Run("user not set in context", func(t *testing.T) {
		payload := models.BookParams{
			Title:  "Book Title",
			Author: "Author Name",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		assert.Panics(t, func() { _, _ = bookUsecase.CreateBook(c, payload) })
	})

	t.Run("create book failed", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		payload := models.BookParams{
			Title:  "Book Title",
			Author: "Author Name",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user", user)

		mockBookRepo.EXPECT().CreateBook(user, payload).Return(nil, errors.New("create book error"))

		result, err := bookUsecase.CreateBook(c, payload)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "create book error")
	})
}

func TestDeleteBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookRepo := mocks.NewMockIBookRepo(ctrl)
	bookUsecase := usecase.NewBookUsecase(mockBookRepo)

	t.Run("success", func(t *testing.T) {
		bookID := "1"

		mockBookRepo.EXPECT().DeleteBook(bookID).Return(nil)

		err := bookUsecase.DeleteBook(bookID)

		assert.Nil(t, err)
	})

	t.Run("delete book failed", func(t *testing.T) {
		bookID := "1"

		mockBookRepo.EXPECT().DeleteBook(bookID).Return(errors.New("delete failed"))

		err := bookUsecase.DeleteBook(bookID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "delete failed")
	})
}

func TestGetBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookRepo := mocks.NewMockIBookRepo(ctrl)
	bookUsecase := usecase.NewBookUsecase(mockBookRepo)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		bookID := "1"
		book := &models.Book{
			Title:  "Test Book",
			Author: "Test Author",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "id", Value: bookID}}
		c.Set("user", user)

		mockBookRepo.EXPECT().GetBook(user, bookID).Return(book, nil)

		result, err := bookUsecase.GetBook(c)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, book.Title, result.Title)
		assert.Equal(t, book.Author, result.Author)
	})

	t.Run("user not set in context", func(t *testing.T) {
		bookID := "1"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "id", Value: bookID}}

		assert.Panics(t, func() { _, _ = bookUsecase.GetBook(c) })
	})

	t.Run("get book failed", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		bookID := "1"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "id", Value: bookID}}
		c.Set("user", user)

		mockBookRepo.EXPECT().GetBook(user, bookID).Return(nil, errors.New("get book error"))

		result, err := bookUsecase.GetBook(c)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "get book error")
	})
}

func TestGetBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookRepo := mocks.NewMockIBookRepo(ctrl)
	bookUsecase := usecase.NewBookUsecase(mockBookRepo)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		books := []*models.Book{
			{Title: "Test Book 1", Author: "Test Author 1"},
			{Title: "Test Book 2", Author: "Test Author 2"},
		}

		booksParams := models.BookQueryParams{
			Title:  "Test Book",
			Author: "Test Author",
		}

		paginator := &pagination.Paginator{}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = BooksListRequest(booksParams)
		c.Set("user", user)

		mockBookRepo.EXPECT().GetBooks(c, user, booksParams).Return(books, paginator, nil)

		result, pag, err := bookUsecase.GetBooks(c)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, pag)
		assert.Equal(t, books, result)
		assert.Equal(t, paginator, pag)
	})

	t.Run("user not set in context", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		assert.Panics(t, func() { _, _, _ = bookUsecase.GetBooks(c) })
	})

	t.Run("bind query params failed", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user", user)
		c.Request = BooksListRequest(models.BookQueryParams{Title: "Test Book", Author: "Test Author"})

		expectedParams := models.BookQueryParams{
			Title:  "Test Book",
			Author: "Test Author",
		}

		mockBookRepo.EXPECT().GetBooks(c, user, expectedParams).Return(nil, nil, errors.New("get books error"))
		_, _, err := bookUsecase.GetBooks(c)

		assert.NotNil(t, err)
	})

	t.Run("get books failed", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user", user)
		c.Request = BooksListRequest(models.BookQueryParams{Title: "Test Book", Author: "Test Author"})

		mockBookRepo.EXPECT().GetBooks(c, user, models.BookQueryParams{Title: "Test Book", Author: "Test Author"}).Return(nil, nil, errors.New("get books error"))

		result, pag, err := bookUsecase.GetBooks(c)

		assert.Nil(t, result)
		assert.Nil(t, pag)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "get books error")
	})
}

func TestUpdateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookRepo := mocks.NewMockIBookRepo(ctrl)
	bookUsecase := usecase.NewBookUsecase(mockBookRepo)

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		bookID := "1"
		payload := models.BookParams{
			Title:  "Updated Book Title",
			Author: "Updated Author Name",
		}

		updatedBook := &models.Book{
			Title:  payload.Title,
			Author: payload.Author,
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "id", Value: bookID}}
		c.Set("user", user)

		mockBookRepo.EXPECT().UpdateBook(user, bookID, payload).Return(updatedBook, nil)

		result, err := bookUsecase.UpdateBook(c, payload)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedBook.Title, result.Title)
		assert.Equal(t, updatedBook.Author, result.Author)
	})

	t.Run("user not set in context", func(t *testing.T) {
		bookID := "1"
		payload := models.BookParams{
			Title:  "Updated Book Title",
			Author: "Updated Author Name",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "id", Value: bookID}}

		assert.Panics(t, func() { _, _ = bookUsecase.UpdateBook(c, payload) })
	})

	t.Run("update book failed", func(t *testing.T) {
		user := &models.User{
			Email: "test@example.com",
			Name:  "John Doe",
		}

		bookID := "1"
		payload := models.BookParams{
			Title:  "Updated Book Title",
			Author: "Updated Author Name",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{gin.Param{Key: "id", Value: bookID}}
		c.Set("user", user)

		mockBookRepo.EXPECT().UpdateBook(user, bookID, payload).Return(nil, errors.New("update book error"))

		result, err := bookUsecase.UpdateBook(c, payload)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "update book error")
	})
}

func BooksListRequest(bookQueryParams models.BookQueryParams) *http.Request {
	query := url.Values{}
	query.Set("title", bookQueryParams.Title)
	query.Set("author", bookQueryParams.Author)

	return httptest.NewRequest("GET", "/books?"+query.Encode(), nil)
}
