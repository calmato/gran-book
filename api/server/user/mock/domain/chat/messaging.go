// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/chat/messaging.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	chat "github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMessaging is a mock of Messaging interface
type MockMessaging struct {
	ctrl     *gomock.Controller
	recorder *MockMessagingMockRecorder
}

// MockMessagingMockRecorder is the mock recorder for MockMessaging
type MockMessagingMockRecorder struct {
	mock *MockMessaging
}

// NewMockMessaging creates a new mock instance
func NewMockMessaging(ctrl *gomock.Controller) *MockMessaging {
	mock := &MockMessaging{ctrl: ctrl}
	mock.recorder = &MockMessagingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessaging) EXPECT() *MockMessagingMockRecorder {
	return m.recorder
}

// PushCreateRoom mocks base method
func (m *MockMessaging) PushCreateRoom(cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushCreateRoom", cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushCreateRoom indicates an expected call of PushCreateRoom
func (mr *MockMessagingMockRecorder) PushCreateRoom(cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushCreateRoom", reflect.TypeOf((*MockMessaging)(nil).PushCreateRoom), cr)
}

// PushNewMessage mocks base method
func (m *MockMessaging) PushNewMessage(cr *chat.Room, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushNewMessage", cr, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushNewMessage indicates an expected call of PushNewMessage
func (mr *MockMessagingMockRecorder) PushNewMessage(cr, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushNewMessage", reflect.TypeOf((*MockMessaging)(nil).PushNewMessage), cr, cm)
}
