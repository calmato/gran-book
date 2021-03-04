package output

// GetUserProfile - ユーザプロフィール取得のレスポンス
type GetUserProfile struct {
	Rating         float32
	ReviewCount    int32
	IsFollow       bool
	IsFollower     bool
	FollowsTotal   int64
	FollowersTotal int64
}
