// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/book/validation.go

// Package mock_book is a generated GoMock package.
package mock_book

import (
	context "context"
	reflect "reflect"

	book "github.com/calmato/gran-book/api/server/book/internal/domain/book"
	gomock "github.com/golang/mock/gomock"
)

// MockValidation is a mock of Validation interface.
type MockValidation struct {
	ctrl     *gomock.Controller
	recorder *MockValidationMockRecorder
}

// MockValidationMockRecorder is the mock recorder for MockValidation.
type MockValidationMockRecorder struct {
	mock *MockValidation
}

// NewMockValidation creates a new mock instance.
func NewMockValidation(ctrl *gomock.Controller) *MockValidation {
	mock := &MockValidation{ctrl: ctrl}
	mock.recorder = &MockValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidation) EXPECT() *MockValidationMockRecorder {
	return m.recorder
}

// Author mocks base method.
func (m *MockValidation) Author(ctx context.Context, a *book.Author) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Author", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Author indicates an expected call of Author.
func (mr *MockValidationMockRecorder) Author(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Author", reflect.TypeOf((*MockValidation)(nil).Author), ctx, a)
}

// Book mocks base method.
func (m *MockValidation) Book(ctx context.Context, b *book.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Book", ctx, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Book indicates an expected call of Book.
func (mr *MockValidationMockRecorder) Book(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Book", reflect.TypeOf((*MockValidation)(nil).Book), ctx, b)
}

// Bookshelf mocks base method.
func (m *MockValidation) Bookshelf(ctx context.Context, bs *book.Bookshelf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bookshelf", ctx, bs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bookshelf indicates an expected call of Bookshelf.
func (mr *MockValidationMockRecorder) Bookshelf(ctx, bs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bookshelf", reflect.TypeOf((*MockValidation)(nil).Bookshelf), ctx, bs)
}

// Review mocks base method.
func (m *MockValidation) Review(ctx context.Context, rv *book.Review) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Review", ctx, rv)
	ret0, _ := ret[0].(error)
	return ret0
}

// Review indicates an expected call of Review.
func (mr *MockValidationMockRecorder) Review(ctx, rv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Review", reflect.TypeOf((*MockValidation)(nil).Review), ctx, rv)
}
