// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/book/book_service_grpc.pb.go

// Package mock_book is a generated GoMock package.
package mock_book

import (
	context "context"
	reflect "reflect"

	book "github.com/calmato/gran-book/api/proto/book"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBookServiceClient is a mock of BookServiceClient interface.
type MockBookServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceClientMockRecorder
}

// MockBookServiceClientMockRecorder is the mock recorder for MockBookServiceClient.
type MockBookServiceClientMockRecorder struct {
	mock *MockBookServiceClient
}

// NewMockBookServiceClient creates a new mock instance.
func NewMockBookServiceClient(ctrl *gomock.Controller) *MockBookServiceClient {
	mock := &MockBookServiceClient{ctrl: ctrl}
	mock.recorder = &MockBookServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServiceClient) EXPECT() *MockBookServiceClientMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookServiceClient) CreateBook(ctx context.Context, in *book.CreateBookRequest, opts ...grpc.CallOption) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBook", varargs...)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookServiceClientMockRecorder) CreateBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookServiceClient)(nil).CreateBook), varargs...)
}

// DeleteBook mocks base method.
func (m *MockBookServiceClient) DeleteBook(ctx context.Context, in *book.DeleteBookRequest, opts ...grpc.CallOption) (*book.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBook", varargs...)
	ret0, _ := ret[0].(*book.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookServiceClientMockRecorder) DeleteBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookServiceClient)(nil).DeleteBook), varargs...)
}

// DeleteBookshelf mocks base method.
func (m *MockBookServiceClient) DeleteBookshelf(ctx context.Context, in *book.DeleteBookshelfRequest, opts ...grpc.CallOption) (*book.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBookshelf", varargs...)
	ret0, _ := ret[0].(*book.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBookshelf indicates an expected call of DeleteBookshelf.
func (mr *MockBookServiceClientMockRecorder) DeleteBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).DeleteBookshelf), varargs...)
}

// GetBook mocks base method.
func (m *MockBookServiceClient) GetBook(ctx context.Context, in *book.GetBookRequest, opts ...grpc.CallOption) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBook", varargs...)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookServiceClientMockRecorder) GetBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookServiceClient)(nil).GetBook), varargs...)
}

// GetBookByIsbn mocks base method.
func (m *MockBookServiceClient) GetBookByIsbn(ctx context.Context, in *book.GetBookByIsbnRequest, opts ...grpc.CallOption) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBookByIsbn", varargs...)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByIsbn indicates an expected call of GetBookByIsbn.
func (mr *MockBookServiceClientMockRecorder) GetBookByIsbn(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByIsbn", reflect.TypeOf((*MockBookServiceClient)(nil).GetBookByIsbn), varargs...)
}

// GetBookshelf mocks base method.
func (m *MockBookServiceClient) GetBookshelf(ctx context.Context, in *book.GetBookshelfRequest, opts ...grpc.CallOption) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBookshelf", varargs...)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookshelf indicates an expected call of GetBookshelf.
func (mr *MockBookServiceClientMockRecorder) GetBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).GetBookshelf), varargs...)
}

// GetReview mocks base method.
func (m *MockBookServiceClient) GetReview(ctx context.Context, in *book.GetReviewRequest, opts ...grpc.CallOption) (*book.ReviewResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReview", varargs...)
	ret0, _ := ret[0].(*book.ReviewResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockBookServiceClientMockRecorder) GetReview(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockBookServiceClient)(nil).GetReview), varargs...)
}

// ListBookReview mocks base method.
func (m *MockBookServiceClient) ListBookReview(ctx context.Context, in *book.ListBookReviewRequest, opts ...grpc.CallOption) (*book.ReviewListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBookReview", varargs...)
	ret0, _ := ret[0].(*book.ReviewListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBookReview indicates an expected call of ListBookReview.
func (mr *MockBookServiceClientMockRecorder) ListBookReview(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBookReview", reflect.TypeOf((*MockBookServiceClient)(nil).ListBookReview), varargs...)
}

// ListBookshelf mocks base method.
func (m *MockBookServiceClient) ListBookshelf(ctx context.Context, in *book.ListBookshelfRequest, opts ...grpc.CallOption) (*book.BookshelfListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBookshelf", varargs...)
	ret0, _ := ret[0].(*book.BookshelfListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBookshelf indicates an expected call of ListBookshelf.
func (mr *MockBookServiceClientMockRecorder) ListBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).ListBookshelf), varargs...)
}

