package input

// Notification - お知らせ登録/更新のリクエスト
type CreateNotification struct {
	Title       string `json:"title" validate:"required,max=64"`
	Description string `json:"description" validate:"required,max=2000"`
	Importance  string `json:"importance" validate:"required,max=64"`
}
