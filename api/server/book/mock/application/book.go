// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application//book.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	context "context"
	reflect "reflect"

	book "github.com/calmato/gran-book/api/server/book/internal/domain/book"
	database "github.com/calmato/gran-book/api/server/book/pkg/database"
	gomock "github.com/golang/mock/gomock"
)

// MockBookApplication is a mock of BookApplication interface.
type MockBookApplication struct {
	ctrl     *gomock.Controller
	recorder *MockBookApplicationMockRecorder
}

// MockBookApplicationMockRecorder is the mock recorder for MockBookApplication.
type MockBookApplicationMockRecorder struct {
	mock *MockBookApplication
}

// NewMockBookApplication creates a new mock instance.
func NewMockBookApplication(ctrl *gomock.Controller) *MockBookApplication {
	mock := &MockBookApplication{ctrl: ctrl}
	mock.recorder = &MockBookApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookApplication) EXPECT() *MockBookApplicationMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBookApplication) Create(ctx context.Context, b *book.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBookApplicationMockRecorder) Create(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookApplication)(nil).Create), ctx, b)
}

// CreateBookshelf mocks base method.
func (m *MockBookApplication) CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBookshelf", ctx, bs)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBookshelf indicates an expected call of CreateBookshelf.
func (mr *MockBookApplicationMockRecorder) CreateBookshelf(ctx, bs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBookshelf", reflect.TypeOf((*MockBookApplication)(nil).CreateBookshelf), ctx, bs)
}

// CreateOrUpdateBookshelf mocks base method.
func (m *MockBookApplication) CreateOrUpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateBookshelf", ctx, bs)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateBookshelf indicates an expected call of CreateOrUpdateBookshelf.
func (mr *MockBookApplicationMockRecorder) CreateOrUpdateBookshelf(ctx, bs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateBookshelf", reflect.TypeOf((*MockBookApplication)(nil).CreateOrUpdateBookshelf), ctx, bs)
}

// Delete mocks base method.
func (m *MockBookApplication) Delete(ctx context.Context, b *book.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBookApplicationMockRecorder) Delete(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookApplication)(nil).Delete), ctx, b)
}

// DeleteBookshelf mocks base method.
func (m *MockBookApplication) DeleteBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookshelf", ctx, bs)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBookshelf indicates an expected call of DeleteBookshelf.
func (mr *MockBookApplicationMockRecorder) DeleteBookshelf(ctx, bs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookshelf", reflect.TypeOf((*MockBookApplication)(nil).DeleteBookshelf), ctx, bs)
}

// Get mocks base method.
func (m *MockBookApplication) Get(ctx context.Context, bookID int) (*book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, bookID)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBookApplicationMockRecorder) Get(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBookApplication)(nil).Get), ctx, bookID)
}

// GetBookshelfByUserIDAndBookID mocks base method.
func (m *MockBookApplication) GetBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookshelfByUserIDAndBookID", ctx, userID, bookID)
	ret0, _ := ret[0].(*book.Bookshelf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookshelfByUserIDAndBookID indicates an expected call of GetBookshelfByUserIDAndBookID.
func (mr *MockBookApplicationMockRecorder) GetBookshelfByUserIDAndBookID(ctx, userID, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookshelfByUserIDAndBookID", reflect.TypeOf((*MockBookApplication)(nil).GetBookshelfByUserIDAndBookID), ctx, userID, bookID)
}

// GetBookshelfByUserIDAndBookIDWithRelated mocks base method.
func (m *MockBookApplication) GetBookshelfByUserIDAndBookIDWithRelated(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookshelfByUserIDAndBookIDWithRelated", ctx, userID, bookID)
	ret0, _ := ret[0].(*book.Bookshelf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookshelfByUserIDAndBookIDWithRelated indicates an expected call of GetBookshelfByUserIDAndBookIDWithRelated.
func (mr *MockBookApplicationMockRecorder) GetBookshelfByUserIDAndBookIDWithRelated(ctx, userID, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookshelfByUserIDAndBookIDWithRelated", reflect.TypeOf((*MockBookApplication)(nil).GetBookshelfByUserIDAndBookIDWithRelated), ctx, userID, bookID)
}

// GetByIsbn mocks base method.
func (m *MockBookApplication) GetByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIsbn", ctx, isbn)
	ret0, _ := ret[0].(*book.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIsbn indicates an expected call of GetByIsbn.
func (mr *MockBookApplicationMockRecorder) GetByIsbn(ctx, isbn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIsbn", reflect.TypeOf((*MockBookApplication)(nil).GetByIsbn), ctx, isbn)
}

// GetReview mocks base method.
func (m *MockBookApplication) GetReview(ctx context.Context, reviewID int) (*book.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", ctx, reviewID)
	ret0, _ := ret[0].(*book.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockBookApplicationMockRecorder) GetReview(ctx, reviewID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockBookApplication)(nil).GetReview), ctx, reviewID)
}

// GetReviewByUserIDAndBookID mocks base method.
func (m *MockBookApplication) GetReviewByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewByUserIDAndBookID", ctx, userID, bookID)
	ret0, _ := ret[0].(*book.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewByUserIDAndBookID indicates an expected call of GetReviewByUserIDAndBookID.
func (mr *MockBookApplicationMockRecorder) GetReviewByUserIDAndBookID(ctx, userID, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewByUserIDAndBookID", reflect.TypeOf((*MockBookApplication)(nil).GetReviewByUserIDAndBookID), ctx, userID, bookID)
}

// List mocks base method.
func (m *MockBookApplication) List(ctx context.Context, q *database.ListQuery) (book.Books, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, q)
	ret0, _ := ret[0].(book.Books)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockBookApplicationMockRecorder) List(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBookApplication)(nil).List), ctx, q)
}

