// Code generated by MockGen. DO NOT EDIT.
// Source: internal/interface/validation/auth.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	user "github.com/calmato/gran-book/api/server/user/proto/service/user"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthRequestValidation is a mock of AuthRequestValidation interface.
type MockAuthRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRequestValidationMockRecorder
}

// MockAuthRequestValidationMockRecorder is the mock recorder for MockAuthRequestValidation.
type MockAuthRequestValidationMockRecorder struct {
	mock *MockAuthRequestValidation
}

// NewMockAuthRequestValidation creates a new mock instance.
func NewMockAuthRequestValidation(ctrl *gomock.Controller) *MockAuthRequestValidation {
	mock := &MockAuthRequestValidation{ctrl: ctrl}
	mock.recorder = &MockAuthRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRequestValidation) EXPECT() *MockAuthRequestValidationMockRecorder {
	return m.recorder
}

// CreateAuth mocks base method.
func (m *MockAuthRequestValidation) CreateAuth(req *user.CreateAuthRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuth", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuth indicates an expected call of CreateAuth.
func (mr *MockAuthRequestValidationMockRecorder) CreateAuth(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuth", reflect.TypeOf((*MockAuthRequestValidation)(nil).CreateAuth), req)
}

// RegisterAuthDevice mocks base method.
func (m *MockAuthRequestValidation) RegisterAuthDevice(req *user.RegisterAuthDeviceRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAuthDevice", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAuthDevice indicates an expected call of RegisterAuthDevice.
func (mr *MockAuthRequestValidationMockRecorder) RegisterAuthDevice(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAuthDevice", reflect.TypeOf((*MockAuthRequestValidation)(nil).RegisterAuthDevice), req)
}

// UpdateAuthAddress mocks base method.
func (m *MockAuthRequestValidation) UpdateAuthAddress(req *user.UpdateAuthAddressRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthAddress", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthAddress indicates an expected call of UpdateAuthAddress.
func (mr *MockAuthRequestValidationMockRecorder) UpdateAuthAddress(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthAddress", reflect.TypeOf((*MockAuthRequestValidation)(nil).UpdateAuthAddress), req)
}

// UpdateAuthEmail mocks base method.
func (m *MockAuthRequestValidation) UpdateAuthEmail(req *user.UpdateAuthEmailRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthEmail", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthEmail indicates an expected call of UpdateAuthEmail.
func (mr *MockAuthRequestValidationMockRecorder) UpdateAuthEmail(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthEmail", reflect.TypeOf((*MockAuthRequestValidation)(nil).UpdateAuthEmail), req)
}

// UpdateAuthPassword mocks base method.
func (m *MockAuthRequestValidation) UpdateAuthPassword(req *user.UpdateAuthPasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthPassword", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthPassword indicates an expected call of UpdateAuthPassword.
func (mr *MockAuthRequestValidationMockRecorder) UpdateAuthPassword(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthPassword", reflect.TypeOf((*MockAuthRequestValidation)(nil).UpdateAuthPassword), req)
}

// UpdateAuthProfile mocks base method.
func (m *MockAuthRequestValidation) UpdateAuthProfile(req *user.UpdateAuthProfileRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthProfile", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthProfile indicates an expected call of UpdateAuthProfile.
func (mr *MockAuthRequestValidationMockRecorder) UpdateAuthProfile(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthProfile", reflect.TypeOf((*MockAuthRequestValidation)(nil).UpdateAuthProfile), req)
}

// UploadAuthThumbnail mocks base method.
func (m *MockAuthRequestValidation) UploadAuthThumbnail(req *user.UploadAuthThumbnailRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAuthThumbnail", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAuthThumbnail indicates an expected call of UploadAuthThumbnail.
func (mr *MockAuthRequestValidationMockRecorder) UploadAuthThumbnail(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAuthThumbnail", reflect.TypeOf((*MockAuthRequestValidation)(nil).UploadAuthThumbnail), req)
}