// ListUserMonthlyResult mocks base method.
func (m *MockBookServiceClient) ListUserMonthlyResult(ctx context.Context, in *book.ListUserMonthlyResultRequest, opts ...grpc.CallOption) (*book.UserMonthlyResultListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserMonthlyResult", varargs...)
	ret0, _ := ret[0].(*book.UserMonthlyResultListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserMonthlyResult indicates an expected call of ListUserMonthlyResult.
func (mr *MockBookServiceClientMockRecorder) ListUserMonthlyResult(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserMonthlyResult", reflect.TypeOf((*MockBookServiceClient)(nil).ListUserMonthlyResult), varargs...)
}

// ListUserReview mocks base method.
func (m *MockBookServiceClient) ListUserReview(ctx context.Context, in *book.ListUserReviewRequest, opts ...grpc.CallOption) (*book.ReviewListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserReview", varargs...)
	ret0, _ := ret[0].(*book.ReviewListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserReview indicates an expected call of ListUserReview.
func (mr *MockBookServiceClientMockRecorder) ListUserReview(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserReview", reflect.TypeOf((*MockBookServiceClient)(nil).ListUserReview), varargs...)
}

// MultiGetBooks mocks base method.
func (m *MockBookServiceClient) MultiGetBooks(ctx context.Context, in *book.MultiGetBooksRequest, opts ...grpc.CallOption) (*book.BookListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGetBooks", varargs...)
	ret0, _ := ret[0].(*book.BookListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetBooks indicates an expected call of MultiGetBooks.
func (mr *MockBookServiceClientMockRecorder) MultiGetBooks(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetBooks", reflect.TypeOf((*MockBookServiceClient)(nil).MultiGetBooks), varargs...)
}

// ReadBookshelf mocks base method.
func (m *MockBookServiceClient) ReadBookshelf(ctx context.Context, in *book.ReadBookshelfRequest, opts ...grpc.CallOption) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadBookshelf", varargs...)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBookshelf indicates an expected call of ReadBookshelf.
func (mr *MockBookServiceClientMockRecorder) ReadBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).ReadBookshelf), varargs...)
}

// ReadingBookshelf mocks base method.
func (m *MockBookServiceClient) ReadingBookshelf(ctx context.Context, in *book.ReadingBookshelfRequest, opts ...grpc.CallOption) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadingBookshelf", varargs...)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadingBookshelf indicates an expected call of ReadingBookshelf.
func (mr *MockBookServiceClientMockRecorder) ReadingBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadingBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).ReadingBookshelf), varargs...)
}

// ReleaseBookshelf mocks base method.
func (m *MockBookServiceClient) ReleaseBookshelf(ctx context.Context, in *book.ReleaseBookshelfRequest, opts ...grpc.CallOption) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReleaseBookshelf", varargs...)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseBookshelf indicates an expected call of ReleaseBookshelf.
func (mr *MockBookServiceClientMockRecorder) ReleaseBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).ReleaseBookshelf), varargs...)
}

// StackedBookshelf mocks base method.
func (m *MockBookServiceClient) StackedBookshelf(ctx context.Context, in *book.StackedBookshelfRequest, opts ...grpc.CallOption) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StackedBookshelf", varargs...)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StackedBookshelf indicates an expected call of StackedBookshelf.
func (mr *MockBookServiceClientMockRecorder) StackedBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackedBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).StackedBookshelf), varargs...)
}

// UpdateBook mocks base method.
func (m *MockBookServiceClient) UpdateBook(ctx context.Context, in *book.UpdateBookRequest, opts ...grpc.CallOption) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBook", varargs...)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookServiceClientMockRecorder) UpdateBook(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookServiceClient)(nil).UpdateBook), varargs...)
}

// WantBookshelf mocks base method.
func (m *MockBookServiceClient) WantBookshelf(ctx context.Context, in *book.WantBookshelfRequest, opts ...grpc.CallOption) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WantBookshelf", varargs...)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WantBookshelf indicates an expected call of WantBookshelf.
func (mr *MockBookServiceClientMockRecorder) WantBookshelf(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WantBookshelf", reflect.TypeOf((*MockBookServiceClient)(nil).WantBookshelf), varargs...)
}

// MockBookServiceServer is a mock of BookServiceServer interface.
type MockBookServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceServerMockRecorder
}

// MockBookServiceServerMockRecorder is the mock recorder for MockBookServiceServer.
type MockBookServiceServerMockRecorder struct {
	mock *MockBookServiceServer
}