// ListBookReview mocks base method.
func (m *MockBookApplication) ListBookReview(ctx context.Context, bookID, limit, offset int) (book.Reviews, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBookReview", ctx, bookID, limit, offset)
	ret0, _ := ret[0].(book.Reviews)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBookReview indicates an expected call of ListBookReview.
func (mr *MockBookApplicationMockRecorder) ListBookReview(ctx, bookID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBookReview", reflect.TypeOf((*MockBookApplication)(nil).ListBookReview), ctx, bookID, limit, offset)
}

// ListBookshelf mocks base method.
func (m *MockBookApplication) ListBookshelf(ctx context.Context, q *database.ListQuery) (book.Bookshelves, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBookshelf", ctx, q)
	ret0, _ := ret[0].(book.Bookshelves)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBookshelf indicates an expected call of ListBookshelf.
func (mr *MockBookApplicationMockRecorder) ListBookshelf(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBookshelf", reflect.TypeOf((*MockBookApplication)(nil).ListBookshelf), ctx, q)
}

// ListUserReview mocks base method.
func (m *MockBookApplication) ListUserReview(ctx context.Context, userID string, limit, offset int) (book.Reviews, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserReview", ctx, userID, limit, offset)
	ret0, _ := ret[0].(book.Reviews)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUserReview indicates an expected call of ListUserReview.
func (mr *MockBookApplicationMockRecorder) ListUserReview(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserReview", reflect.TypeOf((*MockBookApplication)(nil).ListUserReview), ctx, userID, limit, offset)
}

// MultiGet mocks base method.
func (m *MockBookApplication) MultiGet(ctx context.Context, bookIDs []int) (book.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGet", ctx, bookIDs)
	ret0, _ := ret[0].(book.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockBookApplicationMockRecorder) MultiGet(ctx, bookIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockBookApplication)(nil).MultiGet), ctx, bookIDs)
}

// MultipleCreate mocks base method.
func (m *MockBookApplication) MultipleCreate(ctx context.Context, bs book.Books) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipleCreate", ctx, bs)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultipleCreate indicates an expected call of MultipleCreate.
func (mr *MockBookApplicationMockRecorder) MultipleCreate(ctx, bs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipleCreate", reflect.TypeOf((*MockBookApplication)(nil).MultipleCreate), ctx, bs)
}

// MultipleUpdate mocks base method.
func (m *MockBookApplication) MultipleUpdate(ctx context.Context, bs book.Books) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipleUpdate", ctx, bs)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultipleUpdate indicates an expected call of MultipleUpdate.
func (mr *MockBookApplicationMockRecorder) MultipleUpdate(ctx, bs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipleUpdate", reflect.TypeOf((*MockBookApplication)(nil).MultipleUpdate), ctx, bs)
}

// Update mocks base method.
func (m *MockBookApplication) Update(ctx context.Context, b *book.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBookApplicationMockRecorder) Update(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookApplication)(nil).Update), ctx, b)
}

// UpdateBookshelf mocks base method.
func (m *MockBookApplication) UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBookshelf", ctx, bs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBookshelf indicates an expected call of UpdateBookshelf.
func (mr *MockBookApplicationMockRecorder) UpdateBookshelf(ctx, bs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBookshelf", reflect.TypeOf((*MockBookApplication)(nil).UpdateBookshelf), ctx, bs)
}
