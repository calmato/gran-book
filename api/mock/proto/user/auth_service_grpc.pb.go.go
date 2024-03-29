// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/user/auth_service_grpc.pb.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	user "github.com/calmato/gran-book/api/proto/user"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockAuthServiceClient is a mock of AuthServiceClient interface.
type MockAuthServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceClientMockRecorder
}

// MockAuthServiceClientMockRecorder is the mock recorder for MockAuthServiceClient.
type MockAuthServiceClientMockRecorder struct {
	mock *MockAuthServiceClient
}

// NewMockAuthServiceClient creates a new mock instance.
func NewMockAuthServiceClient(ctrl *gomock.Controller) *MockAuthServiceClient {
	mock := &MockAuthServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceClient) EXPECT() *MockAuthServiceClientMockRecorder {
	return m.recorder
}

// CreateAuth mocks base method.
func (m *MockAuthServiceClient) CreateAuth(ctx context.Context, in *user.CreateAuthRequest, opts ...grpc.CallOption) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAuth", varargs...)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuth indicates an expected call of CreateAuth.
func (mr *MockAuthServiceClientMockRecorder) CreateAuth(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuth", reflect.TypeOf((*MockAuthServiceClient)(nil).CreateAuth), varargs...)
}

// DeleteAuth mocks base method.
func (m *MockAuthServiceClient) DeleteAuth(ctx context.Context, in *user.Empty, opts ...grpc.CallOption) (*user.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAuth", varargs...)
	ret0, _ := ret[0].(*user.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAuth indicates an expected call of DeleteAuth.
func (mr *MockAuthServiceClientMockRecorder) DeleteAuth(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuth", reflect.TypeOf((*MockAuthServiceClient)(nil).DeleteAuth), varargs...)
}

// GetAuth mocks base method.
func (m *MockAuthServiceClient) GetAuth(ctx context.Context, in *user.Empty, opts ...grpc.CallOption) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuth", varargs...)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuth indicates an expected call of GetAuth.
func (mr *MockAuthServiceClientMockRecorder) GetAuth(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuth", reflect.TypeOf((*MockAuthServiceClient)(nil).GetAuth), varargs...)
}

// RegisterAuthDevice mocks base method.
func (m *MockAuthServiceClient) RegisterAuthDevice(ctx context.Context, in *user.RegisterAuthDeviceRequest, opts ...grpc.CallOption) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterAuthDevice", varargs...)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAuthDevice indicates an expected call of RegisterAuthDevice.
func (mr *MockAuthServiceClientMockRecorder) RegisterAuthDevice(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAuthDevice", reflect.TypeOf((*MockAuthServiceClient)(nil).RegisterAuthDevice), varargs...)
}

// UpdateAuthAddress mocks base method.
func (m *MockAuthServiceClient) UpdateAuthAddress(ctx context.Context, in *user.UpdateAuthAddressRequest, opts ...grpc.CallOption) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthAddress", varargs...)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthAddress indicates an expected call of UpdateAuthAddress.
func (mr *MockAuthServiceClientMockRecorder) UpdateAuthAddress(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthAddress", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateAuthAddress), varargs...)
}

// UpdateAuthEmail mocks base method.
func (m *MockAuthServiceClient) UpdateAuthEmail(ctx context.Context, in *user.UpdateAuthEmailRequest, opts ...grpc.CallOption) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthEmail", varargs...)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthEmail indicates an expected call of UpdateAuthEmail.
func (mr *MockAuthServiceClientMockRecorder) UpdateAuthEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthEmail", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateAuthEmail), varargs...)
}

// UpdateAuthPassword mocks base method.
func (m *MockAuthServiceClient) UpdateAuthPassword(ctx context.Context, in *user.UpdateAuthPasswordRequest, opts ...grpc.CallOption) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthPassword", varargs...)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthPassword indicates an expected call of UpdateAuthPassword.
func (mr *MockAuthServiceClientMockRecorder) UpdateAuthPassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthPassword", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateAuthPassword), varargs...)
}

// UpdateAuthProfile mocks base method.
func (m *MockAuthServiceClient) UpdateAuthProfile(ctx context.Context, in *user.UpdateAuthProfileRequest, opts ...grpc.CallOption) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAuthProfile", varargs...)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthProfile indicates an expected call of UpdateAuthProfile.
func (mr *MockAuthServiceClientMockRecorder) UpdateAuthProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthProfile", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateAuthProfile), varargs...)
}

