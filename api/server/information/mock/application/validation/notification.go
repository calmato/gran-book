// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/validation/notification.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	input "github.com/calmato/gran-book/api/server/information/internal/application/input"
	gomock "github.com/golang/mock/gomock"
)

// MockNotificationRequestValidation is a mock of NotificationRequestValidation interface.
type MockNotificationRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationRequestValidationMockRecorder
}

// MockNotificationRequestValidationMockRecorder is the mock recorder for MockNotificationRequestValidation.
type MockNotificationRequestValidationMockRecorder struct {
	mock *MockNotificationRequestValidation
}

// NewMockNotificationRequestValidation creates a new mock instance.
func NewMockNotificationRequestValidation(ctrl *gomock.Controller) *MockNotificationRequestValidation {
	mock := &MockNotificationRequestValidation{ctrl: ctrl}
	mock.recorder = &MockNotificationRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationRequestValidation) EXPECT() *MockNotificationRequestValidationMockRecorder {
	return m.recorder
}

// CreateNotification mocks base method.
func (m *MockNotificationRequestValidation) CreateNotification(in *input.CreateNotification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotification", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNotification indicates an expected call of CreateNotification.
func (mr *MockNotificationRequestValidationMockRecorder) CreateNotification(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotification", reflect.TypeOf((*MockNotificationRequestValidation)(nil).CreateNotification), in)
}
