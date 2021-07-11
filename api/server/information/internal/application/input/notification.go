package input

// Notification - お知らせ登録のリクエスト
type CreateNotification struct {
	Title       string `json:"title" validate:"required,max=64"`
	Description string `json:"description" validate:"required,max=2000"`
	Importance  string `json:"importance" validate:"required,max=64"`
}

// Notification - お知らせの更新のリクエスト
type UpdaeteNotification struct {
	Title       string `json:"title" validate:"required,max=64"`
	Description string `json:"description" validate:"required,max=2000"`
	Importance  string `json:"importance" validate:"required,max=64"`
}

// Notifcation - お知らせ詳細の取得のリクエスト
type ShowNotification struct {
	ID int `json:"id" validate:"required"`
}

// Notification - お知らせの削除のリクエスト
type DeleteNotification struct {
	ID int `json:"id" validate:"required"`
}
