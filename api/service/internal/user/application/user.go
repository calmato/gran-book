package application

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/service/internal/user/domain/user"
	"github.com/calmato/gran-book/api/service/pkg/array"
	"github.com/calmato/gran-book/api/service/pkg/database"
	"github.com/google/uuid"
)

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
	current := time.Now().Local()
	ou := &user.User{
		ID:        userID,
		Gender:    user.UnkownGender,
		Role:      user.UserRole,
		CreatedAt: current,
		UpdatedAt: current,
	}

	err = a.userRepository.CreateWithOAuth(ctx, ou)
	if err != nil {
		return nil, err
	}

	return ou, nil
}

func (a *userApplication) List(ctx context.Context, q *database.ListQuery) (user.Users, int, error) {
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

func (a *userApplication) ListAdmin(ctx context.Context, q *database.ListQuery) (user.Users, int, error) {
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
	ctx context.Context, currentUserID, targetUserID string, limit int, offset int,
) (user.Follows, int, error) {
	// フォローしているユーザーの一覧を取得
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
		Conditions: []*database.ConditionQuery{
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    targetUserID,
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

	// ログインユーザーのフォロー/フォローワー検証
	followIDs, err := a.userRepository.ListFollowID(ctx, currentUserID, followingUserIDs...)
	if err != nil {
		return nil, 0, err
	}

	followerIDs, err := a.userRepository.ListFollowerID(ctx, currentUserID, followingUserIDs...)
	if err != nil {
		return nil, 0, err
	}

	for _, f := range fs {
		isFollowing, _ := array.Contains(followerIDs, f.FollowID)
		isFollowed, _ := array.Contains(followIDs, f.FollowID)

		f.IsFollowing = isFollowing
		f.IsFollowed = isFollowed
	}

	total, err := a.userRepository.CountRelationship(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return fs, total, nil
}

func (a *userApplication) ListFollower(
	ctx context.Context, currentUserID, targetUserID string, limit int, offset int,
) (user.Followers, int, error) {
	// フォローされているユーザーの一覧を取得
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
		Conditions: []*database.ConditionQuery{
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    targetUserID,
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

	// ログインユーザーのフォロー/フォローワー検証
	followIDs, err := a.userRepository.ListFollowID(ctx, currentUserID, followedUserIDs...)
	if err != nil {
		return nil, 0, err
	}

	followerIDs, err := a.userRepository.ListFollowerID(ctx, currentUserID, followedUserIDs...)
	if err != nil {
		return nil, 0, err
	}

	for _, f := range fs {
		isFollowing, _ := array.Contains(followerIDs, f.FollowerID)
		isFollowed, _ := array.Contains(followIDs, f.FollowerID)

		f.IsFollowing = isFollowing
		f.IsFollowed = isFollowed
	}

	total, err := a.userRepository.CountRelationship(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return fs, total, nil
}

func (a *userApplication) MultiGet(ctx context.Context, userIDs []string) (user.Users, error) {
	return a.userRepository.MultiGet(ctx, userIDs)
}

func (a *userApplication) Get(ctx context.Context, userID string) (*user.User, error) {
	return a.userRepository.Get(ctx, userID)
}

func (a *userApplication) GetAdmin(ctx context.Context, userID string) (*user.User, error) {
	return a.userRepository.GetAdmin(ctx, userID)
}

func (a *userApplication) GetUserProfile(ctx context.Context, userID, targetID string) (*user.User, error) {
	u, err := a.userRepository.Get(ctx, targetID)
	if err != nil {
		return nil, err
	}

	isFollowing, isFollowed, followCount, followerCount, err := a.getRelationship(ctx, userID, u.ID)
	if err != nil {
		return nil, err
	}

	u.IsFollowing = isFollowing
	u.IsFollowed = isFollowed
	u.FollowCount = followCount
	u.FollowerCount = followerCount

	return u, nil
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
	u.ID = uuid.New().String()

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
	err := a.userDomainValidation.User(ctx, u)
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now().Local()
	u.Role = user.UserRole

	return a.userRepository.Update(ctx, u)
}

func (a *userApplication) Follow(ctx context.Context, userID string, followerID string) (*user.User, error) {
	fu, err := a.userRepository.Get(ctx, followerID)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	err = a.userRepository.CreateRelationship(ctx, r)
	if err != nil {
		return nil, err
	}

	isFollowing, isFollowed, followCount, followerCount, err := a.getRelationship(ctx, userID, fu.ID)
	if err != nil {
		return nil, err
	}

	fu.IsFollowing = isFollowing
	fu.IsFollowed = isFollowed
	fu.FollowCount = followCount
	fu.FollowerCount = followerCount

	return fu, nil
}

func (a *userApplication) Unfollow(ctx context.Context, userID string, followerID string) (*user.User, error) {
	fu, err := a.userRepository.Get(ctx, followerID)
	if err != nil {
		return nil, err
	}

	relationshipID, err := a.userRepository.GetRelationshipIDByUserID(ctx, userID, fu.ID)
	if err != nil {
		return nil, err
	}

	err = a.userRepository.DeleteRelationship(ctx, relationshipID)
	if err != nil {
		return nil, err
	}

	isFollowing, isFollowed, followCount, followerCount, err := a.getRelationship(ctx, userID, fu.ID)
	if err != nil {
		return nil, err
	}

	fu.IsFollowing = isFollowing
	fu.IsFollowed = isFollowed
	fu.FollowCount = followCount
	fu.FollowerCount = followerCount

	return fu, nil
}

func (a *userApplication) UploadThumbnail(ctx context.Context, userID string, thumbnail []byte) (string, error) {
	return a.userUploader.Thumbnail(ctx, userID, thumbnail)
}

func (a *userApplication) getRelationship(
	ctx context.Context, currentUserID, targetUserID string,
) (isFollowing, isFollowed bool, followCount, followerCount int, err error) {
	var followIDs, followerIDs []string
	followIDs, err = a.userRepository.ListFollowID(ctx, targetUserID)
	if err != nil {
		return
	}

	followerIDs, err = a.userRepository.ListFollowerID(ctx, targetUserID)
	if err != nil {
		return
	}

	isFollowing, _ = array.Contains(followIDs, currentUserID)
	isFollowed, _ = array.Contains(followerIDs, currentUserID)

	followCount = len(followIDs)
	followerCount = len(followerIDs)

	return
}
