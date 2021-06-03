package v1

import (
	"github.com/calmato/gran-book/api/server/information/internal/application"
	pb "github.com/calmato/gran-book/api/server/information/proto"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
	AuthApplication application.AuthApplication
}
