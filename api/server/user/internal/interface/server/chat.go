package server

import (
	"context"
	"io"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"golang.org/x/xerrors"
)

// ChatServer - Chatサービスインターフェースの構造体
type ChatServer struct {
	pb.UnimplementedChatServiceServer
	ChatRequestValidation validation.ChatRequestValidation
	ChatApplication       application.ChatApplication
}

// ListRoom - チャットルーム一覧取得
func (s *ChatServer) ListRoom(ctx context.Context, req *pb.ListChatRoomRequest) (*pb.ChatRoomListResponse, error) {
	err := s.ChatRequestValidation.ListChatRoom(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	limit := int(req.GetLimit())
	offset := req.GetOffset()
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
	}

	crs, err := s.ChatApplication.ListRoom(ctx, req.GetUserId(), q)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatRoomListResponse(crs)
	return res, nil
}

// CreateRoom - チャットルーム作成
func (s *ChatServer) CreateRoom(ctx context.Context, req *pb.CreateChatRoomRequest) (*pb.ChatRoomResponse, error) {
	err := s.ChatRequestValidation.CreateChatRoom(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	cr := &chat.Room{
		UserIDs: req.GetUserIds(),
	}

	err = s.ChatApplication.CreateRoom(ctx, cr)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatRoomResponse(cr)
	return res, nil
}

// CreateMessage - チャットメッセージ(テキスト)作成
func (s *ChatServer) CreateMessage(ctx context.Context, req *pb.CreateChatMessageRequest) (*pb.ChatMessageResponse, error) {
	err := s.ChatRequestValidation.CreateChatMessage(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	cr, err := s.ChatApplication.GetRoom(ctx, req.GetRoomId(), req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	cm := &chat.Message{
		UserID: req.GetUserId(),
		Text:   req.GetText(),
	}

	err = s.ChatApplication.CreateMessage(ctx, cr, cm)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatMessageResponse(cm)
	return res, nil
}

// UploadImage - チャットメッセージ(イメージ)作成
func (s *ChatServer) UploadImage(stream pb.ChatService_UploadImageServer) error {
	ctx := stream.Context()
	imageBytes := map[int][]byte{}
	userID := ""
	roomID := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			cr, err := s.ChatApplication.GetRoom(ctx, req.GetRoomId(), req.GetUserId())
			if err != nil {
				return errorHandling(err)
			}

			// 分割して送信されてきたサムネイルのバイナリをまとめる
			image := []byte{}
			for i := 0; i < len(imageBytes); i++ {
				image = append(image, imageBytes[i]...)
			}

			imageURL, err := s.ChatApplication.UploadImage(ctx, cr, image)
			if err != nil {
				return errorHandling(err)
			}

			cm := &chat.Message{
				UserID: req.GetUserId(),
				Image:  imageURL,
			}

			err = s.ChatApplication.CreateMessage(ctx, cr, cm)
			if err != nil {
				return errorHandling(err)
			}

			res := getChatMessageResponse(cm)
			return stream.SendAndClose(res)
		}

		if err != nil {
			return errorHandling(err)
		}

		err = s.ChatRequestValidation.UploadChatImage(req)
		if err != nil {
			return errorHandling(err)
		}

		if userID == "" {
			userID = req.GetUserId()
		}

		if roomID == "" {
			roomID = req.GetRoomId()
		}

		num := int(req.GetPosition())
		if imageBytes[num] != nil {
			err = xerrors.New("Position is duplicated")
			return errorHandling(exception.InvalidRequestValidation.New(err))
		}

		imageBytes[num] = req.GetImage()
	}
}

func getChatRoomListResponse(crs []*chat.Room) *pb.ChatRoomListResponse {
	rooms := make([]*pb.ChatRoomListResponse_Room, len(crs))
	for i, cr := range crs {
		room := &pb.ChatRoomListResponse_Room{
			Id:        cr.ID,
			UserIds:   cr.UserIDs,
			CreatedAt: datetime.TimeToString(cr.CreatedAt),
			UpdatedAt: datetime.TimeToString(cr.UpdatedAt),
		}

		if cr.LatestMessage != nil {
			message := &pb.ChatRoomListResponse_Message{
				UserId:    cr.LatestMessage.UserID,
				Text:      cr.LatestMessage.Text,
				Image:     cr.LatestMessage.Image,
				CreatedAt: datetime.TimeToString(cr.LatestMessage.CreatedAt),
			}

			room.LatestMessage = message
		}

		rooms[i] = room
	}

	return &pb.ChatRoomListResponse{
		Rooms: rooms,
	}
}

func getChatRoomResponse(cr *chat.Room) *pb.ChatRoomResponse {
	return &pb.ChatRoomResponse{
		Id:        cr.ID,
		UserIds:   cr.UserIDs,
		CreatedAt: datetime.TimeToString(cr.CreatedAt),
		UpdatedAt: datetime.TimeToString(cr.UpdatedAt),
	}
}

func getChatMessageResponse(cm *chat.Message) *pb.ChatMessageResponse {
	return &pb.ChatMessageResponse{
		Id:        cm.ID,
		UserId:    cm.UserID,
		Text:      cm.Text,
		Image:     cm.Image,
		CreatedAt: datetime.TimeToString(cm.CreatedAt),
	}
}
