// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/chat/repository.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	reflect "reflect"

	chat "github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	firestore "github.com/calmato/gran-book/api/server/user/pkg/firebase/firestore"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockRepository) CreateMessage(ctx context.Context, roomID string, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", ctx, roomID, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockRepositoryMockRecorder) CreateMessage(ctx, roomID, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockRepository)(nil).CreateMessage), ctx, roomID, cm)
}

// CreateRoom mocks base method.
func (m *MockRepository) CreateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockRepositoryMockRecorder) CreateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRepository)(nil).CreateRoom), ctx, cr)
}

// GetRoom mocks base method.
func (m *MockRepository) GetRoom(ctx context.Context, roomID string) (*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, roomID)
	ret0, _ := ret[0].(*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockRepositoryMockRecorder) GetRoom(ctx, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockRepository)(nil).GetRoom), ctx, roomID)
}

// ListRoom mocks base method.
func (m *MockRepository) ListRoom(ctx context.Context, p *firestore.Params, qs []*firestore.Query) ([]*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", ctx, p, qs)
	ret0, _ := ret[0].([]*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockRepositoryMockRecorder) ListRoom(ctx, p, qs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockRepository)(nil).ListRoom), ctx, p, qs)
}

// UpdateRoom mocks base method.
func (m *MockRepository) UpdateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockRepositoryMockRecorder) UpdateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockRepository)(nil).UpdateRoom), ctx, cr)
}
