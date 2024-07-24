package usecase_test

import (
	"errors"
	"testing"

	mocks "github.com/1rhino/clean_architecture/app/mocks/book_categories"
	"github.com/1rhino/clean_architecture/app/models"
	"github.com/1rhino/clean_architecture/app/modules/book_categories/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateBookCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookCategoryRepo := mocks.NewMockIBookCategoryRepo(ctrl)
	bookCategoryUsecase := usecase.NewBookCategoryUsecase(mockBookCategoryRepo)

	t.Run("success", func(t *testing.T) {
		payload := &models.BookCategoryParams{
			Name: "Fiction",
		}

		createdCategory := &models.BookCategory{
			Name: "Fiction",
		}

		mockBookCategoryRepo.EXPECT().CreateBookCategory(payload).Return(createdCategory, nil).AnyTimes()

		result, err := bookCategoryUsecase.CreateBookCategory(payload)

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, result.Name, createdCategory.Name)
	})

	t.Run("Invalid Name", func(t *testing.T) {
		payload := &models.BookCategoryParams{
			Description: "Description",
		}

		mockBookCategoryRepo.EXPECT().CreateBookCategory(payload).Return(nil, errors.New("create book category error"))

		result, err := bookCategoryUsecase.CreateBookCategory(payload)

		assert.Nil(t, result)
		assert.EqualError(t, err, "create book category error")
	})

	t.Run("Invalid Name", func(t *testing.T) {
		payload := &models.BookCategoryParams{
			Name: "",
		}

		mockBookCategoryRepo.EXPECT().CreateBookCategory(payload).Return(nil, errors.New("create book category error"))

		result, err := bookCategoryUsecase.CreateBookCategory(payload)

		assert.Nil(t, result)
		assert.EqualError(t, err, "create book category error")
	})

	t.Run("Invalid Image URL", func(t *testing.T) {
		payload := &models.BookCategoryParams{
			Name:        "Fiction",
			Description: "Description",
			Image:       "invalid-url",
		}

		mockBookCategoryRepo.EXPECT().CreateBookCategory(payload).Return(nil, errors.New("create book category error"))

		result, err := bookCategoryUsecase.CreateBookCategory(payload)

		assert.Nil(t, result)
		assert.EqualError(t, err, "create book category error")
	})
}

func TestDeleteBookCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookCategoryRepo := mocks.NewMockIBookCategoryRepo(ctrl)
	bookCategoryUsecase := usecase.NewBookCategoryUsecase(mockBookCategoryRepo)

	t.Run("success", func(t *testing.T) {
		bookCategoryID := "1"

		mockBookCategoryRepo.EXPECT().DeleteBookCategory(bookCategoryID).Return(nil)

		err := bookCategoryUsecase.DeleteBookCategory(bookCategoryID)

		assert.Nil(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		bookCategoryID := "2"

		mockBookCategoryRepo.EXPECT().DeleteBookCategory(bookCategoryID).Return(errors.New("delete failed"))

		err := bookCategoryUsecase.DeleteBookCategory(bookCategoryID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "delete failed")
	})
}

func TestGetBookCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookCategoryRepo := mocks.NewMockIBookCategoryRepo(ctrl)
	bookCategoryUsecase := usecase.NewBookCategoryUsecase(mockBookCategoryRepo)

	t.Run("success", func(t *testing.T) {
		expectedCategories := []*models.BookCategory{}
		mockBookCategoryRepo.EXPECT().GetBookCategories().Return(expectedCategories, nil)

		result, err := bookCategoryUsecase.GetBookCategories()
		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, expectedCategories, result)
	})

	t.Run("failure", func(t *testing.T) {
		mockBookCategoryRepo.EXPECT().GetBookCategories().Return(nil, errors.New("repository error"))

		result, err := bookCategoryUsecase.GetBookCategories()

		assert.Nil(t, result)
		assert.EqualError(t, err, "repository error")
	})
}

func TestGetBookCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookCategoryRepo := mocks.NewMockIBookCategoryRepo(ctrl)
	bookCategoryUsecase := usecase.NewBookCategoryUsecase(mockBookCategoryRepo)

	t.Run("success", func(t *testing.T) {
		expectedCategory := &models.BookCategory{
			Name:        "Fiction",
			Description: "Description",
			Image:       "http://images.google.com",
		}

		mockBookCategoryRepo.EXPECT().GetBookCategory("1").Return(expectedCategory, nil)

		result, err := bookCategoryUsecase.GetBookCategory("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, expectedCategory, result)
	})

	t.Run("category not found", func(t *testing.T) {
		mockBookCategoryRepo.EXPECT().GetBookCategory("2").Return(nil, errors.New("category not found"))

		result, err := bookCategoryUsecase.GetBookCategory("2")

		assert.Nil(t, result)
		assert.EqualError(t, err, "category not found")
	})

	t.Run("invalid id", func(t *testing.T) {
		mockBookCategoryRepo.EXPECT().GetBookCategory("xxx").Return(nil, errors.New("invalid params"))

		result, err := bookCategoryUsecase.GetBookCategory("xxx")

		assert.Nil(t, result)
		assert.EqualError(t, err, "invalid params")
	})
}

func TestUpdateBookCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookCategoryRepo := mocks.NewMockIBookCategoryRepo(ctrl)
	bookCategoryUsecase := usecase.NewBookCategoryUsecase(mockBookCategoryRepo)

	t.Run("success", func(t *testing.T) {
		bookCategoryID := "1"
		bookCategory := &models.BookCategory{
			Name: "Fiction",
		}
		params := &models.BookCategoryParams{
			Name: "Updated Fiction",
		}
		updatedBookCategory := &models.BookCategory{
			Name: "Updated Fiction",
		}

		mockBookCategoryRepo.EXPECT().GetBookCategory(bookCategoryID).Return(bookCategory, nil)
		mockBookCategoryRepo.EXPECT().UpdateBookCategory(bookCategory, params).Return(updatedBookCategory, nil)

		result, err := bookCategoryUsecase.UpdateBookCategory(bookCategoryID, params)

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, updatedBookCategory, result)
	})

	t.Run("category not found", func(t *testing.T) {
		bookCategoryID := "2"
		params := &models.BookCategoryParams{
			Name: "Non-existing Category",
		}

		mockBookCategoryRepo.EXPECT().GetBookCategory(bookCategoryID).Return(nil, errors.New("category not found"))

		result, err := bookCategoryUsecase.UpdateBookCategory(bookCategoryID, params)

		assert.Nil(t, result)
		assert.EqualError(t, err, "category not found")
	})

	t.Run("update category failed", func(t *testing.T) {
		bookCategoryID := "3"
		bookCategory := &models.BookCategory{
			Name: "Old Category",
		}
		params := &models.BookCategoryParams{
			Name: "Failed Update",
		}

		mockBookCategoryRepo.EXPECT().GetBookCategory(bookCategoryID).Return(bookCategory, nil)
		mockBookCategoryRepo.EXPECT().UpdateBookCategory(bookCategory, params).Return(nil, errors.New("update failed"))

		result, err := bookCategoryUsecase.UpdateBookCategory(bookCategoryID, params)

		assert.Nil(t, result)
		assert.EqualError(t, err, "update failed")
	})
}
