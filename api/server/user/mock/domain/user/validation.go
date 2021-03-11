// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/user/validation.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	user "github.com/calmato/gran-book/api/server/user/internal/domain/user"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValidation is a mock of Validation interface
type MockValidation struct {
	ctrl     *gomock.Controller
	recorder *MockValidationMockRecorder
}

// MockValidationMockRecorder is the mock recorder for MockValidation
type MockValidationMockRecorder struct {
	mock *MockValidation
}

// NewMockValidation creates a new mock instance
func NewMockValidation(ctrl *gomock.Controller) *MockValidation {
	mock := &MockValidation{ctrl: ctrl}
	mock.recorder = &MockValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidation) EXPECT() *MockValidationMockRecorder {
	return m.recorder
}

// User mocks base method
func (m *MockValidation) User(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// User indicates an expected call of User
func (mr *MockValidationMockRecorder) User(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockValidation)(nil).User), ctx, u)
}

// Relationship mocks base method
func (m *MockValidation) Relationship(ctx context.Context, r *user.Relationship) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relationship", ctx, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Relationship indicates an expected call of Relationship
func (mr *MockValidationMockRecorder) Relationship(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relationship", reflect.TypeOf((*MockValidation)(nil).Relationship), ctx, r)
}
