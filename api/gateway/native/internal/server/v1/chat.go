package v1

import (
	"fmt"
	"io"
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
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
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.LIST_LIMIT_DEFAULT))
	offset := ctx.DefaultQuery("offset", "")

	_, err := h.currentUser(ctx, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &pb.ListChatRoomRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	out, err := h.chatClient.ListRoom(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getChatRoomListResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// CreateRoom - チャットルーム作成
func (h *chatHandler) CreateRoom(ctx *gin.Context) {
	userID := ctx.Param("userID")
	req := &request.CreateChatRoomRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	_, err = h.currentUser(ctx, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userIDs := req.UserIDs
	if isExists, _ := array.Contains(userIDs, userID); !isExists {
		userIDs = append(userIDs, userID)
	}

	usersInput := &pb.MultiGetUserRequest{
		UserIds: userIDs,
	}

	usersOutput, err := h.userClient.MultiGetUser(ctx, usersInput)
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

	roomOutput, err := h.chatClient.CreateRoom(ctx, roomInput)
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
	req := &request.CreateChatMessageRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	userOutput, err := h.currentUser(ctx, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &pb.CreateChatMessageRequest{
		RoomId: roomID,
		UserId: userID,
		Text:   req.Text,
	}

	messageOutput, err := h.chatClient.CreateMessage(ctx, in)
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

	userOutput, err := h.currentUser(ctx, userID)
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

	stream, err := h.chatClient.UploadImage(ctx)
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

	stream.Send(in)
	count += 1

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

		stream.Send(in)
		count += 1
	}

	messageOutput, err := stream.CloseAndRecv()
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getChatMessageResponse(messageOutput, userOutput)
	ctx.JSON(http.StatusOK, res)
}

func (h *chatHandler) currentUser(ctx *gin.Context, userID string) (*pb.AuthResponse, error) {
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
) *response.ChatRoomResponse {
	users := make([]*response.ChatRoomUser, len(roomOutput.GetUserIds()))
	for i, userID := range roomOutput.GetUserIds() {
		user := &response.ChatRoomUser{
			ID:       userID,
			Username: "unknown",
		}

		if usersOutput.GetUsers()[userID] != nil {
			user.Username = usersOutput.GetUsers()[userID].GetUsername()
			user.ThumbnailURL = usersOutput.GetUsers()[userID].GetThumbnailUrl()
		}

		users[i] = user
	}

	return &response.ChatRoomResponse{
		ID:        roomOutput.GetId(),
		Users:     users,
		CreatedAt: roomOutput.GetCreatedAt(),
		UpdatedAt: roomOutput.GetUpdatedAt(),
	}
}

func (h *chatHandler) getChatRoomListResponse(out *pb.ChatRoomListResponse) *response.ChatRoomListResponse {
	return &response.ChatRoomListResponse{}
}

func (h *chatHandler) getChatMessageResponse(
	messageOutput *pb.ChatMessageResponse, userOutput *pb.AuthResponse,
) *response.ChatMessageResponse {
	user := &response.ChatMessageUser{
		ID:           userOutput.GetId(),
		Username:     userOutput.GetUsername(),
		ThumbnailURL: userOutput.GetThumbnailUrl(),
	}

	return &response.ChatMessageResponse{
		Text:      messageOutput.GetText(),
		Image:     messageOutput.GetImage(),
		CreatedAt: messageOutput.GetCreatedAt(),
		User:      user,
	}
}
