package v1

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	request "github.com/calmato/gran-book/api/service/internal/gateway/native/request/v1"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/array"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/proto/chat"
	"github.com/calmato/gran-book/api/service/proto/user"
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

	_, err = h.currentUser(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	roomsInput := &chat.ListRoomRequest{
		UserId: userID,
		Limit:  limit,
		Offset: fmt.Sprint(offset),
	}
	roomsOutput, err := h.Chat.ListRoom(c, roomsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	crs := entity.NewChatRooms(roomsOutput.Rooms)

	usersInput := &user.MultiGetUserRequest{
		UserIds: crs.UserIDs(),
	}
	usersOutput, err := h.User.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	res := response.NewChatRoomListResponse(crs, us.Map())
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
	usersOutput, err := h.User.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	if ok := us.IsExists(userIDs...); !ok {
		err := fmt.Errorf("one of the user ids does not exist")
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	roomInput := &chat.CreateRoomRequest{
		UserIds: userIDs,
	}
	roomOutput, err := h.Chat.CreateRoom(c, roomInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	cr := entity.NewChatRoom(roomOutput.Room)
	res := response.NewChatRoomResponse(cr, us.Map())
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
	out, err := h.Chat.CreateMessage(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	cm := entity.NewChatMessage(out.Message)
	res := response.NewChatMessageResponse(cm, a)
	ctx.JSON(http.StatusOK, res)
}

// createChatImageMessage - チャットメッセージ(画像)作成
func (h *apiV1Handler) createChatImageMessage(ctx *gin.Context) {
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
			if err == io.EOF {
				break
			}

			util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
			return
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

	messageOutput, err := stream.CloseAndRecv()
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	cm := entity.NewChatMessage(messageOutput.Message)
	res := response.NewChatMessageResponse(cm, a)
	ctx.JSON(http.StatusOK, res)
}
