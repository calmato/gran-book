package v1

// 読んだ本の登録
type ReadBookshelfRequest struct {
	ReadOn     string `json:"readOn"`     // 読み終えた日
	Impression string `json:"impression"` // 感想
}
