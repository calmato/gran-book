// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/validation/admin.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	input "github.com/calmato/gran-book/api/server/user/internal/application/input"
	gomock "github.com/golang/mock/gomock"
)

// MockAdminRequestValidation is a mock of AdminRequestValidation interface.
type MockAdminRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRequestValidationMockRecorder
}

// MockAdminRequestValidationMockRecorder is the mock recorder for MockAdminRequestValidation.
type MockAdminRequestValidationMockRecorder struct {
	mock *MockAdminRequestValidation
}

// NewMockAdminRequestValidation creates a new mock instance.
func NewMockAdminRequestValidation(ctrl *gomock.Controller) *MockAdminRequestValidation {
	mock := &MockAdminRequestValidation{ctrl: ctrl}
	mock.recorder = &MockAdminRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRequestValidation) EXPECT() *MockAdminRequestValidationMockRecorder {
	return m.recorder
}

// CreateAdmin mocks base method.
func (m *MockAdminRequestValidation) CreateAdmin(in *input.CreateAdmin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockAdminRequestValidationMockRecorder) CreateAdmin(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminRequestValidation)(nil).CreateAdmin), in)
}

// ListAdmin mocks base method.
func (m *MockAdminRequestValidation) ListAdmin(in *input.ListAdmin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdmin", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAdmin indicates an expected call of ListAdmin.
func (mr *MockAdminRequestValidationMockRecorder) ListAdmin(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdmin", reflect.TypeOf((*MockAdminRequestValidation)(nil).ListAdmin), in)
}

// SearchAdmin mocks base method.
func (m *MockAdminRequestValidation) SearchAdmin(in *input.SearchAdmin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAdmin", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchAdmin indicates an expected call of SearchAdmin.
func (mr *MockAdminRequestValidationMockRecorder) SearchAdmin(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAdmin", reflect.TypeOf((*MockAdminRequestValidation)(nil).SearchAdmin), in)
}

// UpdateAdminContact mocks base method.
func (m *MockAdminRequestValidation) UpdateAdminContact(in *input.UpdateAdminContact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminContact", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminContact indicates an expected call of UpdateAdminContact.
func (mr *MockAdminRequestValidationMockRecorder) UpdateAdminContact(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminContact", reflect.TypeOf((*MockAdminRequestValidation)(nil).UpdateAdminContact), in)
}

// UpdateAdminPassword mocks base method.
func (m *MockAdminRequestValidation) UpdateAdminPassword(in *input.UpdateAdminPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminPassword", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminPassword indicates an expected call of UpdateAdminPassword.
func (mr *MockAdminRequestValidationMockRecorder) UpdateAdminPassword(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminPassword", reflect.TypeOf((*MockAdminRequestValidation)(nil).UpdateAdminPassword), in)
}

// UpdateAdminProfile mocks base method.
func (m *MockAdminRequestValidation) UpdateAdminProfile(in *input.UpdateAdminProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminProfile", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminProfile indicates an expected call of UpdateAdminProfile.
func (mr *MockAdminRequestValidationMockRecorder) UpdateAdminProfile(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminProfile", reflect.TypeOf((*MockAdminRequestValidation)(nil).UpdateAdminProfile), in)
}

// UploadAdminThumbnail mocks base method.
func (m *MockAdminRequestValidation) UploadAdminThumbnail(in *input.UploadAdminThumbnail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAdminThumbnail", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAdminThumbnail indicates an expected call of UploadAdminThumbnail.
func (mr *MockAdminRequestValidationMockRecorder) UploadAdminThumbnail(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAdminThumbnail", reflect.TypeOf((*MockAdminRequestValidation)(nil).UploadAdminThumbnail), in)
}
