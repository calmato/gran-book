// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/notification/service.go

// Package mock_notification is a generated GoMock package.
package mock_notification

import (
	context "context"
	reflect "reflect"

	notification "github.com/calmato/gran-book/api/server/information/internal/domain/notification"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockService) Create(ctx context.Context, n *notification.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, n)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(ctx, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, n)
}

// Delete mocks base method.
func (m *MockService) Delete(ctx context.Context, notificatinID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, notificatinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(ctx, notificatinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), ctx, notificatinID)
}

// List mocks base method.
func (m *MockService) List(ctx context.Context) ([]*notification.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]*notification.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), ctx)
}

// Show mocks base method.
func (m *MockService) Show(ctx context.Context, notificatinID int) (*notification.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, notificatinID)
	ret0, _ := ret[0].(*notification.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockServiceMockRecorder) Show(ctx, notificatinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockService)(nil).Show), ctx, notificatinID)
}

// Update mocks base method.
func (m *MockService) Update(ctx context.Context, n *notification.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, n)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(ctx, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), ctx, n)
}