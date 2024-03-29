// Code generated by MockGen. DO NOT EDIT.
// Source: validation.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	reflect "reflect"

	chat "github.com/calmato/gran-book/api/internal/user/domain/chat"
	gomock "github.com/golang/mock/gomock"
)

// MockValidation is a mock of Validation interface.
type MockValidation struct {
	ctrl     *gomock.Controller
	recorder *MockValidationMockRecorder
}

// MockValidationMockRecorder is the mock recorder for MockValidation.
type MockValidationMockRecorder struct {
	mock *MockValidation
}

// NewMockValidation creates a new mock instance.
func NewMockValidation(ctrl *gomock.Controller) *MockValidation {
	mock := &MockValidation{ctrl: ctrl}
	mock.recorder = &MockValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidation) EXPECT() *MockValidationMockRecorder {
	return m.recorder
}

// Message mocks base method.
func (m *MockValidation) Message(ctx context.Context, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message", ctx, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// Message indicates an expected call of Message.
func (mr *MockValidationMockRecorder) Message(ctx, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockValidation)(nil).Message), ctx, cm)
}

// Room mocks base method.
func (m *MockValidation) Room(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Room", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Room indicates an expected call of Room.
func (mr *MockValidationMockRecorder) Room(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Room", reflect.TypeOf((*MockValidation)(nil).Room), ctx, cr)
}
