// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/validation/book.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	input "github.com/calmato/gran-book/api/server/book/internal/application/input"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBookRequestValidation is a mock of BookRequestValidation interface
type MockBookRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockBookRequestValidationMockRecorder
}

// MockBookRequestValidationMockRecorder is the mock recorder for MockBookRequestValidation
type MockBookRequestValidationMockRecorder struct {
	mock *MockBookRequestValidation
}

// NewMockBookRequestValidation creates a new mock instance
func NewMockBookRequestValidation(ctrl *gomock.Controller) *MockBookRequestValidation {
	mock := &MockBookRequestValidation{ctrl: ctrl}
	mock.recorder = &MockBookRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookRequestValidation) EXPECT() *MockBookRequestValidationMockRecorder {
	return m.recorder
}

// BookItem mocks base method
func (m *MockBookRequestValidation) BookItem(in *input.BookItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookItem", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// BookItem indicates an expected call of BookItem
func (mr *MockBookRequestValidationMockRecorder) BookItem(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookItem", reflect.TypeOf((*MockBookRequestValidation)(nil).BookItem), in)
}

// CreateBookshelf mocks base method
func (m *MockBookRequestValidation) CreateBookshelf(in *input.CreateBookshelf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBookshelf", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBookshelf indicates an expected call of CreateBookshelf
func (mr *MockBookRequestValidationMockRecorder) CreateBookshelf(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBookshelf", reflect.TypeOf((*MockBookRequestValidation)(nil).CreateBookshelf), in)
}

// UpdateBookshelf mocks base method
func (m *MockBookRequestValidation) UpdateBookshelf(in *input.UpdateBookshelf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBookshelf", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBookshelf indicates an expected call of UpdateBookshelf
func (mr *MockBookRequestValidationMockRecorder) UpdateBookshelf(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBookshelf", reflect.TypeOf((*MockBookRequestValidation)(nil).UpdateBookshelf), in)
}

// CreateAndUpdateBooks mocks base method
func (m *MockBookRequestValidation) CreateAndUpdateBooks(in *input.CreateAndUpdateBooks) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAndUpdateBooks", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAndUpdateBooks indicates an expected call of CreateAndUpdateBooks
func (mr *MockBookRequestValidationMockRecorder) CreateAndUpdateBooks(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndUpdateBooks", reflect.TypeOf((*MockBookRequestValidation)(nil).CreateAndUpdateBooks), in)
}
