package output

// GetUserProfile - ユーザプロフィール取得のレスポンス
type GetUserProfile struct {
	IsFollow       bool
	IsFollower     bool
	FollowsTotal   int64
	FollowersTotal int64
}
