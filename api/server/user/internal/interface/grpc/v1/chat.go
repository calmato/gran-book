package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

// ChatService - Chatインターフェースの構造体
type ChatServer struct {
	pb.UnimplementedChatServiceServer
	AuthApplication application.AuthApplication
	ChatApplication application.ChatApplication
}

// CreateRoom - チャットルーム作成
func (s *ChatServer) CreateRoom(ctx context.Context, req *pb.CreateChatRoomRequest) (*pb.ChatRoomResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.CreateRoom{
		UserIDs: req.GetUserIds(),
	}

	cr, err := s.ChatApplication.CreateRoom(ctx, in, cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatRoomResponse(cr)
	return res, nil
}

func getChatRoomResponse(cr *chat.Room) *pb.ChatRoomResponse {
	return &pb.ChatRoomResponse{
		Id:        cr.ID,
		UserIds:   cr.UserIDs,
		CreatedAt: datetime.TimeToString(cr.CreatedAt),
		UpdatedAt: datetime.TimeToString(cr.CreatedAt),
	}
}
