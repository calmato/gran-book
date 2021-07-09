package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/internal/application"
	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/calmato/gran-book/api/server/information/internal/domain/notification"
	"github.com/calmato/gran-book/api/server/information/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/information/proto"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
	AuthApplication         application.AuthApplication
	NotificationApplication application.NotificationApplication
}

// CreateNotification - お知らせ登録
func (s *NotificationServer) Create(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.NotificationResponse, error) {
	in := &input.CreateNotification{
		Title:       req.Title,
		Description: req.Description,
		Importance:  req.Importance,
	}

	n, err := s.NotificationApplication.Create(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getNotificationResponse(n)
	return res, nil
}

func getNotificationResponse(n *notification.Notification) *pb.NotificationResponse {
	return &pb.NotificationResponse{
		Id:          int64(n.ID),
		AuthorId:    n.AuthorID,
		EditorId:    n.EditorID,
		Title:       n.Title,
		Description: n.Description,
		Importance:  n.Importance,
		Category:    "",
		CreatedAt:   datetime.TimeToString(n.CreatedAt),
		UpdatedAt:   datetime.TimeToString(n.UpdatedAt),
	}
}
