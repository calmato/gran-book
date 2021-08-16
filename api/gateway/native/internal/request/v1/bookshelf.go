package v1

type ReadBookshelfRequest struct {
	ReadOn     string `json:"readOn"`
	Impression string `json:"impression"`
}
