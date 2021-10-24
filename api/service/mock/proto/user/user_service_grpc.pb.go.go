// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/user/user_service_grpc.pb.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	user "github.com/calmato/gran-book/api/service/proto/user"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// Follow mocks base method.
func (m *MockUserServiceClient) Follow(ctx context.Context, in *user.FollowRequest, opts ...grpc.CallOption) (*user.UserProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Follow", varargs...)
	ret0, _ := ret[0].(*user.UserProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Follow indicates an expected call of Follow.
func (mr *MockUserServiceClientMockRecorder) Follow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockUserServiceClient)(nil).Follow), varargs...)
}

// GetUser mocks base method.
func (m *MockUserServiceClient) GetUser(ctx context.Context, in *user.GetUserRequest, opts ...grpc.CallOption) (*user.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*user.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceClient)(nil).GetUser), varargs...)
}

// GetUserProfile mocks base method.
func (m *MockUserServiceClient) GetUserProfile(ctx context.Context, in *user.GetUserProfileRequest, opts ...grpc.CallOption) (*user.UserProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserProfile", varargs...)
	ret0, _ := ret[0].(*user.UserProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockUserServiceClientMockRecorder) GetUserProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockUserServiceClient)(nil).GetUserProfile), varargs...)
}

// ListFollow mocks base method.
func (m *MockUserServiceClient) ListFollow(ctx context.Context, in *user.ListFollowRequest, opts ...grpc.CallOption) (*user.FollowListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFollow", varargs...)
	ret0, _ := ret[0].(*user.FollowListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFollow indicates an expected call of ListFollow.
func (mr *MockUserServiceClientMockRecorder) ListFollow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollow", reflect.TypeOf((*MockUserServiceClient)(nil).ListFollow), varargs...)
}

// ListFollower mocks base method.
func (m *MockUserServiceClient) ListFollower(ctx context.Context, in *user.ListFollowerRequest, opts ...grpc.CallOption) (*user.FollowerListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFollower", varargs...)
	ret0, _ := ret[0].(*user.FollowerListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFollower indicates an expected call of ListFollower.
func (mr *MockUserServiceClientMockRecorder) ListFollower(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollower", reflect.TypeOf((*MockUserServiceClient)(nil).ListFollower), varargs...)
}

// ListUser mocks base method.
func (m *MockUserServiceClient) ListUser(ctx context.Context, in *user.ListUserRequest, opts ...grpc.CallOption) (*user.UserListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUser", varargs...)
	ret0, _ := ret[0].(*user.UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser indicates an expected call of ListUser.
func (mr *MockUserServiceClientMockRecorder) ListUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockUserServiceClient)(nil).ListUser), varargs...)
}

// MultiGetUser mocks base method.
func (m *MockUserServiceClient) MultiGetUser(ctx context.Context, in *user.MultiGetUserRequest, opts ...grpc.CallOption) (*user.UserListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGetUser", varargs...)
	ret0, _ := ret[0].(*user.UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetUser indicates an expected call of MultiGetUser.
func (mr *MockUserServiceClientMockRecorder) MultiGetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetUser", reflect.TypeOf((*MockUserServiceClient)(nil).MultiGetUser), varargs...)
}

// Unfollow mocks base method.
func (m *MockUserServiceClient) Unfollow(ctx context.Context, in *user.UnfollowRequest, opts ...grpc.CallOption) (*user.UserProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unfollow", varargs...)
	ret0, _ := ret[0].(*user.UserProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow.
func (mr *MockUserServiceClientMockRecorder) Unfollow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockUserServiceClient)(nil).Unfollow), varargs...)
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// Follow mocks base method.
func (m *MockUserServiceServer) Follow(arg0 context.Context, arg1 *user.FollowRequest) (*user.UserProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", arg0, arg1)
	ret0, _ := ret[0].(*user.UserProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Follow indicates an expected call of Follow.
func (mr *MockUserServiceServerMockRecorder) Follow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockUserServiceServer)(nil).Follow), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockUserServiceServer) GetUser(arg0 context.Context, arg1 *user.GetUserRequest) (*user.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*user.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceServer)(nil).GetUser), arg0, arg1)
}

// GetUserProfile mocks base method.
func (m *MockUserServiceServer) GetUserProfile(arg0 context.Context, arg1 *user.GetUserProfileRequest) (*user.UserProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", arg0, arg1)
	ret0, _ := ret[0].(*user.UserProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockUserServiceServerMockRecorder) GetUserProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockUserServiceServer)(nil).GetUserProfile), arg0, arg1)
}

// ListFollow mocks base method.
func (m *MockUserServiceServer) ListFollow(arg0 context.Context, arg1 *user.ListFollowRequest) (*user.FollowListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollow", arg0, arg1)
	ret0, _ := ret[0].(*user.FollowListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFollow indicates an expected call of ListFollow.
func (mr *MockUserServiceServerMockRecorder) ListFollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollow", reflect.TypeOf((*MockUserServiceServer)(nil).ListFollow), arg0, arg1)
}

// ListFollower mocks base method.
func (m *MockUserServiceServer) ListFollower(arg0 context.Context, arg1 *user.ListFollowerRequest) (*user.FollowerListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollower", arg0, arg1)
	ret0, _ := ret[0].(*user.FollowerListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFollower indicates an expected call of ListFollower.
func (mr *MockUserServiceServerMockRecorder) ListFollower(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollower", reflect.TypeOf((*MockUserServiceServer)(nil).ListFollower), arg0, arg1)
}

// ListUser mocks base method.
func (m *MockUserServiceServer) ListUser(arg0 context.Context, arg1 *user.ListUserRequest) (*user.UserListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", arg0, arg1)
	ret0, _ := ret[0].(*user.UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser indicates an expected call of ListUser.
func (mr *MockUserServiceServerMockRecorder) ListUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockUserServiceServer)(nil).ListUser), arg0, arg1)
}

// MultiGetUser mocks base method.
func (m *MockUserServiceServer) MultiGetUser(arg0 context.Context, arg1 *user.MultiGetUserRequest) (*user.UserListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetUser", arg0, arg1)
	ret0, _ := ret[0].(*user.UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetUser indicates an expected call of MultiGetUser.
func (mr *MockUserServiceServerMockRecorder) MultiGetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetUser", reflect.TypeOf((*MockUserServiceServer)(nil).MultiGetUser), arg0, arg1)
}

// Unfollow mocks base method.
func (m *MockUserServiceServer) Unfollow(arg0 context.Context, arg1 *user.UnfollowRequest) (*user.UserProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", arg0, arg1)
	ret0, _ := ret[0].(*user.UserProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow.
func (mr *MockUserServiceServerMockRecorder) Unfollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockUserServiceServer)(nil).Unfollow), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}