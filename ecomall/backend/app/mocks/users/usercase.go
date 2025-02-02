// Code generated by MockGen. DO NOT EDIT.
// Source: app/modules/users/usecase/usecase.go
//
// Generated by this command:
//
//	mockgen -source=app/modules/users/usecase/usecase.go -destination=app/mocks/users/user_usercase.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/1rhino/clean_architecture/app/models"
	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserUsecase is a mock of IUserUsecase interface.
type MockIUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUserUsecaseMockRecorder
}

// MockIUserUsecaseMockRecorder is the mock recorder for MockIUserUsecase.
type MockIUserUsecaseMockRecorder struct {
	mock *MockIUserUsecase
}

// NewMockIUserUsecase creates a new mock instance.
func NewMockIUserUsecase(ctrl *gomock.Controller) *MockIUserUsecase {
	mock := &MockIUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserUsecase) EXPECT() *MockIUserUsecaseMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockIUserUsecase) Authenticate(ctx *gin.Context, token *jwt.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockIUserUsecaseMockRecorder) Authenticate(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockIUserUsecase)(nil).Authenticate), ctx, token)
}

// DeleteProfile mocks base method.
func (m *MockIUserUsecase) DeleteProfile(ctx *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockIUserUsecaseMockRecorder) DeleteProfile(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockIUserUsecase)(nil).DeleteProfile), ctx)
}

// GetProfile mocks base method.
func (m *MockIUserUsecase) GetProfile(ctx *gin.Context) *models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", ctx)
	ret0, _ := ret[0].(*models.User)
	return ret0
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockIUserUsecaseMockRecorder) GetProfile(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockIUserUsecase)(nil).GetProfile), ctx)
}

// SignIn mocks base method.
func (m *MockIUserUsecase) SignIn(ctx *gin.Context, payload *models.SignInInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignIn indicates an expected call of SignIn.
func (mr *MockIUserUsecaseMockRecorder) SignIn(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockIUserUsecase)(nil).SignIn), ctx, payload)
}

// SignUp mocks base method.
func (m *MockIUserUsecase) SignUp(ctx *gin.Context, payload *models.SignUpInput) (*models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, payload)
	ret0, _ := ret[0].(*models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockIUserUsecaseMockRecorder) SignUp(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockIUserUsecase)(nil).SignUp), ctx, payload)
}

// UpdateProfile mocks base method.
func (m *MockIUserUsecase) UpdateProfile(ctx *gin.Context, payload *models.UserParams) (*models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, payload)
	ret0, _ := ret[0].(*models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockIUserUsecaseMockRecorder) UpdateProfile(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockIUserUsecase)(nil).UpdateProfile), ctx, payload)
}
