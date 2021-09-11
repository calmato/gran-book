package server

import (
	"context"
	"io"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/domain/chat"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase/firestore"
	pb "github.com/calmato/gran-book/api/server/user/proto/chat"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	chatRequestValidation validation.ChatRequestValidation
	chatApplication       application.ChatApplication
}

func NewChatServer(crv validation.ChatRequestValidation, ca application.ChatApplication) pb.ChatServiceServer {
	return &chatServer{
		chatRequestValidation: crv,
		chatApplication:       ca,
	}
}

// ListRoom - チャットルーム一覧取得
func (s *chatServer) ListRoom(ctx context.Context, req *pb.ListRoomRequest) (*pb.RoomListResponse, error) {
	err := s.chatRequestValidation.ListRoom(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	p := &firestore.Params{
		Limit:      int(req.GetLimit()),
		Offset:     req.GetOffset(),
		OrderField: "updatedAt",
		OrderBy:    firestore.OrderByDesc,
	}

	crs, err := s.chatApplication.ListRoom(ctx, req.GetUserId(), p)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatRoomListResponse(crs)
	return res, nil
}

// CreateRoom - チャットルーム作成
func (s *chatServer) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.RoomResponse, error) {
	err := s.chatRequestValidation.CreateRoom(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	cr := &chat.Room{
		UserIDs: req.GetUserIds(),
	}

	err = s.chatApplication.CreateRoom(ctx, cr)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatRoomResponse(cr)
	return res, nil
}

// CreateMessage - チャットメッセージ(テキスト)作成
func (s *chatServer) CreateMessage(
	ctx context.Context, req *pb.CreateMessageRequest,
) (*pb.MessageResponse, error) {
	err := s.chatRequestValidation.CreateMessage(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	cr, err := s.chatApplication.GetRoom(ctx, req.GetRoomId(), req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	cm := &chat.Message{
		UserID: req.GetUserId(),
		Text:   req.GetText(),
	}

	err = s.chatApplication.CreateMessage(ctx, cr, cm)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getChatMessageResponse(cm)
	return res, nil
}

// UploadImage - チャットメッセージ(イメージ)作成
func (s *chatServer) UploadImage(stream pb.ChatService_UploadImageServer) error {
	ctx := stream.Context()
	imageBytes := map[int][]byte{}
	userID := ""
	roomID := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			cr, err := s.chatApplication.GetRoom(ctx, req.GetRoomId(), req.GetUserId())
			if err != nil {
				return errorHandling(err)
			}

			// 分割して送信されてきたサムネイルのバイナリをまとめる
			image := []byte{}
			for i := 0; i < len(imageBytes); i++ {
				image = append(image, imageBytes[i]...)
			}

			imageURL, err := s.chatApplication.UploadImage(ctx, cr, image)
			if err != nil {
				return errorHandling(err)
			}

			cm := &chat.Message{
				UserID: req.GetUserId(),
				Image:  imageURL,
			}

			err = s.chatApplication.CreateMessage(ctx, cr, cm)
			if err != nil {
				return errorHandling(err)
			}

			res := getChatMessageResponse(cm)
			return stream.SendAndClose(res)
		}

		if err != nil {
			return errorHandling(err)
		}

		err = s.chatRequestValidation.UploadChatImage(req)
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
			return errorHandling(exception.InvalidRequestValidation.New(errInvalidUploadRequest))
		}

		imageBytes[num] = req.GetImage()
	}
}

func getChatRoomListResponse(crs chat.Rooms) *pb.RoomListResponse {
	return &pb.RoomListResponse{
		Rooms: crs.Proto(),
	}
}

func getChatRoomResponse(cr *chat.Room) *pb.RoomResponse {
	return &pb.RoomResponse{
		Room: cr.Proto(),
	}
}

func getChatMessageResponse(cm *chat.Message) *pb.MessageResponse {
	return &pb.MessageResponse{
		Message: cm.Proto(),
	}
}
