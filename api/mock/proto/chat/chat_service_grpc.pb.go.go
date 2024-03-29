// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/chat/chat_service_grpc.pb.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	context "context"
	reflect "reflect"

	chat "github.com/calmato/gran-book/api/proto/chat"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockChatServiceClient is a mock of ChatServiceClient interface.
type MockChatServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceClientMockRecorder
}

// MockChatServiceClientMockRecorder is the mock recorder for MockChatServiceClient.
type MockChatServiceClientMockRecorder struct {
	mock *MockChatServiceClient
}

// NewMockChatServiceClient creates a new mock instance.
func NewMockChatServiceClient(ctrl *gomock.Controller) *MockChatServiceClient {
	mock := &MockChatServiceClient{ctrl: ctrl}
	mock.recorder = &MockChatServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceClient) EXPECT() *MockChatServiceClientMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockChatServiceClient) CreateMessage(ctx context.Context, in *chat.CreateMessageRequest, opts ...grpc.CallOption) (*chat.MessageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMessage", varargs...)
	ret0, _ := ret[0].(*chat.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockChatServiceClientMockRecorder) CreateMessage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockChatServiceClient)(nil).CreateMessage), varargs...)
}

// CreateRoom mocks base method.
func (m *MockChatServiceClient) CreateRoom(ctx context.Context, in *chat.CreateRoomRequest, opts ...grpc.CallOption) (*chat.RoomResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRoom", varargs...)
	ret0, _ := ret[0].(*chat.RoomResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockChatServiceClientMockRecorder) CreateRoom(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockChatServiceClient)(nil).CreateRoom), varargs...)
}

// ListRoom mocks base method.
func (m *MockChatServiceClient) ListRoom(ctx context.Context, in *chat.ListRoomRequest, opts ...grpc.CallOption) (*chat.RoomListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoom", varargs...)
	ret0, _ := ret[0].(*chat.RoomListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockChatServiceClientMockRecorder) ListRoom(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockChatServiceClient)(nil).ListRoom), varargs...)
}

// UploadImage mocks base method.
func (m *MockChatServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (chat.ChatService_UploadImageClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadImage", varargs...)
	ret0, _ := ret[0].(chat.ChatService_UploadImageClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockChatServiceClientMockRecorder) UploadImage(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockChatServiceClient)(nil).UploadImage), varargs...)
}

// MockChatService_UploadImageClient is a mock of ChatService_UploadImageClient interface.
type MockChatService_UploadImageClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatService_UploadImageClientMockRecorder
}

// MockChatService_UploadImageClientMockRecorder is the mock recorder for MockChatService_UploadImageClient.
type MockChatService_UploadImageClientMockRecorder struct {
	mock *MockChatService_UploadImageClient
}

// NewMockChatService_UploadImageClient creates a new mock instance.
func NewMockChatService_UploadImageClient(ctrl *gomock.Controller) *MockChatService_UploadImageClient {
	mock := &MockChatService_UploadImageClient{ctrl: ctrl}
	mock.recorder = &MockChatService_UploadImageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatService_UploadImageClient) EXPECT() *MockChatService_UploadImageClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockChatService_UploadImageClient) CloseAndRecv() (*chat.MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*chat.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockChatService_UploadImageClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockChatService_UploadImageClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockChatService_UploadImageClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockChatService_UploadImageClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockChatService_UploadImageClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).Context))
}

// Header mocks base method.
func (m *MockChatService_UploadImageClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockChatService_UploadImageClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m_2 *MockChatService_UploadImageClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockChatService_UploadImageClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockChatService_UploadImageClient) Send(arg0 *chat.UploadChatImageRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockChatService_UploadImageClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockChatService_UploadImageClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockChatService_UploadImageClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockChatService_UploadImageClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockChatService_UploadImageClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockChatService_UploadImageClient)(nil).Trailer))
}

// MockChatServiceServer is a mock of ChatServiceServer interface.
type MockChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceServerMockRecorder
}

// MockChatServiceServerMockRecorder is the mock recorder for MockChatServiceServer.
type MockChatServiceServerMockRecorder struct {
	mock *MockChatServiceServer
}

