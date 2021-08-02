package application

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/array"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"golang.org/x/xerrors"
)

// UserApplication - Userアプリケーションのインターフェース
type UserApplication interface {
	Authentication(ctx context.Context) (*user.User, error)
	Authorization(ctx context.Context) (int, error)
	List(ctx context.Context, q *database.ListQuery) ([]*user.User, int, error)
	ListAdmin(ctx context.Context, q *database.ListQuery) ([]*user.User, int, error)
	ListFollow(ctx context.Context, userID string, limit int, offset int) ([]*user.Follow, int, error)
	ListFollower(ctx context.Context, userID string, limit int, offset int) ([]*user.Follower, int, error)
	MultiGet(ctx context.Context, userIDs []string) ([]*user.User, error)
	Get(ctx context.Context, userID string) (*user.User, error)
	GetAdmin(ctx context.Context, userID string) (*user.User, error)
	GetUserProfile(ctx context.Context, userID string) (*user.User, bool, bool, int, int, error)
	Create(ctx context.Context, u *user.User) error
	Update(ctx context.Context, u *user.User) error
	UpdatePassword(ctx context.Context, u *user.User) error
	Delete(ctx context.Context, u *user.User) error
	DeleteAdmin(ctx context.Context, u *user.User) error
	Follow(ctx context.Context, userID string, followerID string) (*user.User, bool, bool, int, int, error)
	Unfollow(ctx context.Context, userID string, followerID string) (*user.User, bool, bool, int, int, error)
	UploadThumbnail(ctx context.Context, userID string, thumbnail []byte) (string, error)
	HasAdminRole(role int) error
}

type userApplication struct {
	userDomainValidation user.Validation
	userRepository       user.Repository
	userUploader         user.Uploader
}

// NewUserApplication - UserApplicationの生成
func NewUserApplication(udv user.Validation, ur user.Repository, uu user.Uploader) UserApplication {
	return &userApplication{
		userDomainValidation: udv,
		userRepository:       ur,
		userUploader:         uu,
	}
}

func (a *userApplication) Authentication(ctx context.Context) (*user.User, error) {
	userID, err := a.userRepository.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	u, err := a.userRepository.Get(ctx, userID)
	if err == nil {
		return u, nil
	}

	// err: Auth APIにはデータがあるが、User DBにはレコードがない
	// -> Auth APIのデータを基にUser DBに登録
	ou := &user.User{
		ID:     userID,
		Gender: user.UnkownGender,
		Role:   user.UserRole,
	}

	// TODO: domain validation
	err = a.userRepository.CreateWithOAuth(ctx, ou)
	if err != nil {
		return nil, err
	}

	return ou, nil
}

func (a *userApplication) Authorization(ctx context.Context) (int, error) {
	userID, err := a.userRepository.Authentication(ctx)
	if err != nil {
		return user.UserRole, err
	}

	u, err := a.userRepository.Get(ctx, userID)
	if err != nil {
		return user.UserRole, err
	}

	switch u.Role {
	case user.AdminRole, user.DeveloperRole, user.OperatorRole:
		return u.Role, nil
	default:
		err := xerrors.New("This account doesn't have administrator privileges")
		return user.UserRole, exception.Forbidden.New(err)
	}
}

