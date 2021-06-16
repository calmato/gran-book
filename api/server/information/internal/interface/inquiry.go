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
		SenderId:    req.GetSenderId(),
		Subject:     req.GetSubject(),
		Description: req.GetDescription(),
		Email:       req.GetEmail(),
	}

	ir, err := s.InquiryApplication.Create(ctx, in)

	res := &pb.InquiryResponse{
		InquiryId:   int64(ir.ID),
		SenderId:    ir.SenderId,
		AdminId:     ir.AdminId,
		Subject:     ir.Subject,
		Description: ir.Description,
		Email:       ir.Email,
		IsReplied:   ir.IsReplied,
		CreatedAt:   datetime.DateToString(ir.CreatedAt),
		UpdatedAt:   datetime.DateToString(ir.UpdatedAt),
	}

	return res, err
}
