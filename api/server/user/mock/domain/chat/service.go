// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/chat/service.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	reflect "reflect"

	domain "github.com/calmato/gran-book/api/server/user/internal/domain"
	chat "github.com/calmato/gran-book/api/server/user/internal/domain/chat"
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

// CreateMessage mocks base method.
func (m *MockService) CreateMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", ctx, cr, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockServiceMockRecorder) CreateMessage(ctx, cr, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockService)(nil).CreateMessage), ctx, cr, cm)
}

// CreateRoom mocks base method.
func (m *MockService) CreateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockServiceMockRecorder) CreateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockService)(nil).CreateRoom), ctx, cr)
}

// GetRoom mocks base method.
func (m *MockService) GetRoom(ctx context.Context, roomID string) (*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, roomID)
	ret0, _ := ret[0].(*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockServiceMockRecorder) GetRoom(ctx, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockService)(nil).GetRoom), ctx, roomID)
}

// ListRoom mocks base method.
func (m *MockService) ListRoom(ctx context.Context, q *domain.ListQuery, uid string) ([]*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", ctx, q, uid)
	ret0, _ := ret[0].([]*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockServiceMockRecorder) ListRoom(ctx, q, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockService)(nil).ListRoom), ctx, q, uid)
}

// PushCreateRoom mocks base method.
func (m *MockService) PushCreateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushCreateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushCreateRoom indicates an expected call of PushCreateRoom.
func (mr *MockServiceMockRecorder) PushCreateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushCreateRoom", reflect.TypeOf((*MockService)(nil).PushCreateRoom), ctx, cr)
}

// PushNewMessage mocks base method.
func (m *MockService) PushNewMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushNewMessage", ctx, cr, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushNewMessage indicates an expected call of PushNewMessage.
func (mr *MockServiceMockRecorder) PushNewMessage(ctx, cr, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushNewMessage", reflect.TypeOf((*MockService)(nil).PushNewMessage), ctx, cr, cm)
}

// UploadImage mocks base method.
func (m *MockService) UploadImage(ctx context.Context, roomID string, image []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadImage", ctx, roomID, image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockServiceMockRecorder) UploadImage(ctx, roomID, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockService)(nil).UploadImage), ctx, roomID, image)
}

// ValidationMessage mocks base method.
func (m *MockService) ValidationMessage(ctx context.Context, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationMessage", ctx, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidationMessage indicates an expected call of ValidationMessage.
func (mr *MockServiceMockRecorder) ValidationMessage(ctx, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationMessage", reflect.TypeOf((*MockService)(nil).ValidationMessage), ctx, cm)
}

// ValidationRoom mocks base method.
func (m *MockService) ValidationRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidationRoom indicates an expected call of ValidationRoom.
func (mr *MockServiceMockRecorder) ValidationRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationRoom", reflect.TypeOf((*MockService)(nil).ValidationRoom), ctx, cr)
}