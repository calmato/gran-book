// Code generated by MockGen. DO NOT EDIT.
// Source: uploader.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUploader is a mock of Uploader interface.
type MockUploader struct {
	ctrl     *gomock.Controller
	recorder *MockUploaderMockRecorder
}

// MockUploaderMockRecorder is the mock recorder for MockUploader.
type MockUploaderMockRecorder struct {
	mock *MockUploader
}

// NewMockUploader creates a new mock instance.
func NewMockUploader(ctrl *gomock.Controller) *MockUploader {
	mock := &MockUploader{ctrl: ctrl}
	mock.recorder = &MockUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploader) EXPECT() *MockUploaderMockRecorder {
	return m.recorder
}

// Image mocks base method.
func (m *MockUploader) Image(ctx context.Context, roomID string, data []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Image", ctx, roomID, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Image indicates an expected call of Image.
func (mr *MockUploaderMockRecorder) Image(ctx, roomID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Image", reflect.TypeOf((*MockUploader)(nil).Image), ctx, roomID, data)
}
