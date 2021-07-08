// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/validation/chat.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	input "github.com/calmato/gran-book/api/server/user/internal/application/input"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChatRequestValidation is a mock of ChatRequestValidation interface
type MockChatRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockChatRequestValidationMockRecorder
}

// MockChatRequestValidationMockRecorder is the mock recorder for MockChatRequestValidation
type MockChatRequestValidationMockRecorder struct {
	mock *MockChatRequestValidation
}

// NewMockChatRequestValidation creates a new mock instance
func NewMockChatRequestValidation(ctrl *gomock.Controller) *MockChatRequestValidation {
	mock := &MockChatRequestValidation{ctrl: ctrl}
	mock.recorder = &MockChatRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatRequestValidation) EXPECT() *MockChatRequestValidationMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method
func (m *MockChatRequestValidation) CreateRoom(in *input.CreateRoom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom
func (mr *MockChatRequestValidationMockRecorder) CreateRoom(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockChatRequestValidation)(nil).CreateRoom), in)
}