package server

import (
	"context"

	"github.com/calmato/gran-book/api/internal/information/application"
	"github.com/calmato/gran-book/api/internal/information/domain/inquiry"
	"github.com/calmato/gran-book/api/internal/information/interface/validation"
	pb "github.com/calmato/gran-book/api/proto/information"
)

type inquiryServer struct {
	pb.UnimplementedInquiryServiceServer
	inquiryRequestValidation validation.InquiryRequestValidation
	inquiryApplication       application.InquiryApplication
}

func NewInquiryServer(
	irv validation.InquiryRequestValidation, ia application.InquiryApplication,
) pb.InquiryServiceServer {
	return &inquiryServer{
		inquiryRequestValidation: irv,
		inquiryApplication:       ia,
	}
}

// CreateInquiry - お問い合わせ登録
func (s *inquiryServer) CreateInquiry(
	ctx context.Context, req *pb.CreateInquiryRequest,
) (*pb.InquiryResponse, error) {
	err := s.inquiryRequestValidation.CreateInquiry(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	i := &inquiry.Inquiry{
		SenderID:    req.UserId,
		Subject:     req.Subject,
		Description: req.Description,
		Email:       req.Email,
		IsReplied:   false,
	}
	err = s.inquiryApplication.Create(ctx, i)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getInquiryResponse(i)
	return res, nil
}

func getInquiryResponse(i *inquiry.Inquiry) *pb.InquiryResponse {
	return &pb.InquiryResponse{
		Inquiry: i.Proto(),
	}
}
