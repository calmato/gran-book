// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/validation/user.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	input "github.com/calmato/gran-book/api/server/user/internal/application/input"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserRequestValidation is a mock of UserRequestValidation interface
type MockUserRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockUserRequestValidationMockRecorder
}

// MockUserRequestValidationMockRecorder is the mock recorder for MockUserRequestValidation
type MockUserRequestValidationMockRecorder struct {
	mock *MockUserRequestValidation
}

// NewMockUserRequestValidation creates a new mock instance
func NewMockUserRequestValidation(ctrl *gomock.Controller) *MockUserRequestValidation {
	mock := &MockUserRequestValidation{ctrl: ctrl}
	mock.recorder = &MockUserRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRequestValidation) EXPECT() *MockUserRequestValidationMockRecorder {
	return m.recorder
}

// ListUser mocks base method
func (m *MockUserRequestValidation) ListUser(in *input.ListUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListUser indicates an expected call of ListUser
func (mr *MockUserRequestValidationMockRecorder) ListUser(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockUserRequestValidation)(nil).ListUser), in)
}

// ListUserByUserIDs mocks base method
func (m *MockUserRequestValidation) ListUserByUserIDs(in *input.ListUserByUserIDs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserByUserIDs", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListUserByUserIDs indicates an expected call of ListUserByUserIDs
func (mr *MockUserRequestValidationMockRecorder) ListUserByUserIDs(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserByUserIDs", reflect.TypeOf((*MockUserRequestValidation)(nil).ListUserByUserIDs), in)
}

// ListFollow mocks base method
func (m *MockUserRequestValidation) ListFollow(in *input.ListFollow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollow", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFollow indicates an expected call of ListFollow
func (mr *MockUserRequestValidationMockRecorder) ListFollow(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollow", reflect.TypeOf((*MockUserRequestValidation)(nil).ListFollow), in)
}

// ListFollower mocks base method
func (m *MockUserRequestValidation) ListFollower(in *input.ListFollower) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollower", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListFollower indicates an expected call of ListFollower
func (mr *MockUserRequestValidationMockRecorder) ListFollower(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollower", reflect.TypeOf((*MockUserRequestValidation)(nil).ListFollower), in)
}

// SearchUser mocks base method
func (m *MockUserRequestValidation) SearchUser(in *input.SearchUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUser", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchUser indicates an expected call of SearchUser
func (mr *MockUserRequestValidationMockRecorder) SearchUser(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockUserRequestValidation)(nil).SearchUser), in)
}
