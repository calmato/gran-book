// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/user/admin_service_grpc.pb.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	user "github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockAdminServiceClient is a mock of AdminServiceClient interface.
type MockAdminServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceClientMockRecorder
}

// MockAdminServiceClientMockRecorder is the mock recorder for MockAdminServiceClient.
type MockAdminServiceClientMockRecorder struct {
	mock *MockAdminServiceClient
}

// NewMockAdminServiceClient creates a new mock instance.
func NewMockAdminServiceClient(ctrl *gomock.Controller) *MockAdminServiceClient {
	mock := &MockAdminServiceClient{ctrl: ctrl}
	mock.recorder = &MockAdminServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminServiceClient) EXPECT() *MockAdminServiceClientMockRecorder {
	return m.recorder
}

// CreateAdmin mocks base method.
func (m *MockAdminServiceClient) CreateAdmin(ctx context.Context, in *user.CreateAdminRequest, opts ...grpc.CallOption) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAdmin", varargs...)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockAdminServiceClientMockRecorder) CreateAdmin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminServiceClient)(nil).CreateAdmin), varargs...)
}

// DeleteAdmin mocks base method.
func (m *MockAdminServiceClient) DeleteAdmin(ctx context.Context, in *user.DeleteAdminRequest, opts ...grpc.CallOption) (*user.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAdmin", varargs...)
	ret0, _ := ret[0].(*user.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockAdminServiceClientMockRecorder) DeleteAdmin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockAdminServiceClient)(nil).DeleteAdmin), varargs...)
}

// GetAdmin mocks base method.
func (m *MockAdminServiceClient) GetAdmin(ctx context.Context, in *user.GetAdminRequest, opts ...grpc.CallOption) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAdmin", varargs...)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockAdminServiceClientMockRecorder) GetAdmin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockAdminServiceClient)(nil).GetAdmin), varargs...)
}

// ListAdmin mocks base method.
func (m *MockAdminServiceClient) ListAdmin(ctx context.Context, in *user.ListAdminRequest, opts ...grpc.CallOption) (*user.AdminListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAdmin", varargs...)
	ret0, _ := ret[0].(*user.AdminListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAdmin indicates an expected call of ListAdmin.
func (mr *MockAdminServiceClientMockRecorder) ListAdmin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdmin", reflect.TypeOf((*MockAdminServiceClient)(nil).ListAdmin), varargs...)
}

// UpdateAdminContact mocks base method.
func (m *MockAdminServiceClient) UpdateAdminContact(ctx context.Context, in *user.UpdateAdminContactRequest, opts ...grpc.CallOption) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAdminContact", varargs...)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdminContact indicates an expected call of UpdateAdminContact.
func (mr *MockAdminServiceClientMockRecorder) UpdateAdminContact(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminContact", reflect.TypeOf((*MockAdminServiceClient)(nil).UpdateAdminContact), varargs...)
}

// UpdateAdminPassword mocks base method.
func (m *MockAdminServiceClient) UpdateAdminPassword(ctx context.Context, in *user.UpdateAdminPasswordRequest, opts ...grpc.CallOption) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAdminPassword", varargs...)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdminPassword indicates an expected call of UpdateAdminPassword.
func (mr *MockAdminServiceClientMockRecorder) UpdateAdminPassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminPassword", reflect.TypeOf((*MockAdminServiceClient)(nil).UpdateAdminPassword), varargs...)
}

// UpdateAdminProfile mocks base method.
func (m *MockAdminServiceClient) UpdateAdminProfile(ctx context.Context, in *user.UpdateAdminProfileRequest, opts ...grpc.CallOption) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAdminProfile", varargs...)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdminProfile indicates an expected call of UpdateAdminProfile.
func (mr *MockAdminServiceClientMockRecorder) UpdateAdminProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminProfile", reflect.TypeOf((*MockAdminServiceClient)(nil).UpdateAdminProfile), varargs...)
}

