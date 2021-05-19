// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/auth/service.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Authentication mocks base method.
func (m *MockService) Authentication(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authentication", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authentication indicates an expected call of Authentication.
func (mr *MockServiceMockRecorder) Authentication(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockService)(nil).Authentication), ctx)
}
