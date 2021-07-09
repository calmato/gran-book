// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/user/service.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	domain "github.com/calmato/gran-book/api/server/user/internal/domain"
	user "github.com/calmato/gran-book/api/server/user/internal/domain/user"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Authentication mocks base method
func (m *MockService) Authentication(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authentication", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authentication indicates an expected call of Authentication
func (mr *MockServiceMockRecorder) Authentication(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockService)(nil).Authentication), ctx)
}

// List mocks base method
func (m *MockService) List(ctx context.Context, q *domain.ListQuery) ([]*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, q)
	ret0, _ := ret[0].([]*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServiceMockRecorder) List(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), ctx, q)
}

// ListFollow mocks base method
func (m *MockService) ListFollow(ctx context.Context, q *domain.ListQuery, uid string) ([]*user.Follow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollow", ctx, q, uid)
	ret0, _ := ret[0].([]*user.Follow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFollow indicates an expected call of ListFollow
func (mr *MockServiceMockRecorder) ListFollow(ctx, q, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollow", reflect.TypeOf((*MockService)(nil).ListFollow), ctx, q, uid)
}

// ListFollower mocks base method
func (m *MockService) ListFollower(ctx context.Context, q *domain.ListQuery, uid string) ([]*user.Follower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollower", ctx, q, uid)
	ret0, _ := ret[0].([]*user.Follower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFollower indicates an expected call of ListFollower
func (mr *MockServiceMockRecorder) ListFollower(ctx, q, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollower", reflect.TypeOf((*MockService)(nil).ListFollower), ctx, q, uid)
}

// ListInstanceID mocks base method
func (m *MockService) ListInstanceID(ctx context.Context, userIDs []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstanceID", ctx, userIDs)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstanceID indicates an expected call of ListInstanceID
func (mr *MockServiceMockRecorder) ListInstanceID(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstanceID", reflect.TypeOf((*MockService)(nil).ListInstanceID), ctx, userIDs)
}

// ListCount mocks base method
func (m *MockService) ListCount(ctx context.Context, q *domain.ListQuery) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCount", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCount indicates an expected call of ListCount
func (mr *MockServiceMockRecorder) ListCount(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCount", reflect.TypeOf((*MockService)(nil).ListCount), ctx, q)
}

// ListFriendCount mocks base method
func (m *MockService) ListFriendCount(ctx context.Context, uid string) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFriendCount", ctx, uid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFriendCount indicates an expected call of ListFriendCount
func (mr *MockServiceMockRecorder) ListFriendCount(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFriendCount", reflect.TypeOf((*MockService)(nil).ListFriendCount), ctx, uid)
}

// Show mocks base method
func (m *MockService) Show(ctx context.Context, uid string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, uid)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show
func (mr *MockServiceMockRecorder) Show(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockService)(nil).Show), ctx, uid)
}

// ShowRelationship mocks base method
func (m *MockService) ShowRelationship(ctx context.Context, id int) (*user.Relationship, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowRelationship", ctx, id)
	ret0, _ := ret[0].(*user.Relationship)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowRelationship indicates an expected call of ShowRelationship
func (mr *MockServiceMockRecorder) ShowRelationship(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowRelationship", reflect.TypeOf((*MockService)(nil).ShowRelationship), ctx, id)
}

// ShowRelationshipByUID mocks base method
func (m *MockService) ShowRelationshipByUID(ctx context.Context, followID, followerID string) (*user.Relationship, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowRelationshipByUID", ctx, followID, followerID)
	ret0, _ := ret[0].(*user.Relationship)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowRelationshipByUID indicates an expected call of ShowRelationshipByUID
func (mr *MockServiceMockRecorder) ShowRelationshipByUID(ctx, followID, followerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowRelationshipByUID", reflect.TypeOf((*MockService)(nil).ShowRelationshipByUID), ctx, followID, followerID)
}

// Create mocks base method
func (m *MockService) Create(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockServiceMockRecorder) Create(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, u)
}

// CreateWithOAuth mocks base method
func (m *MockService) CreateWithOAuth(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithOAuth", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWithOAuth indicates an expected call of CreateWithOAuth
func (mr *MockServiceMockRecorder) CreateWithOAuth(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithOAuth", reflect.TypeOf((*MockService)(nil).CreateWithOAuth), ctx, u)
}

// CreateRelationship mocks base method
func (m *MockService) CreateRelationship(ctx context.Context, r *user.Relationship) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelationship", ctx, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRelationship indicates an expected call of CreateRelationship
func (mr *MockServiceMockRecorder) CreateRelationship(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelationship", reflect.TypeOf((*MockService)(nil).CreateRelationship), ctx, r)
}

// Update mocks base method
func (m *MockService) Update(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockServiceMockRecorder) Update(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), ctx, u)
}

// UpdatePassword mocks base method
func (m *MockService) UpdatePassword(ctx context.Context, uid, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, uid, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockServiceMockRecorder) UpdatePassword(ctx, uid, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockService)(nil).UpdatePassword), ctx, uid, password)
}

// Delete mocks base method
func (m *MockService) Delete(ctx context.Context, uid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServiceMockRecorder) Delete(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), ctx, uid)
}

// DeleteRelationship mocks base method
func (m *MockService) DeleteRelationship(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRelationship", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRelationship indicates an expected call of DeleteRelationship
func (mr *MockServiceMockRecorder) DeleteRelationship(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRelationship", reflect.TypeOf((*MockService)(nil).DeleteRelationship), ctx, id)
}

// UploadThumbnail mocks base method
func (m *MockService) UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadThumbnail", ctx, uid, thumbnail)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadThumbnail indicates an expected call of UploadThumbnail
func (mr *MockServiceMockRecorder) UploadThumbnail(ctx, uid, thumbnail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadThumbnail", reflect.TypeOf((*MockService)(nil).UploadThumbnail), ctx, uid, thumbnail)
}

// IsFriend mocks base method
func (m *MockService) IsFriend(ctx context.Context, friendID, uid string) (bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFriend", ctx, friendID, uid)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IsFriend indicates an expected call of IsFriend
func (mr *MockServiceMockRecorder) IsFriend(ctx, friendID, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFriend", reflect.TypeOf((*MockService)(nil).IsFriend), ctx, friendID, uid)
}

// Validation mocks base method
func (m *MockService) Validation(ctx context.Context, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validation", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validation indicates an expected call of Validation
func (mr *MockServiceMockRecorder) Validation(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validation", reflect.TypeOf((*MockService)(nil).Validation), ctx, u)
}

// ValidationRelationship mocks base method
func (m *MockService) ValidationRelationship(ctx context.Context, r *user.Relationship) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationRelationship", ctx, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidationRelationship indicates an expected call of ValidationRelationship
func (mr *MockServiceMockRecorder) ValidationRelationship(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationRelationship", reflect.TypeOf((*MockService)(nil).ValidationRelationship), ctx, r)
}