// UploadAdminThumbnail mocks base method.
func (m *MockAdminServiceClient) UploadAdminThumbnail(ctx context.Context, opts ...grpc.CallOption) (user.AdminService_UploadAdminThumbnailClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadAdminThumbnail", varargs...)
	ret0, _ := ret[0].(user.AdminService_UploadAdminThumbnailClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAdminThumbnail indicates an expected call of UploadAdminThumbnail.
func (mr *MockAdminServiceClientMockRecorder) UploadAdminThumbnail(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAdminThumbnail", reflect.TypeOf((*MockAdminServiceClient)(nil).UploadAdminThumbnail), varargs...)
}

// MockAdminService_UploadAdminThumbnailClient is a mock of AdminService_UploadAdminThumbnailClient interface.
type MockAdminService_UploadAdminThumbnailClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdminService_UploadAdminThumbnailClientMockRecorder
}

// MockAdminService_UploadAdminThumbnailClientMockRecorder is the mock recorder for MockAdminService_UploadAdminThumbnailClient.
type MockAdminService_UploadAdminThumbnailClientMockRecorder struct {
	mock *MockAdminService_UploadAdminThumbnailClient
}

// NewMockAdminService_UploadAdminThumbnailClient creates a new mock instance.
func NewMockAdminService_UploadAdminThumbnailClient(ctrl *gomock.Controller) *MockAdminService_UploadAdminThumbnailClient {
	mock := &MockAdminService_UploadAdminThumbnailClient{ctrl: ctrl}
	mock.recorder = &MockAdminService_UploadAdminThumbnailClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminService_UploadAdminThumbnailClient) EXPECT() *MockAdminService_UploadAdminThumbnailClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockAdminService_UploadAdminThumbnailClient) CloseAndRecv() (*user.AdminThumbnailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*user.AdminThumbnailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockAdminService_UploadAdminThumbnailClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockAdminService_UploadAdminThumbnailClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).Context))
}

// Header mocks base method.
func (m *MockAdminService_UploadAdminThumbnailClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m_2 *MockAdminService_UploadAdminThumbnailClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockAdminService_UploadAdminThumbnailClient) Send(arg0 *user.UploadAdminThumbnailRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockAdminService_UploadAdminThumbnailClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockAdminService_UploadAdminThumbnailClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockAdminService_UploadAdminThumbnailClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailClient)(nil).Trailer))
}

// MockAdminServiceServer is a mock of AdminServiceServer interface.
type MockAdminServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceServerMockRecorder
}

// MockAdminServiceServerMockRecorder is the mock recorder for MockAdminServiceServer.
type MockAdminServiceServerMockRecorder struct {
	mock *MockAdminServiceServer
}

// NewMockAdminServiceServer creates a new mock instance.
func NewMockAdminServiceServer(ctrl *gomock.Controller) *MockAdminServiceServer {
	mock := &MockAdminServiceServer{ctrl: ctrl}
	mock.recorder = &MockAdminServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminServiceServer) EXPECT() *MockAdminServiceServerMockRecorder {
	return m.recorder
}

// CreateAdmin mocks base method.
func (m *MockAdminServiceServer) CreateAdmin(arg0 context.Context, arg1 *user.CreateAdminRequest) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", arg0, arg1)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockAdminServiceServerMockRecorder) CreateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminServiceServer)(nil).CreateAdmin), arg0, arg1)
}

// DeleteAdmin mocks base method.
func (m *MockAdminServiceServer) DeleteAdmin(arg0 context.Context, arg1 *user.DeleteAdminRequest) (*user.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", arg0, arg1)
	ret0, _ := ret[0].(*user.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockAdminServiceServerMockRecorder) DeleteAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockAdminServiceServer)(nil).DeleteAdmin), arg0, arg1)
}

// GetAdmin mocks base method.
func (m *MockAdminServiceServer) GetAdmin(arg0 context.Context, arg1 *user.GetAdminRequest) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", arg0, arg1)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockAdminServiceServerMockRecorder) GetAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockAdminServiceServer)(nil).GetAdmin), arg0, arg1)
}

// ListAdmin mocks base method.
func (m *MockAdminServiceServer) ListAdmin(arg0 context.Context, arg1 *user.ListAdminRequest) (*user.AdminListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdmin", arg0, arg1)
	ret0, _ := ret[0].(*user.AdminListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAdmin indicates an expected call of ListAdmin.
func (mr *MockAdminServiceServerMockRecorder) ListAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdmin", reflect.TypeOf((*MockAdminServiceServer)(nil).ListAdmin), arg0, arg1)
}

// UpdateAdminContact mocks base method.
func (m *MockAdminServiceServer) UpdateAdminContact(arg0 context.Context, arg1 *user.UpdateAdminContactRequest) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminContact", arg0, arg1)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdminContact indicates an expected call of UpdateAdminContact.
func (mr *MockAdminServiceServerMockRecorder) UpdateAdminContact(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminContact", reflect.TypeOf((*MockAdminServiceServer)(nil).UpdateAdminContact), arg0, arg1)
}

// UpdateAdminPassword mocks base method.
func (m *MockAdminServiceServer) UpdateAdminPassword(arg0 context.Context, arg1 *user.UpdateAdminPasswordRequest) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminPassword", arg0, arg1)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdminPassword indicates an expected call of UpdateAdminPassword.
func (mr *MockAdminServiceServerMockRecorder) UpdateAdminPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminPassword", reflect.TypeOf((*MockAdminServiceServer)(nil).UpdateAdminPassword), arg0, arg1)
}