// UploadAuthThumbnail mocks base method.
func (m *MockAuthServiceClient) UploadAuthThumbnail(ctx context.Context, opts ...grpc.CallOption) (user.AuthService_UploadAuthThumbnailClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadAuthThumbnail", varargs...)
	ret0, _ := ret[0].(user.AuthService_UploadAuthThumbnailClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAuthThumbnail indicates an expected call of UploadAuthThumbnail.
func (mr *MockAuthServiceClientMockRecorder) UploadAuthThumbnail(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAuthThumbnail", reflect.TypeOf((*MockAuthServiceClient)(nil).UploadAuthThumbnail), varargs...)
}

// MockAuthService_UploadAuthThumbnailClient is a mock of AuthService_UploadAuthThumbnailClient interface.
type MockAuthService_UploadAuthThumbnailClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthService_UploadAuthThumbnailClientMockRecorder
}

// MockAuthService_UploadAuthThumbnailClientMockRecorder is the mock recorder for MockAuthService_UploadAuthThumbnailClient.
type MockAuthService_UploadAuthThumbnailClientMockRecorder struct {
	mock *MockAuthService_UploadAuthThumbnailClient
}

// NewMockAuthService_UploadAuthThumbnailClient creates a new mock instance.
func NewMockAuthService_UploadAuthThumbnailClient(ctrl *gomock.Controller) *MockAuthService_UploadAuthThumbnailClient {
	mock := &MockAuthService_UploadAuthThumbnailClient{ctrl: ctrl}
	mock.recorder = &MockAuthService_UploadAuthThumbnailClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService_UploadAuthThumbnailClient) EXPECT() *MockAuthService_UploadAuthThumbnailClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockAuthService_UploadAuthThumbnailClient) CloseAndRecv() (*user.AuthThumbnailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*user.AuthThumbnailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockAuthService_UploadAuthThumbnailClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockAuthService_UploadAuthThumbnailClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).Context))
}

// Header mocks base method.
func (m *MockAuthService_UploadAuthThumbnailClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m_2 *MockAuthService_UploadAuthThumbnailClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockAuthService_UploadAuthThumbnailClient) Send(arg0 *user.UploadAuthThumbnailRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockAuthService_UploadAuthThumbnailClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockAuthService_UploadAuthThumbnailClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockAuthService_UploadAuthThumbnailClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailClient)(nil).Trailer))
}

// MockAuthServiceServer is a mock of AuthServiceServer interface.
type MockAuthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceServerMockRecorder
}

// MockAuthServiceServerMockRecorder is the mock recorder for MockAuthServiceServer.
type MockAuthServiceServerMockRecorder struct {
	mock *MockAuthServiceServer
}

// NewMockAuthServiceServer creates a new mock instance.
func NewMockAuthServiceServer(ctrl *gomock.Controller) *MockAuthServiceServer {
	mock := &MockAuthServiceServer{ctrl: ctrl}
	mock.recorder = &MockAuthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceServer) EXPECT() *MockAuthServiceServerMockRecorder {
	return m.recorder
}

// CreateAuth mocks base method.
func (m *MockAuthServiceServer) CreateAuth(arg0 context.Context, arg1 *user.CreateAuthRequest) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuth", arg0, arg1)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuth indicates an expected call of CreateAuth.
func (mr *MockAuthServiceServerMockRecorder) CreateAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuth", reflect.TypeOf((*MockAuthServiceServer)(nil).CreateAuth), arg0, arg1)
}

// DeleteAuth mocks base method.
func (m *MockAuthServiceServer) DeleteAuth(arg0 context.Context, arg1 *user.Empty) (*user.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuth", arg0, arg1)
	ret0, _ := ret[0].(*user.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAuth indicates an expected call of DeleteAuth.
func (mr *MockAuthServiceServerMockRecorder) DeleteAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuth", reflect.TypeOf((*MockAuthServiceServer)(nil).DeleteAuth), arg0, arg1)
}

// GetAuth mocks base method.
func (m *MockAuthServiceServer) GetAuth(arg0 context.Context, arg1 *user.Empty) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuth", arg0, arg1)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuth indicates an expected call of GetAuth.
func (mr *MockAuthServiceServerMockRecorder) GetAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuth", reflect.TypeOf((*MockAuthServiceServer)(nil).GetAuth), arg0, arg1)
}

// RegisterAuthDevice mocks base method.
func (m *MockAuthServiceServer) RegisterAuthDevice(arg0 context.Context, arg1 *user.RegisterAuthDeviceRequest) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAuthDevice", arg0, arg1)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAuthDevice indicates an expected call of RegisterAuthDevice.
func (mr *MockAuthServiceServerMockRecorder) RegisterAuthDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAuthDevice", reflect.TypeOf((*MockAuthServiceServer)(nil).RegisterAuthDevice), arg0, arg1)
}

// UpdateAuthAddress mocks base method.
func (m *MockAuthServiceServer) UpdateAuthAddress(arg0 context.Context, arg1 *user.UpdateAuthAddressRequest) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthAddress", arg0, arg1)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthAddress indicates an expected call of UpdateAuthAddress.
func (mr *MockAuthServiceServerMockRecorder) UpdateAuthAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthAddress", reflect.TypeOf((*MockAuthServiceServer)(nil).UpdateAuthAddress), arg0, arg1)
}

// UpdateAuthEmail mocks base method.
func (m *MockAuthServiceServer) UpdateAuthEmail(arg0 context.Context, arg1 *user.UpdateAuthEmailRequest) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthEmail", arg0, arg1)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthEmail indicates an expected call of UpdateAuthEmail.
func (mr *MockAuthServiceServerMockRecorder) UpdateAuthEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthEmail", reflect.TypeOf((*MockAuthServiceServer)(nil).UpdateAuthEmail), arg0, arg1)
}