// NewMockBookServiceServer creates a new mock instance.
func NewMockBookServiceServer(ctrl *gomock.Controller) *MockBookServiceServer {
	mock := &MockBookServiceServer{ctrl: ctrl}
	mock.recorder = &MockBookServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServiceServer) EXPECT() *MockBookServiceServerMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockBookServiceServer) CreateBook(arg0 context.Context, arg1 *book.CreateBookRequest) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0, arg1)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockBookServiceServerMockRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookServiceServer)(nil).CreateBook), arg0, arg1)
}

// DeleteBook mocks base method.
func (m *MockBookServiceServer) DeleteBook(arg0 context.Context, arg1 *book.DeleteBookRequest) (*book.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0, arg1)
	ret0, _ := ret[0].(*book.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookServiceServerMockRecorder) DeleteBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBookServiceServer)(nil).DeleteBook), arg0, arg1)
}

// DeleteBookshelf mocks base method.
func (m *MockBookServiceServer) DeleteBookshelf(arg0 context.Context, arg1 *book.DeleteBookshelfRequest) (*book.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBookshelf indicates an expected call of DeleteBookshelf.
func (mr *MockBookServiceServerMockRecorder) DeleteBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).DeleteBookshelf), arg0, arg1)
}

// GetBook mocks base method.
func (m *MockBookServiceServer) GetBook(arg0 context.Context, arg1 *book.GetBookRequest) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", arg0, arg1)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookServiceServerMockRecorder) GetBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookServiceServer)(nil).GetBook), arg0, arg1)
}

// GetBookByIsbn mocks base method.
func (m *MockBookServiceServer) GetBookByIsbn(arg0 context.Context, arg1 *book.GetBookByIsbnRequest) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByIsbn", arg0, arg1)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByIsbn indicates an expected call of GetBookByIsbn.
func (mr *MockBookServiceServerMockRecorder) GetBookByIsbn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByIsbn", reflect.TypeOf((*MockBookServiceServer)(nil).GetBookByIsbn), arg0, arg1)
}

// GetBookshelf mocks base method.
func (m *MockBookServiceServer) GetBookshelf(arg0 context.Context, arg1 *book.GetBookshelfRequest) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookshelf indicates an expected call of GetBookshelf.
func (mr *MockBookServiceServerMockRecorder) GetBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).GetBookshelf), arg0, arg1)
}

// GetReview mocks base method.
func (m *MockBookServiceServer) GetReview(arg0 context.Context, arg1 *book.GetReviewRequest) (*book.ReviewResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", arg0, arg1)
	ret0, _ := ret[0].(*book.ReviewResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockBookServiceServerMockRecorder) GetReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockBookServiceServer)(nil).GetReview), arg0, arg1)
}

// ListBookReview mocks base method.
func (m *MockBookServiceServer) ListBookReview(arg0 context.Context, arg1 *book.ListBookReviewRequest) (*book.ReviewListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBookReview", arg0, arg1)
	ret0, _ := ret[0].(*book.ReviewListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBookReview indicates an expected call of ListBookReview.
func (mr *MockBookServiceServerMockRecorder) ListBookReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBookReview", reflect.TypeOf((*MockBookServiceServer)(nil).ListBookReview), arg0, arg1)
}

// ListBookshelf mocks base method.
func (m *MockBookServiceServer) ListBookshelf(arg0 context.Context, arg1 *book.ListBookshelfRequest) (*book.BookshelfListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.BookshelfListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBookshelf indicates an expected call of ListBookshelf.
func (mr *MockBookServiceServerMockRecorder) ListBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).ListBookshelf), arg0, arg1)
}

// ListUserMonthlyResult mocks base method.
func (m *MockBookServiceServer) ListUserMonthlyResult(arg0 context.Context, arg1 *book.ListUserMonthlyResultRequest) (*book.UserMonthlyResultListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserMonthlyResult", arg0, arg1)
	ret0, _ := ret[0].(*book.UserMonthlyResultListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserMonthlyResult indicates an expected call of ListUserMonthlyResult.
func (mr *MockBookServiceServerMockRecorder) ListUserMonthlyResult(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserMonthlyResult", reflect.TypeOf((*MockBookServiceServer)(nil).ListUserMonthlyResult), arg0, arg1)
}

// ListUserReview mocks base method.
func (m *MockBookServiceServer) ListUserReview(arg0 context.Context, arg1 *book.ListUserReviewRequest) (*book.ReviewListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserReview", arg0, arg1)
	ret0, _ := ret[0].(*book.ReviewListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserReview indicates an expected call of ListUserReview.
func (mr *MockBookServiceServerMockRecorder) ListUserReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserReview", reflect.TypeOf((*MockBookServiceServer)(nil).ListUserReview), arg0, arg1)
}

// MultiGetBooks mocks base method.
func (m *MockBookServiceServer) MultiGetBooks(arg0 context.Context, arg1 *book.MultiGetBooksRequest) (*book.BookListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetBooks", arg0, arg1)
	ret0, _ := ret[0].(*book.BookListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetBooks indicates an expected call of MultiGetBooks.
func (mr *MockBookServiceServerMockRecorder) MultiGetBooks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetBooks", reflect.TypeOf((*MockBookServiceServer)(nil).MultiGetBooks), arg0, arg1)
}

// ReadBookshelf mocks base method.
func (m *MockBookServiceServer) ReadBookshelf(arg0 context.Context, arg1 *book.ReadBookshelfRequest) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBookshelf indicates an expected call of ReadBookshelf.
func (mr *MockBookServiceServerMockRecorder) ReadBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).ReadBookshelf), arg0, arg1)
}

