// Code generated by MockGen. DO NOT EDIT.
// Source: internal/interface/validation/admin.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	proto "github.com/calmato/gran-book/api/server/user/proto"
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
func (m *MockAdminRequestValidation) CreateAdmin(req *proto.CreateAdminRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockAdminRequestValidationMockRecorder) CreateAdmin(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminRequestValidation)(nil).CreateAdmin), req)
}

// DeleteAdmin mocks base method.
func (m *MockAdminRequestValidation) DeleteAdmin(req *proto.DeleteAdminRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockAdminRequestValidationMockRecorder) DeleteAdmin(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockAdminRequestValidation)(nil).DeleteAdmin), req)
}

// GetAdmin mocks base method.
func (m *MockAdminRequestValidation) GetAdmin(req *proto.GetAdminRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockAdminRequestValidationMockRecorder) GetAdmin(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockAdminRequestValidation)(nil).GetAdmin), req)
}

// ListAdmin mocks base method.
func (m *MockAdminRequestValidation) ListAdmin(req *proto.ListAdminRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdmin", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListAdmin indicates an expected call of ListAdmin.
func (mr *MockAdminRequestValidationMockRecorder) ListAdmin(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdmin", reflect.TypeOf((*MockAdminRequestValidation)(nil).ListAdmin), req)
}

// UpdateAdminContact mocks base method.
func (m *MockAdminRequestValidation) UpdateAdminContact(req *proto.UpdateAdminContactRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminContact", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminContact indicates an expected call of UpdateAdminContact.
func (mr *MockAdminRequestValidationMockRecorder) UpdateAdminContact(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminContact", reflect.TypeOf((*MockAdminRequestValidation)(nil).UpdateAdminContact), req)
}

// UpdateAdminPassword mocks base method.
func (m *MockAdminRequestValidation) UpdateAdminPassword(req *proto.UpdateAdminPasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminPassword", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminPassword indicates an expected call of UpdateAdminPassword.
func (mr *MockAdminRequestValidationMockRecorder) UpdateAdminPassword(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminPassword", reflect.TypeOf((*MockAdminRequestValidation)(nil).UpdateAdminPassword), req)
}

// UpdateAdminProfile mocks base method.
func (m *MockAdminRequestValidation) UpdateAdminProfile(req *proto.UpdateAdminProfileRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminProfile", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminProfile indicates an expected call of UpdateAdminProfile.
func (mr *MockAdminRequestValidationMockRecorder) UpdateAdminProfile(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminProfile", reflect.TypeOf((*MockAdminRequestValidation)(nil).UpdateAdminProfile), req)
}

// UploadAdminThumbnail mocks base method.
func (m *MockAdminRequestValidation) UploadAdminThumbnail(req *proto.UploadAdminThumbnailRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAdminThumbnail", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAdminThumbnail indicates an expected call of UploadAdminThumbnail.
func (mr *MockAdminRequestValidationMockRecorder) UploadAdminThumbnail(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAdminThumbnail", reflect.TypeOf((*MockAdminRequestValidation)(nil).UploadAdminThumbnail), req)
}