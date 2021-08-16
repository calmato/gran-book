// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application//chat.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	context "context"
	reflect "reflect"

	chat "github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	firestore "github.com/calmato/gran-book/api/server/user/pkg/firebase/firestore"
	gomock "github.com/golang/mock/gomock"
)

// MockChatApplication is a mock of ChatApplication interface.
type MockChatApplication struct {
	ctrl     *gomock.Controller
	recorder *MockChatApplicationMockRecorder
}

// MockChatApplicationMockRecorder is the mock recorder for MockChatApplication.
type MockChatApplicationMockRecorder struct {
	mock *MockChatApplication
}

// NewMockChatApplication creates a new mock instance.
func NewMockChatApplication(ctrl *gomock.Controller) *MockChatApplication {
	mock := &MockChatApplication{ctrl: ctrl}
	mock.recorder = &MockChatApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatApplication) EXPECT() *MockChatApplicationMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockChatApplication) CreateMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", ctx, cr, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockChatApplicationMockRecorder) CreateMessage(ctx, cr, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockChatApplication)(nil).CreateMessage), ctx, cr, cm)
}

// CreateRoom mocks base method.
func (m *MockChatApplication) CreateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockChatApplicationMockRecorder) CreateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockChatApplication)(nil).CreateRoom), ctx, cr)
}

// GetRoom mocks base method.
func (m *MockChatApplication) GetRoom(ctx context.Context, roomID, userID string) (*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, roomID, userID)
	ret0, _ := ret[0].(*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockChatApplicationMockRecorder) GetRoom(ctx, roomID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockChatApplication)(nil).GetRoom), ctx, roomID, userID)
}

// ListRoom mocks base method.
func (m *MockChatApplication) ListRoom(ctx context.Context, userID string, p *firestore.Params) ([]*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", ctx, userID, p)
	ret0, _ := ret[0].([]*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockChatApplicationMockRecorder) ListRoom(ctx, userID, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockChatApplication)(nil).ListRoom), ctx, userID, p)
}

// UploadImage mocks base method.
func (m *MockChatApplication) UploadImage(ctx context.Context, cr *chat.Room, image []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadImage", ctx, cr, image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockChatApplicationMockRecorder) UploadImage(ctx, cr, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockChatApplication)(nil).UploadImage), ctx, cr, image)
}
