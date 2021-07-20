// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/chat/service.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	domain "github.com/calmato/gran-book/api/server/user/internal/domain"
	chat "github.com/calmato/gran-book/api/server/user/internal/domain/chat"
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

// ListRoom mocks base method
func (m *MockService) ListRoom(ctx context.Context, q *domain.ListQuery, uid string) ([]*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", ctx, q, uid)
	ret0, _ := ret[0].([]*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom
func (mr *MockServiceMockRecorder) ListRoom(ctx, q, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockService)(nil).ListRoom), ctx, q, uid)
}

// CreateRoom mocks base method
func (m *MockService) CreateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom
func (mr *MockServiceMockRecorder) CreateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockService)(nil).CreateRoom), ctx, cr)
}

// ValidationRoom mocks base method
func (m *MockService) ValidationRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidationRoom indicates an expected call of ValidationRoom
func (mr *MockServiceMockRecorder) ValidationRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationRoom", reflect.TypeOf((*MockService)(nil).ValidationRoom), ctx, cr)
}

// PushCreateRoom mocks base method
func (m *MockService) PushCreateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushCreateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushCreateRoom indicates an expected call of PushCreateRoom
func (mr *MockServiceMockRecorder) PushCreateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushCreateRoom", reflect.TypeOf((*MockService)(nil).PushCreateRoom), ctx, cr)
}

// PushNewMessage mocks base method
func (m *MockService) PushNewMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushNewMessage", ctx, cr, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushNewMessage indicates an expected call of PushNewMessage
func (mr *MockServiceMockRecorder) PushNewMessage(ctx, cr, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushNewMessage", reflect.TypeOf((*MockService)(nil).PushNewMessage), ctx, cr, cm)
}
