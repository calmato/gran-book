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

// ListRoom - チャットルーム一覧取得
func (s *ChatServer) ListRoom(ctx context.Context, req *pb.ListChatRoomRequest) (*pb.ChatRoomListResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	crs, err := s.ChatApplication.ListRoom(ctx, cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatRoomListResponse(crs)
	return res, nil
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

func getChatRoomListResponse(crs []*chat.Room) *pb.ChatRoomListResponse {
	rs := make([]*pb.ChatRoomListResponse_Room, len(crs))
	for i, cr := range crs {
		r := &pb.ChatRoomListResponse_Room{
			Id:      cr.ID,
			UserIds: cr.UserIDs,
		}

		if cr.LatestMessage != nil {
			m := &pb.ChatRoomListResponse_Message{
				UserId:    cr.LatestMessage.UserID,
				Text:      cr.LatestMessage.Text,
				Image:     cr.LatestMessage.Image,
				CreatedAt: datetime.TimeToString(cr.LatestMessage.CreatedAt),
			}

			r.LatestMessage = m
		}

		rs[i] = r
	}

	return &pb.ChatRoomListResponse{
		Rooms: rs,
	}
}
