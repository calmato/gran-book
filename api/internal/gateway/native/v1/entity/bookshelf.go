package entity

import "github.com/calmato/gran-book/api/internal/gateway/entity"

type Bookshelf struct {
	ID         int64  `json:"id"`                   // 本棚ID
	Status     string `json:"status"`               // 読書ステータス
	ReadOn     string `json:"readOn"`               // 読み終えた日
	Impression string `json:"impression,omitempty"` // 感想
	CreatedAt  string `json:"createdAt"`            // 登録日時
	UpdatedAt  string `json:"updatedAt"`            // 更新日時
}

func NewBookshelf(bs *entity.Bookshelf) *Bookshelf {
	if bs == nil {
		return nil
	}

	return &Bookshelf{
		ID:         bs.Id,
		Status:     bs.Status().Name(),
		Impression: "",
		ReadOn:     bs.ReadOn,
		CreatedAt:  bs.CreatedAt,
		UpdatedAt:  bs.UpdatedAt,
	}
}
