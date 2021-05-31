package v1

import (
	"context"
	"fmt"

	"github.com/calmato/gran-book/api/server/information/internal/application"
	"github.com/calmato/gran-book/api/server/information/internal/domain/exception"
	pb "github.com/calmato/gran-book/api/server/information/proto"
	"golang.org/x/xerrors"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
	AuthApplication application.AuthApplication
}

func (s *NotificationServer) Reply(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	name := req.GetName()
	if name == "" {
		err := exception.InvalidRequestValidation.New(xerrors.New("Name is nil"))
		return nil, errorHandling(err)
	}

	message := fmt.Sprintf("Hello, %s!!", name)

	res := &pb.HelloResponse{Message: message}
	return res, nil
}
