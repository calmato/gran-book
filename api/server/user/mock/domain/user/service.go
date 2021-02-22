// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/user/service.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	domain "github.com/calmato/gran-book/api/server/user/internal/domain"
	user "github.com/calmato/gran-book/api/server/user/internal/domain/user"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Authentication mocks base method
func (m *MockService) Authentication(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authentication", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authentication indicates an expected call of Authentication
func (mr *MockServiceMockRecorder) Authentication(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockService)(nil).Authentication), ctx)
}

// List mocks base method
func (m *MockService) List(ctx context.Context, query *domain.ListQuery) ([]*user.User, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, query)
	ret0, _ := ret[0].([]*user.User)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockServiceMockRecorder) List(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), ctx, query)
}

// Show mocks base method
func (m *MockService) Show(ctx context.Context, uid string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, uid)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show
func (mr *MockServiceMockRecorder) Show(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockService)(nil).Show), ctx, uid)
}

// Create mocks base method
func (m *MockService) Create(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockServiceMockRecorder) Create(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, u)
}

// Update mocks base method
func (m *MockService) Update(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockServiceMockRecorder) Update(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), ctx, u)
}

// UpdatePassword mocks base method
func (m *MockService) UpdatePassword(ctx context.Context, uid, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, uid, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockServiceMockRecorder) UpdatePassword(ctx, uid, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockService)(nil).UpdatePassword), ctx, uid, password)
}

// UploadThumbnail mocks base method
func (m *MockService) UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadThumbnail", ctx, uid, thumbnail)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadThumbnail indicates an expected call of UploadThumbnail
func (mr *MockServiceMockRecorder) UploadThumbnail(ctx, uid, thumbnail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadThumbnail", reflect.TypeOf((*MockService)(nil).UploadThumbnail), ctx, uid, thumbnail)
}
