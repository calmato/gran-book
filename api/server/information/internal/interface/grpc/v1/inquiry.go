package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/application"
	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/information/proto"
)

type InquiryServer struct {
	pb.UnimplementedInquiryServiceServer
	InquiryApplication application.InquiryApplication
}

// CreateInquiry お問い合わせ登録
func (s *InquiryServer) CreateInquiry(ctx context.Context, req *pb.CreateInquiryRequest) (*pb.InquiryResponse, error) {
	in := &input.CreateInquiry{
		SenderID:    req.GetSenderId(),
		Subject:     req.GetSubject(),
		Description: req.GetDescription(),
		Email:       req.GetEmail(),
	}

	ir, err := s.InquiryApplication.Create(ctx, in)

	if err != nil {
		return nil, errorHandling(err)
	}

	res := &pb.InquiryResponse{
		InquiryId: int64(ir.ID),
		SenderId:  ir.SenderID,
		// AdminId:     ir.AdminID,
		Subject:     ir.Subject,
		Description: ir.Description,
		Email:       ir.Email,
		IsReplied:   ir.IsReplied,
		CreatedAt:   datetime.DateToString(ir.CreatedAt),
		UpdatedAt:   datetime.DateToString(ir.UpdatedAt),
	}

	return res, err
}
