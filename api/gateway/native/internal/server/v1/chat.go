package v1

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/pkg/array"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ChatHandler interface {
	ListRoom(ctx *gin.Context)
	CreateRoom(ctx *gin.Context)
	CreateTextMessage(ctx *gin.Context)
	CreateImageMessage(ctx *gin.Context)
}

type chatHandler struct {
	chatClient pb.ChatServiceClient
	authClient pb.AuthServiceClient
	userClient pb.UserServiceClient
}

func NewChatHandler(chatConn *grpc.ClientConn, authConn *grpc.ClientConn, userConn *grpc.ClientConn) ChatHandler {
	cc := pb.NewChatServiceClient(chatConn)
	ac := pb.NewAuthServiceClient(authConn)
	uc := pb.NewUserServiceClient(userConn)

	return &chatHandler{
		chatClient: cc,
		authClient: ac,
		userClient: uc,
	}
}

// ListRoom - チャットルーム一覧取得
func (h *chatHandler) ListRoom(ctx *gin.Context) {
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.DefaultQuery("offset", "")

	c := util.SetMetadata(ctx)
	_, err := h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	roomsInput := &pb.ListChatRoomRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	roomsOutput, err := h.chatClient.ListRoom(c, roomsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	log.Printf("debug: %v\n", roomsOutput)

	userIDs := []string{}
	users := map[string]bool{}
	for _, r := range roomsOutput.GetRooms() {
		for _, userID := range r.GetUserIds() {
			if !users[userID] {
				users[userID] = true
				userIDs = append(userIDs, userID)
			}
		}
	}

	usersInput := &pb.MultiGetUserRequest{
		UserIds: userIDs,
	}

	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getChatRoomListResponse(roomsOutput, usersOutput)
	ctx.JSON(http.StatusOK, res)
}

// CreateRoom - チャットルーム作成
func (h *chatHandler) CreateRoom(ctx *gin.Context) {
	userID := ctx.Param("userID")
	req := &pb.CreateChatRoomV1Request{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	_, err = h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userIDs := req.UserIds
	if isExists, _ := array.Contains(userIDs, userID); !isExists {
		userIDs = append(userIDs, userID)
	}

	usersInput := &pb.MultiGetUserRequest{
		UserIds: userIDs,
	}

	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	for _, userID := range userIDs {
		if usersOutput.GetUsers()[userID] != nil {
			continue
		}

		err := fmt.Errorf("one of the user ids does not exist")
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	roomInput := &pb.CreateChatRoomRequest{
		UserIds: userIDs,
	}

	roomOutput, err := h.chatClient.CreateRoom(c, roomInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getChatRoomResponse(roomOutput, usersOutput)
	ctx.JSON(http.StatusOK, res)
}

// CreateTextMessage - チャットメッセージ(テキスト)作成
func (h *chatHandler) CreateTextMessage(ctx *gin.Context) {
	roomID := ctx.Param("roomID")
	userID := ctx.Param("userID")
	req := &pb.CreateChatMessageV1Request{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	userOutput, err := h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &pb.CreateChatMessageRequest{
		RoomId: roomID,
		UserId: userID,
		Text:   req.Text,
	}

	messageOutput, err := h.chatClient.CreateMessage(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getChatMessageResponse(messageOutput, userOutput)
	ctx.JSON(http.StatusOK, res)
}

// CreateImageMessage - チャットメッセージ(画像)作成
func (h *chatHandler) CreateImageMessage(ctx *gin.Context) {
	roomID := ctx.Param("roomID")
	userID := ctx.Param("userID")

	c := util.SetMetadata(ctx)
	userOutput, err := h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}
	defer file.Close()

	stream, err := h.chatClient.UploadImage(c)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrInternalServerError.New(err))
		return
	}

	var count int64           // 読み込み回数
	buf := make([]byte, 1024) // 1リクエストの上限設定

	in := &pb.UploadChatImageRequest{
		RoomId:   roomID,
		UserId:   userID,
		Position: count,
	}

	err = stream.Send(in)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrInternalServerError.New(err))
		return
	}

	count++

	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
			return
		}

		in = &pb.UploadChatImageRequest{
			Image:    buf,
			Position: count,
		}

		err = stream.Send(in)
		if err != nil {
			util.ErrorHandling(ctx, entity.ErrInternalServerError.New(err))
			return
		}

		count++
	}

	messageOutput, err := stream.CloseAndRecv()
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getChatMessageResponse(messageOutput, userOutput)
	ctx.JSON(http.StatusOK, res)
}

