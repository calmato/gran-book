// Code generated by MockGen. DO NOT EDIT.
// Source: application.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	context "context"
	reflect "reflect"

	chat "github.com/calmato/gran-book/api/internal/user/domain/chat"
	user "github.com/calmato/gran-book/api/internal/user/domain/user"
	database "github.com/calmato/gran-book/api/pkg/database"
	firestore "github.com/calmato/gran-book/api/pkg/firebase/firestore"
	gomock "github.com/golang/mock/gomock"
)

// MockUserApplication is a mock of UserApplication interface.
type MockUserApplication struct {
	ctrl     *gomock.Controller
	recorder *MockUserApplicationMockRecorder
}

// MockUserApplicationMockRecorder is the mock recorder for MockUserApplication.
type MockUserApplicationMockRecorder struct {
	mock *MockUserApplication
}

// NewMockUserApplication creates a new mock instance.
func NewMockUserApplication(ctrl *gomock.Controller) *MockUserApplication {
	mock := &MockUserApplication{ctrl: ctrl}
	mock.recorder = &MockUserApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserApplication) EXPECT() *MockUserApplicationMockRecorder {
	return m.recorder
}

// Authentication mocks base method.
func (m *MockUserApplication) Authentication(ctx context.Context) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authentication", ctx)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authentication indicates an expected call of Authentication.
func (mr *MockUserApplicationMockRecorder) Authentication(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockUserApplication)(nil).Authentication), ctx)
}

// Create mocks base method.
func (m *MockUserApplication) Create(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserApplicationMockRecorder) Create(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserApplication)(nil).Create), ctx, u)
}

// Delete mocks base method.
func (m *MockUserApplication) Delete(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserApplicationMockRecorder) Delete(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserApplication)(nil).Delete), ctx, u)
}

// DeleteAdmin mocks base method.
func (m *MockUserApplication) DeleteAdmin(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockUserApplicationMockRecorder) DeleteAdmin(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockUserApplication)(nil).DeleteAdmin), ctx, u)
}

// Follow mocks base method.
func (m *MockUserApplication) Follow(ctx context.Context, userID, followerID string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", ctx, userID, followerID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Follow indicates an expected call of Follow.
func (mr *MockUserApplicationMockRecorder) Follow(ctx, userID, followerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockUserApplication)(nil).Follow), ctx, userID, followerID)
}

// Get mocks base method.
func (m *MockUserApplication) Get(ctx context.Context, userID string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserApplicationMockRecorder) Get(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserApplication)(nil).Get), ctx, userID)
}

// GetAdmin mocks base method.
func (m *MockUserApplication) GetAdmin(ctx context.Context, userID string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", ctx, userID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockUserApplicationMockRecorder) GetAdmin(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockUserApplication)(nil).GetAdmin), ctx, userID)
}

// GetUserProfile mocks base method.
func (m *MockUserApplication) GetUserProfile(ctx context.Context, userID, targetID string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", ctx, userID, targetID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockUserApplicationMockRecorder) GetUserProfile(ctx, userID, targetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockUserApplication)(nil).GetUserProfile), ctx, userID, targetID)
}

// List mocks base method.
func (m *MockUserApplication) List(ctx context.Context, q *database.ListQuery) (user.Users, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, q)
	ret0, _ := ret[0].(user.Users)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockUserApplicationMockRecorder) List(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserApplication)(nil).List), ctx, q)
}

