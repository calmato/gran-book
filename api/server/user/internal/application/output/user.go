package output

// UserProfile - ユーザプロフィール関連のレスポンス
type UserProfile struct {
	IsFollow      bool
	IsFollower    bool
	FollowCount   int
	FollowerCount int
}