// UpdateAdminProfile mocks base method.
func (m *MockAdminServiceServer) UpdateAdminProfile(arg0 context.Context, arg1 *user.UpdateAdminProfileRequest) (*user.AdminResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminProfile", arg0, arg1)
	ret0, _ := ret[0].(*user.AdminResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdminProfile indicates an expected call of UpdateAdminProfile.
func (mr *MockAdminServiceServerMockRecorder) UpdateAdminProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminProfile", reflect.TypeOf((*MockAdminServiceServer)(nil).UpdateAdminProfile), arg0, arg1)
}

// UploadAdminThumbnail mocks base method.
func (m *MockAdminServiceServer) UploadAdminThumbnail(arg0 user.AdminService_UploadAdminThumbnailServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAdminThumbnail", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAdminThumbnail indicates an expected call of UploadAdminThumbnail.
func (mr *MockAdminServiceServerMockRecorder) UploadAdminThumbnail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAdminThumbnail", reflect.TypeOf((*MockAdminServiceServer)(nil).UploadAdminThumbnail), arg0)
}

// mustEmbedUnimplementedAdminServiceServer mocks base method.
func (m *MockAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAdminServiceServer")
}

// mustEmbedUnimplementedAdminServiceServer indicates an expected call of mustEmbedUnimplementedAdminServiceServer.
func (mr *MockAdminServiceServerMockRecorder) mustEmbedUnimplementedAdminServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAdminServiceServer", reflect.TypeOf((*MockAdminServiceServer)(nil).mustEmbedUnimplementedAdminServiceServer))
}

// MockUnsafeAdminServiceServer is a mock of UnsafeAdminServiceServer interface.
type MockUnsafeAdminServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAdminServiceServerMockRecorder
}

// MockUnsafeAdminServiceServerMockRecorder is the mock recorder for MockUnsafeAdminServiceServer.
type MockUnsafeAdminServiceServerMockRecorder struct {
	mock *MockUnsafeAdminServiceServer
}

// NewMockUnsafeAdminServiceServer creates a new mock instance.
func NewMockUnsafeAdminServiceServer(ctrl *gomock.Controller) *MockUnsafeAdminServiceServer {
	mock := &MockUnsafeAdminServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAdminServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAdminServiceServer) EXPECT() *MockUnsafeAdminServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAdminServiceServer mocks base method.
func (m *MockUnsafeAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAdminServiceServer")
}

// mustEmbedUnimplementedAdminServiceServer indicates an expected call of mustEmbedUnimplementedAdminServiceServer.
func (mr *MockUnsafeAdminServiceServerMockRecorder) mustEmbedUnimplementedAdminServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAdminServiceServer", reflect.TypeOf((*MockUnsafeAdminServiceServer)(nil).mustEmbedUnimplementedAdminServiceServer))
}

// MockAdminService_UploadAdminThumbnailServer is a mock of AdminService_UploadAdminThumbnailServer interface.
type MockAdminService_UploadAdminThumbnailServer struct {
	ctrl     *gomock.Controller
	recorder *MockAdminService_UploadAdminThumbnailServerMockRecorder
}

// MockAdminService_UploadAdminThumbnailServerMockRecorder is the mock recorder for MockAdminService_UploadAdminThumbnailServer.
type MockAdminService_UploadAdminThumbnailServerMockRecorder struct {
	mock *MockAdminService_UploadAdminThumbnailServer
}

// NewMockAdminService_UploadAdminThumbnailServer creates a new mock instance.
func NewMockAdminService_UploadAdminThumbnailServer(ctrl *gomock.Controller) *MockAdminService_UploadAdminThumbnailServer {
	mock := &MockAdminService_UploadAdminThumbnailServer{ctrl: ctrl}
	mock.recorder = &MockAdminService_UploadAdminThumbnailServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminService_UploadAdminThumbnailServer) EXPECT() *MockAdminService_UploadAdminThumbnailServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockAdminService_UploadAdminThumbnailServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockAdminService_UploadAdminThumbnailServer) Recv() (*user.UploadAdminThumbnailRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*user.UploadAdminThumbnailRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockAdminService_UploadAdminThumbnailServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).RecvMsg), m)
}

// SendAndClose mocks base method.
func (m *MockAdminService_UploadAdminThumbnailServer) SendAndClose(arg0 *user.AdminThumbnailResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).SendAndClose), arg0)
}

// SendHeader mocks base method.
func (m *MockAdminService_UploadAdminThumbnailServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockAdminService_UploadAdminThumbnailServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockAdminService_UploadAdminThumbnailServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockAdminService_UploadAdminThumbnailServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockAdminService_UploadAdminThumbnailServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockAdminService_UploadAdminThumbnailServer)(nil).SetTrailer), arg0)
}