// ReadingBookshelf mocks base method.
func (m *MockBookServiceServer) ReadingBookshelf(arg0 context.Context, arg1 *book.ReadingBookshelfRequest) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadingBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadingBookshelf indicates an expected call of ReadingBookshelf.
func (mr *MockBookServiceServerMockRecorder) ReadingBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadingBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).ReadingBookshelf), arg0, arg1)
}

// ReleaseBookshelf mocks base method.
func (m *MockBookServiceServer) ReleaseBookshelf(arg0 context.Context, arg1 *book.ReleaseBookshelfRequest) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseBookshelf indicates an expected call of ReleaseBookshelf.
func (mr *MockBookServiceServerMockRecorder) ReleaseBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).ReleaseBookshelf), arg0, arg1)
}

// StackedBookshelf mocks base method.
func (m *MockBookServiceServer) StackedBookshelf(arg0 context.Context, arg1 *book.StackedBookshelfRequest) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackedBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StackedBookshelf indicates an expected call of StackedBookshelf.
func (mr *MockBookServiceServerMockRecorder) StackedBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackedBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).StackedBookshelf), arg0, arg1)
}

// UpdateBook mocks base method.
func (m *MockBookServiceServer) UpdateBook(arg0 context.Context, arg1 *book.UpdateBookRequest) (*book.BookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", arg0, arg1)
	ret0, _ := ret[0].(*book.BookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockBookServiceServerMockRecorder) UpdateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookServiceServer)(nil).UpdateBook), arg0, arg1)
}

// WantBookshelf mocks base method.
func (m *MockBookServiceServer) WantBookshelf(arg0 context.Context, arg1 *book.WantBookshelfRequest) (*book.BookshelfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WantBookshelf", arg0, arg1)
	ret0, _ := ret[0].(*book.BookshelfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WantBookshelf indicates an expected call of WantBookshelf.
func (mr *MockBookServiceServerMockRecorder) WantBookshelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WantBookshelf", reflect.TypeOf((*MockBookServiceServer)(nil).WantBookshelf), arg0, arg1)
}

// mustEmbedUnimplementedBookServiceServer mocks base method.
func (m *MockBookServiceServer) mustEmbedUnimplementedBookServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBookServiceServer")
}

// mustEmbedUnimplementedBookServiceServer indicates an expected call of mustEmbedUnimplementedBookServiceServer.
func (mr *MockBookServiceServerMockRecorder) mustEmbedUnimplementedBookServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBookServiceServer", reflect.TypeOf((*MockBookServiceServer)(nil).mustEmbedUnimplementedBookServiceServer))
}

// MockUnsafeBookServiceServer is a mock of UnsafeBookServiceServer interface.
type MockUnsafeBookServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBookServiceServerMockRecorder
}

// MockUnsafeBookServiceServerMockRecorder is the mock recorder for MockUnsafeBookServiceServer.
type MockUnsafeBookServiceServerMockRecorder struct {
	mock *MockUnsafeBookServiceServer
}

// NewMockUnsafeBookServiceServer creates a new mock instance.
func NewMockUnsafeBookServiceServer(ctrl *gomock.Controller) *MockUnsafeBookServiceServer {
	mock := &MockUnsafeBookServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBookServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBookServiceServer) EXPECT() *MockUnsafeBookServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBookServiceServer mocks base method.
func (m *MockUnsafeBookServiceServer) mustEmbedUnimplementedBookServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBookServiceServer")
}

// mustEmbedUnimplementedBookServiceServer indicates an expected call of mustEmbedUnimplementedBookServiceServer.
func (mr *MockUnsafeBookServiceServerMockRecorder) mustEmbedUnimplementedBookServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBookServiceServer", reflect.TypeOf((*MockUnsafeBookServiceServer)(nil).mustEmbedUnimplementedBookServiceServer))
}