// UpdateAuthPassword mocks base method.
func (m *MockAuthServiceServer) UpdateAuthPassword(arg0 context.Context, arg1 *user.UpdateAuthPasswordRequest) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthPassword", arg0, arg1)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthPassword indicates an expected call of UpdateAuthPassword.
func (mr *MockAuthServiceServerMockRecorder) UpdateAuthPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthPassword", reflect.TypeOf((*MockAuthServiceServer)(nil).UpdateAuthPassword), arg0, arg1)
}

// UpdateAuthProfile mocks base method.
func (m *MockAuthServiceServer) UpdateAuthProfile(arg0 context.Context, arg1 *user.UpdateAuthProfileRequest) (*user.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthProfile", arg0, arg1)
	ret0, _ := ret[0].(*user.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthProfile indicates an expected call of UpdateAuthProfile.
func (mr *MockAuthServiceServerMockRecorder) UpdateAuthProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthProfile", reflect.TypeOf((*MockAuthServiceServer)(nil).UpdateAuthProfile), arg0, arg1)
}

// UploadAuthThumbnail mocks base method.
func (m *MockAuthServiceServer) UploadAuthThumbnail(arg0 user.AuthService_UploadAuthThumbnailServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAuthThumbnail", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAuthThumbnail indicates an expected call of UploadAuthThumbnail.
func (mr *MockAuthServiceServerMockRecorder) UploadAuthThumbnail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAuthThumbnail", reflect.TypeOf((*MockAuthServiceServer)(nil).UploadAuthThumbnail), arg0)
}

// mustEmbedUnimplementedAuthServiceServer mocks base method.
func (m *MockAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServiceServer")
}

// mustEmbedUnimplementedAuthServiceServer indicates an expected call of mustEmbedUnimplementedAuthServiceServer.
func (mr *MockAuthServiceServerMockRecorder) mustEmbedUnimplementedAuthServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServiceServer", reflect.TypeOf((*MockAuthServiceServer)(nil).mustEmbedUnimplementedAuthServiceServer))
}

// MockUnsafeAuthServiceServer is a mock of UnsafeAuthServiceServer interface.
type MockUnsafeAuthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAuthServiceServerMockRecorder
}

// MockUnsafeAuthServiceServerMockRecorder is the mock recorder for MockUnsafeAuthServiceServer.
type MockUnsafeAuthServiceServerMockRecorder struct {
	mock *MockUnsafeAuthServiceServer
}

// NewMockUnsafeAuthServiceServer creates a new mock instance.
func NewMockUnsafeAuthServiceServer(ctrl *gomock.Controller) *MockUnsafeAuthServiceServer {
	mock := &MockUnsafeAuthServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAuthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAuthServiceServer) EXPECT() *MockUnsafeAuthServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAuthServiceServer mocks base method.
func (m *MockUnsafeAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAuthServiceServer")
}

// mustEmbedUnimplementedAuthServiceServer indicates an expected call of mustEmbedUnimplementedAuthServiceServer.
func (mr *MockUnsafeAuthServiceServerMockRecorder) mustEmbedUnimplementedAuthServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthServiceServer", reflect.TypeOf((*MockUnsafeAuthServiceServer)(nil).mustEmbedUnimplementedAuthServiceServer))
}

// MockAuthService_UploadAuthThumbnailServer is a mock of AuthService_UploadAuthThumbnailServer interface.
type MockAuthService_UploadAuthThumbnailServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthService_UploadAuthThumbnailServerMockRecorder
}

// MockAuthService_UploadAuthThumbnailServerMockRecorder is the mock recorder for MockAuthService_UploadAuthThumbnailServer.
type MockAuthService_UploadAuthThumbnailServerMockRecorder struct {
	mock *MockAuthService_UploadAuthThumbnailServer
}

// NewMockAuthService_UploadAuthThumbnailServer creates a new mock instance.
func NewMockAuthService_UploadAuthThumbnailServer(ctrl *gomock.Controller) *MockAuthService_UploadAuthThumbnailServer {
	mock := &MockAuthService_UploadAuthThumbnailServer{ctrl: ctrl}
	mock.recorder = &MockAuthService_UploadAuthThumbnailServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService_UploadAuthThumbnailServer) EXPECT() *MockAuthService_UploadAuthThumbnailServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockAuthService_UploadAuthThumbnailServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockAuthService_UploadAuthThumbnailServer) Recv() (*user.UploadAuthThumbnailRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*user.UploadAuthThumbnailRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockAuthService_UploadAuthThumbnailServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).RecvMsg), m)
}

// SendAndClose mocks base method.
func (m *MockAuthService_UploadAuthThumbnailServer) SendAndClose(arg0 *user.AuthThumbnailResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).SendAndClose), arg0)
}

// SendHeader mocks base method.
func (m *MockAuthService_UploadAuthThumbnailServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockAuthService_UploadAuthThumbnailServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockAuthService_UploadAuthThumbnailServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockAuthService_UploadAuthThumbnailServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockAuthService_UploadAuthThumbnailServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockAuthService_UploadAuthThumbnailServer)(nil).SetTrailer), arg0)
}