func (a *userApplication) List(ctx context.Context, q *database.ListQuery) ([]*user.User, int, error) {
	us, err := a.userRepository.List(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	total, err := a.userRepository.Count(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return us, total, nil
}

func (a *userApplication) ListAdmin(ctx context.Context, q *database.ListQuery) ([]*user.User, int, error) {
	cq := &database.ConditionQuery{
		Field:    "role",
		Operator: "IN",
		Value:    []int{user.AdminRole, user.DeveloperRole, user.OperatorRole},
	}
	q.Conditions = append(q.Conditions, cq)

	us, err := a.userRepository.List(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	total, err := a.userRepository.Count(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return us, total, nil
}

func (a *userApplication) ListFollow(
	ctx context.Context, userID string, limit int, offset int,
) ([]*user.Follow, int, error) {
	// フォローしているユーザーの一覧を取得
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
		Conditions: []*database.ConditionQuery{
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    userID,
			},
		},
	}

	fs, err := a.userRepository.ListFollow(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	followingUserIDs := make([]string, len(fs))
	for i, f := range fs {
		followingUserIDs[i] = f.FollowID
	}

	// フォローされているかの検証 (フォローしているかの検証はすべてtrueになるため不要)
	followIDs, err := a.userRepository.ListFollowID(ctx, userID, followingUserIDs...)

	for _, f := range fs {
		isFollower, _ := array.Contains(followIDs, f.FollowID)
		f.IsFollowing = true
		f.IsFollowed = isFollower
	}

	total, err := a.userRepository.CountRelationship(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return fs, total, nil
}

func (a *userApplication) ListFollower(
	ctx context.Context, userID string, limit int, offset int,
) ([]*user.Follower, int, error) {
	// フォローされているユーザーの一覧を取得
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
		Conditions: []*database.ConditionQuery{
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    userID,
			},
		},
	}

	fs, err := a.userRepository.ListFollower(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	followedUserIDs := make([]string, len(fs))
	for i, f := range fs {
		followedUserIDs[i] = f.FollowerID
	}

	// フォローしているかの検証 (フォローされているかの検証はすべてtrueになるため不要)
	followerIDs, err := a.userRepository.ListFollowerID(ctx, userID, followedUserIDs...)

	for _, f := range fs {
		isFollow, _ := array.Contains(followerIDs, f.FollowerID)
		f.IsFollowing = isFollow
		f.IsFollowed = true
	}

	total, err := a.userRepository.CountRelationship(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return fs, total, nil
}

func (a *userApplication) MultiGet(ctx context.Context, userIDs []string) ([]*user.User, error) {
	return a.userRepository.MultiGet(ctx, userIDs)
}

func (a *userApplication) Get(ctx context.Context, userID string) (*user.User, error) {
	return a.userRepository.Get(ctx, userID)
}

func (a *userApplication) GetAdmin(ctx context.Context, userID string) (*user.User, error) {
	return a.userRepository.GetAdmin(ctx, userID)
}

func (a *userApplication) GetUserProfile(
	ctx context.Context, userID string,
) (*user.User, bool, bool, int, int, error) {
	u, err := a.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	isFollow, isFollower, followCount, followerCount, err := a.getRelationship(ctx, u.ID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	return u, isFollow, isFollower, followCount, followerCount, nil
}

func (a *userApplication) Create(ctx context.Context, u *user.User) error {
	if u.Gender == 0 {
		u.Gender = user.UnkownGender
	}

	if u.Role == 0 {
		u.Role = user.UserRole
	}

	err := a.userDomainValidation.User(ctx, u)
	if err != nil {
		return err
	}

	current := time.Now().Local()
	u.CreatedAt = current
	u.UpdatedAt = current

	return a.userRepository.Create(ctx, u)
}

func (a *userApplication) Update(ctx context.Context, u *user.User) error {
	err := a.userDomainValidation.User(ctx, u)
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now().Local()

	return a.userRepository.Update(ctx, u)
}

func (a *userApplication) UpdatePassword(ctx context.Context, u *user.User) error {
	err := a.userDomainValidation.User(ctx, u)
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now().Local()

	return a.userRepository.UpdatePassword(ctx, u.ID, u.Password)
}

func (a *userApplication) Delete(ctx context.Context, u *user.User) error {
	return a.userRepository.Delete(ctx, u.ID)
}

func (a *userApplication) DeleteAdmin(ctx context.Context, u *user.User) error {
	u.UpdatedAt = time.Now().Local()
	u.Role = user.UserRole

	return a.userRepository.Update(ctx, u)
}

func (a *userApplication) Follow(
	ctx context.Context, userID string, followerID string,
) (*user.User, bool, bool, int, int, error) {
	fu, err := a.userRepository.Get(ctx, followerID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	current := time.Now().Local()
	r := &user.Relationship{
		FollowID:   userID,
		FollowerID: fu.ID,
		CreatedAt:  current,
		UpdatedAt:  current,
	}

	err = a.userDomainValidation.Relationship(ctx, r)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	err = a.userRepository.CreateRelationship(ctx, r)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	isFollow, isFollower, followCount, followerCount, err := a.getRelationship(ctx, fu.ID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	return fu, isFollow, isFollower, followCount, followerCount, nil
}

func (a *userApplication) Unfollow(
	ctx context.Context, userID string, followerID string,
) (*user.User, bool, bool, int, int, error) {
	fu, err := a.userRepository.Get(ctx, followerID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	relationshipID, err := a.userRepository.GetRelationshipIDByUserID(ctx, userID, followerID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	err = a.userRepository.DeleteRelationship(ctx, relationshipID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	isFollow, isFollower, followCount, followerCount, err := a.getRelationship(ctx, fu.ID)
	if err != nil {
		return nil, false, false, 0, 0, err
	}

	return fu, isFollow, isFollower, followCount, followerCount, nil
}

func (a *userApplication) UploadThumbnail(ctx context.Context, userID string, thumbnail []byte) (string, error) {
	return a.userUploader.Thumbnail(ctx, userID, thumbnail)
}

func (a *userApplication) HasAdminRole(role int) error {
	if role != user.AdminRole {
		err := xerrors.New("This account doesn't have administrator privileges")
		return exception.Forbidden.New(err)
	}

	return nil
}

func (a *userApplication) getRelationship(ctx context.Context, userID string) (bool, bool, int, int, error) {
	followIDs, err := a.userRepository.ListFollowID(ctx, userID)
	if err != nil {
		return false, false, 0, 0, err
	}

	followerIDs, err := a.userRepository.ListFollowID(ctx, userID)
	if err != nil {
		return false, false, 0, 0, err
	}

	isFollow, _ := array.Contains(followIDs, userID)
	isFollower, _ := array.Contains(followerIDs, userID)

	followCount := len(followIDs)
	followerCount := len(followerIDs)

	return isFollow, isFollower, followCount, followerCount, nil
}
