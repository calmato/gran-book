package inquiry

import (
	"time"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/proto/information"
)

// Inquiry - お問い合わせ
type Inquiry struct {
	ID          int64     `gorm:"default:null;primaryKey;autoIncrement;<-:create"`
	SenderID    string    `gorm:"default:null"`
	AdminID     string    `gorm:"default:null"`
	Subject     string    `gorm:"default:null"`
	Description string    `gorm:"default:null"`
	Email       string    `gorm:"default:null"`
	IsReplied   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"default:null;<-:create"`
	UpdatedAt   time.Time `gorm:"default:null"`
}

type Inquiries []*Inquiry

func (i *Inquiry) Proto() *information.Inquiry {
	return &information.Inquiry{
		Id:          i.ID,
		SenderId:    i.SenderID,
		AdminId:     i.AdminID,
		Subject:     i.Subject,
		Description: i.Description,
		Email:       i.Email,
		IsReplied:   i.IsReplied,
		CreatedAt:   datetime.FormatTime(i.CreatedAt),
		UpdatedAt:   datetime.FormatTime(i.UpdatedAt),
	}
}

func (is Inquiries) Proto() []*information.Inquiry {
	res := make([]*information.Inquiry, len(is))
	for i := range is {
		res[i] = is[i].Proto()
	}
	return res
}
