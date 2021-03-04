package input

// GetUserProfile - ユーザプロフィールのリクエスト
type GetUserProfile struct {
	ID string `json:"id" validate:"required"`
}
