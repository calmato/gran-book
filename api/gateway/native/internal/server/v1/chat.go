package v1

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/pkg/array"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

type ChatHandler interface {
	ListRoom(ctx *gin.Context)
	CreateRoom(ctx *gin.Context)
	CreateTextMessage(ctx *gin.Context)
	CreateImageMessage(ctx *gin.Context)
}

type chatHandler struct {
	authClient user.AuthServiceClient
	chatClient chat.ChatServiceClient
	userClient user.UserServiceClient
}

func NewChatHandler(ac user.AuthServiceClient, cc chat.ChatServiceClient, uc user.UserServiceClient) ChatHandler {
	return &chatHandler{
		authClient: ac,
		chatClient: cc,
		userClient: uc,
	}
}

// ListRoom - チャットルーム一覧取得
func (h *chatHandler) ListRoom(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.DefaultQuery("offset", "")

	_, err := h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	roomsInput := &chat.ListRoomRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	roomsOutput, err := h.chatClient.ListRoom(c, roomsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	crs := entity.NewChatRooms(roomsOutput.Rooms)

	usersInput := &user.MultiGetUserRequest{
		UserIds: crs.UserIDs(),
	}
	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	res := response.NewChatRoomListResponse(crs, us.Map())
	ctx.JSON(http.StatusOK, res)
}

// CreateRoom - チャットルーム作成
func (h *chatHandler) CreateRoom(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")

	req := &request.CreateChatRoomRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	_, err := h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userIDs := req.UserIDs
	if isExists, _ := array.Contains(userIDs, userID); !isExists {
		userIDs = append(userIDs, userID)
	}

	usersInput := &user.MultiGetUserRequest{
		UserIds: userIDs,
	}
	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	if ok := us.IsExists(userIDs...); !ok {
		err := fmt.Errorf("one of the user ids does not exist")
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	roomInput := &chat.CreateRoomRequest{
		UserIds: userIDs,
	}
	roomOutput, err := h.chatClient.CreateRoom(c, roomInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	cr := entity.NewChatRoom(roomOutput.Room)
	res := response.NewChatRoomResponse(cr, us.Map())
	ctx.JSON(http.StatusOK, res)
}

// CreateTextMessage - チャットメッセージ(テキスト)作成
func (h *chatHandler) CreateTextMessage(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	roomID := ctx.Param("roomID")
	userID := ctx.Param("userID")

	req := &request.CreateChatMessageRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	a, err := h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &chat.CreateMessageRequest{
		RoomId: roomID,
		UserId: userID,
		Text:   req.Text,
	}
	out, err := h.chatClient.CreateMessage(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	cm := entity.NewChatMessage(out.Message)
	res := response.NewChatMessageResponse(cm, a)
	ctx.JSON(http.StatusOK, res)
}

// CreateImageMessage - チャットメッセージ(画像)作成
func (h *chatHandler) CreateImageMessage(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	roomID := ctx.Param("roomID")
	userID := ctx.Param("userID")

	a, err := h.currentUser(c, userID)
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

	in := &chat.UploadChatImageRequest{
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
		if _, err = file.Read(buf); err != nil {
			if err == io.EOF {
				break
			}

			util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
			return
		}

		in = &chat.UploadChatImageRequest{
			Image:    buf,
			Position: count,
		}

		if err = stream.Send(in); err != nil {
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

	cm := entity.NewChatMessage(messageOutput.Message)
	res := response.NewChatMessageResponse(cm, a)
	ctx.JSON(http.StatusOK, res)
}

func (h *chatHandler) currentUser(ctx context.Context, userID string) (*entity.Auth, error) {
	out, err := h.authClient.GetAuth(ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}

	a := entity.NewAuth(out.Auth)

	if a.Id != userID {
		return nil, entity.ErrForbidden.New(err)
	}

	return a, nil
}
