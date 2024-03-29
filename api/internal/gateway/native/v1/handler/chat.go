package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/native/v1/entity"
	request "github.com/calmato/gran-book/api/internal/gateway/native/v1/request"
	response "github.com/calmato/gran-book/api/internal/gateway/native/v1/response"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/array"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/chat"
	"github.com/gin-gonic/gin"
)

// listChatRoom - チャットルーム一覧取得
func (h *apiV1Handler) listChatRoom(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", entity.ListLimitDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", entity.ListOffsetDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	ok, err := h.correctUser(c, userID)
	if err != nil || !ok {
		err = fmt.Errorf("v1: user id is not correct: %w", err)
		util.ErrorHandling(ctx, exception.ErrForbidden.New(err))
		return
	}

	crs, err := h.chatListRoom(c, userID, limit, offset)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us, err := h.userMultiGetUser(c, crs.UserIDs())
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.ChatRoomListResponse{
		Rooms: entity.NewChatRooms(crs, us.Map()),
	}
	ctx.JSON(http.StatusOK, res)
}

// createChatRoom - チャットルーム作成
func (h *apiV1Handler) createChatRoom(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")

	req := &request.CreateChatRoomRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	ok, err := h.correctUser(c, userID)
	if err != nil || !ok {
		err = fmt.Errorf("v1: user id is not correct: %w", err)
		util.ErrorHandling(ctx, exception.ErrForbidden.New(err))
		return
	}

	userIDs := req.UserIDs
	if isExists, _ := array.Contains(userIDs, userID); !isExists {
		userIDs = append(userIDs, userID)
	}

	us, err := h.userMultiGetUser(c, userIDs)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	if ok := us.IsExists(userIDs...); !ok {
		err := fmt.Errorf("v1: one of the user ids does not exist")
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &chat.CreateRoomRequest{
		UserIds: userIDs,
	}
	out, err := h.Chat.CreateRoom(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	cr := gentity.NewChatRoom(out.Room)

	res := &response.ChatRoomResponse{
		ChatRoom: entity.NewChatRoom(cr, us.Map()),
	}
	ctx.JSON(http.StatusOK, res)
}

// createChatTextMessage - チャットメッセージ(テキスト)作成
func (h *apiV1Handler) createChatTextMessage(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	roomID := ctx.Param("roomID")
	userID := ctx.Param("userID")

	req := &request.CreateChatMessageRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	ok, err := h.correctUser(c, userID)
	if err != nil || !ok {
		err = fmt.Errorf("v1: user id is not correct: %w", err)
		util.ErrorHandling(ctx, exception.ErrForbidden.New(err))
		return
	}

	u, err := h.userGetUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &chat.CreateMessageRequest{
		RoomId: roomID,
		UserId: userID,
		Text:   req.Text,
	}
	out, err := h.Chat.CreateMessage(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	cm := gentity.NewChatMessage(out.Message)

	res := &response.ChatMessageResponse{
		ChatMessage: entity.NewChatMessage(cm, u),
	}
	ctx.JSON(http.StatusOK, res)
}

// createChatImageMessage - チャットメッセージ(画像)作成
func (h *apiV1Handler) createChatImageMessage(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	roomID := ctx.Param("roomID")
	userID := ctx.Param("userID")

	ok, err := h.correctUser(c, userID)
	if err != nil || !ok {
		err = fmt.Errorf("v1: user id is not correct: %w", err)
		util.ErrorHandling(ctx, exception.ErrForbidden.New(err))
		return
	}

	u, err := h.userGetUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	defer file.Close()

	stream, err := h.Chat.UploadImage(c)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInternal.New(err))
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
		util.ErrorHandling(ctx, exception.ErrInternal.New(err))
		return
	}

	count++

	for {
		if _, err = file.Read(buf); err != nil {
			break
		}

		in = &chat.UploadChatImageRequest{
			Image:    buf,
			Position: count,
		}

		if err = stream.Send(in); err != nil {
			util.ErrorHandling(ctx, exception.ErrInternal.New(err))
			return
		}

		count++
	}

	if err != nil && err != io.EOF {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	out, err := stream.CloseAndRecv()
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	cm := gentity.NewChatMessage(out.Message)

	res := &response.ChatMessageResponse{
		ChatMessage: entity.NewChatMessage(cm, u),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) chatListRoom(
	ctx context.Context, userID string, limit, offset int64,
) (gentity.ChatRooms, error) {
	in := &chat.ListRoomRequest{
		UserId: userID,
		Limit:  limit,
		Offset: fmt.Sprint(offset),
	}
	out, err := h.Chat.ListRoom(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewChatRooms(out.Rooms), nil
}
