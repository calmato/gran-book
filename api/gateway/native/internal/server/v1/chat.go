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
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.DefaultQuery("offset", "")

	c := util.SetMetadata(ctx)
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
	res := h.getChatRoomListResponse(crs, us.Map())
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

	c := util.SetMetadata(ctx)
	_, err = h.currentUser(c, userID)
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
	res := h.getChatRoomResponse(cr, us.Map())
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

	c := util.SetMetadata(ctx)
	userOutput, err := h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAuth(userOutput.Auth)

	in := &chat.CreateMessageRequest{
		RoomId: roomID,
		UserId: userID,
		Text:   req.Text,
	}

	messageOutput, err := h.chatClient.CreateMessage(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	cm := entity.NewChatMessage(messageOutput.Message)
	res := h.getChatMessageResponse(cm, a)
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

	a := entity.NewAuth(userOutput.Auth)
	cm := entity.NewChatMessage(messageOutput.Message)
	res := h.getChatMessageResponse(cm, a)
	ctx.JSON(http.StatusOK, res)
}

func (h *chatHandler) currentUser(ctx context.Context, userID string) (*user.AuthResponse, error) {
	out, err := h.authClient.GetAuth(ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}

	a := entity.NewAuth(out.Auth)

	if a.Id != userID {
		return nil, entity.ErrForbidden.New(err)
	}

	return out, nil
}

func (h *chatHandler) getChatRoomResponse(cr *entity.ChatRoom, us map[string]*entity.User) *response.ChatRoomResponse {
	users := make([]*response.ChatRoomUser, len(cr.UserIds))
	for i, userID := range cr.UserIds {
		user := &response.ChatRoomUser{
			ID:       userID,
			Username: "unknown",
		}

		if us[userID] != nil {
			user.Username = us[userID].Username
			user.ThumbnailURL = us[userID].ThumbnailUrl
		}

		users[i] = user
	}

	return &response.ChatRoomResponse{
		ID:        cr.Id,
		Users:     users,
		CreatedAt: cr.CreatedAt,
		UpdatedAt: cr.UpdatedAt,
	}
}

func (h *chatHandler) getChatRoomListResponse(
	crs entity.ChatRooms, us map[string]*entity.User,
) *response.ChatRoomListResponse {
	rooms := make([]*response.ChatRoomListRoom, len(crs))
	for i, r := range crs {
		users := make([]*response.ChatRoomListUser, len(r.UserIds))
		for j, userID := range r.UserIds {
			user := &response.ChatRoomListUser{
				ID:       userID,
				Username: "unknown",
			}

			if us[userID] != nil {
				user.Username = us[userID].Username
				user.ThumbnailURL = us[userID].ThumbnailUrl
			}

			users[j] = user
		}

		message := &response.ChatRoomListMessage{}
		if r.LatestMessage != nil {
			message.UserID = r.LatestMessage.UserId
			message.Text = r.LatestMessage.Text
			message.Image = r.LatestMessage.Image
			message.CreatedAt = r.LatestMessage.CreatedAt
		}

		room := &response.ChatRoomListRoom{
			ID:            r.Id,
			CreatedAt:     r.CreatedAt,
			UpdatedAt:     r.UpdatedAt,
			Users:         users,
			LatestMessage: message,
		}

		rooms[i] = room
	}

	return &response.ChatRoomListResponse{
		Rooms: rooms,
	}
}

func (h *chatHandler) getChatMessageResponse(cm *entity.ChatMessage, a *entity.Auth) *response.ChatMessageResponse {
	user := &response.ChatMessageUser{
		ID:           a.Id,
		Username:     a.Username,
		ThumbnailURL: a.ThumbnailUrl,
	}

	return &response.ChatMessageResponse{
		Text:      cm.Text,
		Image:     cm.Image,
		CreatedAt: cm.CreatedAt,
		User:      user,
	}
}