func (h *chatHandler) currentUser(ctx context.Context, userID string) (*pb.AuthResponse, error) {
	out, err := h.authClient.GetAuth(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	if out.GetId() != userID {
		return nil, entity.ErrForbidden.New(err)
	}

	return out, nil
}

func (h *chatHandler) getChatRoomResponse(
	roomOutput *pb.ChatRoomResponse, usersOutput *pb.UserMapResponse,
) *pb.ChatRoomV1Response {
	users := make([]*pb.ChatRoomV1Response_User, len(roomOutput.GetUserIds()))
	for i, userID := range roomOutput.GetUserIds() {
		user := &pb.ChatRoomV1Response_User{
			Id:       userID,
			Username: "unknown",
		}

		if usersOutput.GetUsers()[userID] != nil {
			user.Username = usersOutput.GetUsers()[userID].GetUsername()
			user.ThumbnailUrl = usersOutput.GetUsers()[userID].GetThumbnailUrl()
		}

		users[i] = user
	}

	return &pb.ChatRoomV1Response{
		Id:        roomOutput.GetId(),
		Users:     users,
		CreatedAt: roomOutput.GetCreatedAt(),
		UpdatedAt: roomOutput.GetUpdatedAt(),
	}
}

type ChatRoomListV1Response struct {
	Rooms []*pb.ChatRoomListV1Response_Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"roomsList,omitempty"` // チャットルーム一覧
}

func (h *chatHandler) getChatRoomListResponse(
	roomsOutput *pb.ChatRoomListResponse, usersOutput *pb.UserMapResponse,
) *ChatRoomListV1Response {
	rooms := make([]*pb.ChatRoomListV1Response_Room, len(roomsOutput.GetRooms()))
	for i, r := range roomsOutput.GetRooms() {
		users := make([]*pb.ChatRoomListV1Response_User, len(r.GetUserIds()))
		for j, userID := range r.GetUserIds() {
			user := &pb.ChatRoomListV1Response_User{
				Id:       userID,
				Username: "unknown",
			}

			if usersOutput.GetUsers()[userID] != nil {
				user.Username = usersOutput.GetUsers()[userID].GetUsername()
				user.ThumbnailUrl = usersOutput.GetUsers()[userID].GetThumbnailUrl()
			}

			users[j] = user
		}

		message := &pb.ChatRoomListV1Response_Message{}
		if r.GetLatestMessage() != nil {
			message.UserId = r.GetLatestMessage().GetUserId()
			message.Text = r.GetLatestMessage().GetText()
			message.Image = r.GetLatestMessage().GetImage()
			message.CreatedAt = r.GetLatestMessage().GetCreatedAt()
		}

		room := &pb.ChatRoomListV1Response_Room{
			Id:            r.GetId(),
			CreatedAt:     r.GetCreatedAt(),
			UpdatedAt:     r.GetUpdatedAt(),
			Users:         users,
			LatestMessage: message,
		}

		rooms[i] = room
	}

	return &ChatRoomListV1Response{
		Rooms: rooms,
	}
}

func (h *chatHandler) getChatMessageResponse(
	messageOutput *pb.ChatMessageResponse, userOutput *pb.AuthResponse,
) *pb.ChatMessageV1Response {
	user := &pb.ChatMessageV1Response_User{
		Id:           userOutput.GetId(),
		Username:     userOutput.GetUsername(),
		ThumbnailUrl: userOutput.GetThumbnailUrl(),
	}

	return &pb.ChatMessageV1Response{
		Text:      messageOutput.GetText(),
		Image:     messageOutput.GetImage(),
		CreatedAt: messageOutput.GetCreatedAt(),
		User:      user,
	}
}