// NewMockChatServiceServer creates a new mock instance.
func NewMockChatServiceServer(ctrl *gomock.Controller) *MockChatServiceServer {
	mock := &MockChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceServer) EXPECT() *MockChatServiceServerMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockChatServiceServer) CreateMessage(arg0 context.Context, arg1 *chat.CreateMessageRequest) (*chat.MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", arg0, arg1)
	ret0, _ := ret[0].(*chat.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockChatServiceServerMockRecorder) CreateMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockChatServiceServer)(nil).CreateMessage), arg0, arg1)
}

// CreateRoom mocks base method.
func (m *MockChatServiceServer) CreateRoom(arg0 context.Context, arg1 *chat.CreateRoomRequest) (*chat.RoomResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", arg0, arg1)
	ret0, _ := ret[0].(*chat.RoomResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockChatServiceServerMockRecorder) CreateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockChatServiceServer)(nil).CreateRoom), arg0, arg1)
}

// ListRoom mocks base method.
func (m *MockChatServiceServer) ListRoom(arg0 context.Context, arg1 *chat.ListRoomRequest) (*chat.RoomListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", arg0, arg1)
	ret0, _ := ret[0].(*chat.RoomListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockChatServiceServerMockRecorder) ListRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockChatServiceServer)(nil).ListRoom), arg0, arg1)
}

// UploadImage mocks base method.
func (m *MockChatServiceServer) UploadImage(arg0 chat.ChatService_UploadImageServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadImage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockChatServiceServerMockRecorder) UploadImage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockChatServiceServer)(nil).UploadImage), arg0)
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}

// MockUnsafeChatServiceServer is a mock of UnsafeChatServiceServer interface.
type MockUnsafeChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeChatServiceServerMockRecorder
}

// MockUnsafeChatServiceServerMockRecorder is the mock recorder for MockUnsafeChatServiceServer.
type MockUnsafeChatServiceServerMockRecorder struct {
	mock *MockUnsafeChatServiceServer
}

// NewMockUnsafeChatServiceServer creates a new mock instance.
func NewMockUnsafeChatServiceServer(ctrl *gomock.Controller) *MockUnsafeChatServiceServer {
	mock := &MockUnsafeChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeChatServiceServer) EXPECT() *MockUnsafeChatServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockUnsafeChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockUnsafeChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockUnsafeChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}

// MockChatService_UploadImageServer is a mock of ChatService_UploadImageServer interface.
type MockChatService_UploadImageServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatService_UploadImageServerMockRecorder
}

// MockChatService_UploadImageServerMockRecorder is the mock recorder for MockChatService_UploadImageServer.
type MockChatService_UploadImageServerMockRecorder struct {
	mock *MockChatService_UploadImageServer
}

// NewMockChatService_UploadImageServer creates a new mock instance.
func NewMockChatService_UploadImageServer(ctrl *gomock.Controller) *MockChatService_UploadImageServer {
	mock := &MockChatService_UploadImageServer{ctrl: ctrl}
	mock.recorder = &MockChatService_UploadImageServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatService_UploadImageServer) EXPECT() *MockChatService_UploadImageServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockChatService_UploadImageServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockChatService_UploadImageServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockChatService_UploadImageServer) Recv() (*chat.UploadChatImageRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*chat.UploadChatImageRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockChatService_UploadImageServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockChatService_UploadImageServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockChatService_UploadImageServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).RecvMsg), m)
}

// SendAndClose mocks base method.
func (m *MockChatService_UploadImageServer) SendAndClose(arg0 *chat.MessageResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockChatService_UploadImageServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).SendAndClose), arg0)
}

// SendHeader mocks base method.
func (m *MockChatService_UploadImageServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockChatService_UploadImageServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockChatService_UploadImageServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockChatService_UploadImageServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockChatService_UploadImageServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockChatService_UploadImageServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockChatService_UploadImageServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockChatService_UploadImageServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockChatService_UploadImageServer)(nil).SetTrailer), arg0)
}