// ListAdmin mocks base method.
func (m *MockUserApplication) ListAdmin(ctx context.Context, q *database.ListQuery) (user.Users, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdmin", ctx, q)
	ret0, _ := ret[0].(user.Users)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAdmin indicates an expected call of ListAdmin.
func (mr *MockUserApplicationMockRecorder) ListAdmin(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdmin", reflect.TypeOf((*MockUserApplication)(nil).ListAdmin), ctx, q)
}

// ListFollow mocks base method.
func (m *MockUserApplication) ListFollow(ctx context.Context, userID, targetID string, limit, offset int) (user.Follows, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollow", ctx, userID, targetID, limit, offset)
	ret0, _ := ret[0].(user.Follows)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFollow indicates an expected call of ListFollow.
func (mr *MockUserApplicationMockRecorder) ListFollow(ctx, userID, targetID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollow", reflect.TypeOf((*MockUserApplication)(nil).ListFollow), ctx, userID, targetID, limit, offset)
}

// ListFollower mocks base method.
func (m *MockUserApplication) ListFollower(ctx context.Context, userID, targetID string, limit, offset int) (user.Followers, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollower", ctx, userID, targetID, limit, offset)
	ret0, _ := ret[0].(user.Followers)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFollower indicates an expected call of ListFollower.
func (mr *MockUserApplicationMockRecorder) ListFollower(ctx, userID, targetID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollower", reflect.TypeOf((*MockUserApplication)(nil).ListFollower), ctx, userID, targetID, limit, offset)
}

// MultiGet mocks base method.
func (m *MockUserApplication) MultiGet(ctx context.Context, userIDs []string) (user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGet", ctx, userIDs)
	ret0, _ := ret[0].(user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockUserApplicationMockRecorder) MultiGet(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockUserApplication)(nil).MultiGet), ctx, userIDs)
}

// Unfollow mocks base method.
func (m *MockUserApplication) Unfollow(ctx context.Context, userID, followerID string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", ctx, userID, followerID)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow.
func (mr *MockUserApplicationMockRecorder) Unfollow(ctx, userID, followerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockUserApplication)(nil).Unfollow), ctx, userID, followerID)
}

// Update mocks base method.
func (m *MockUserApplication) Update(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserApplicationMockRecorder) Update(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserApplication)(nil).Update), ctx, u)
}

// UpdatePassword mocks base method.
func (m *MockUserApplication) UpdatePassword(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserApplicationMockRecorder) UpdatePassword(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserApplication)(nil).UpdatePassword), ctx, u)
}

// UploadThumbnail mocks base method.
func (m *MockUserApplication) UploadThumbnail(ctx context.Context, userID string, thumbnail []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadThumbnail", ctx, userID, thumbnail)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadThumbnail indicates an expected call of UploadThumbnail.
func (mr *MockUserApplicationMockRecorder) UploadThumbnail(ctx, userID, thumbnail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadThumbnail", reflect.TypeOf((*MockUserApplication)(nil).UploadThumbnail), ctx, userID, thumbnail)
}

// MockChatApplication is a mock of ChatApplication interface.
type MockChatApplication struct {
	ctrl     *gomock.Controller
	recorder *MockChatApplicationMockRecorder
}

// MockChatApplicationMockRecorder is the mock recorder for MockChatApplication.
type MockChatApplicationMockRecorder struct {
	mock *MockChatApplication
}

// NewMockChatApplication creates a new mock instance.
func NewMockChatApplication(ctrl *gomock.Controller) *MockChatApplication {
	mock := &MockChatApplication{ctrl: ctrl}
	mock.recorder = &MockChatApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatApplication) EXPECT() *MockChatApplicationMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockChatApplication) CreateMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", ctx, cr, cm)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockChatApplicationMockRecorder) CreateMessage(ctx, cr, cm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockChatApplication)(nil).CreateMessage), ctx, cr, cm)
}

// CreateRoom mocks base method.
func (m *MockChatApplication) CreateRoom(ctx context.Context, cr *chat.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, cr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockChatApplicationMockRecorder) CreateRoom(ctx, cr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockChatApplication)(nil).CreateRoom), ctx, cr)
}

// GetRoom mocks base method.
func (m *MockChatApplication) GetRoom(ctx context.Context, roomID, userID string) (*chat.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, roomID, userID)
	ret0, _ := ret[0].(*chat.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockChatApplicationMockRecorder) GetRoom(ctx, roomID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockChatApplication)(nil).GetRoom), ctx, roomID, userID)
}

// ListRoom mocks base method.
func (m *MockChatApplication) ListRoom(ctx context.Context, userID string, p *firestore.Params) (chat.Rooms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", ctx, userID, p)
	ret0, _ := ret[0].(chat.Rooms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockChatApplicationMockRecorder) ListRoom(ctx, userID, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockChatApplication)(nil).ListRoom), ctx, userID, p)
}

// UploadImage mocks base method.
func (m *MockChatApplication) UploadImage(ctx context.Context, cr *chat.Room, image []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadImage", ctx, cr, image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockChatApplicationMockRecorder) UploadImage(ctx, cr, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockChatApplication)(nil).UploadImage), ctx, cr, image)
}
