package v1

import (
	"io"
	"net/http"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/admin/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/admin/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/admin/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type AdminHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateContact(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	UploadThumbnail(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type adminHandler struct {
	adminClient pb.AdminServiceClient
}

func NewAdminHandler(adminConn *grpc.ClientConn) AdminHandler {
	ac := pb.NewAdminServiceClient(adminConn)

	return &adminHandler{
		adminClient: ac,
	}
}

// List - 管理者一覧取得
func (h *adminHandler) List(ctx *gin.Context) {
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	field := ctx.DefaultQuery("field", "")
	value := ctx.DefaultQuery("value", "")
	by := ctx.DefaultQuery("by", "")
	direction := ctx.DefaultQuery("direction", "")

	in := &pb.ListAdminRequest{
		Limit:  limit,
		Offset: offset,
	}

	if field != "" {
		search := &pb.Search{
			Field: field,
			Value: value,
		}

		in.Search = search
	}

	// TODO: CamelCase -> snake_case に変換する関数作成したい..
	if by != "" {
		order := &pb.Order{
			Field:   by,
			OrderBy: entity.OrderBy(0).Value(direction).Proto(),
		}

		in.Order = order
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.ListAdmin(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAdminListResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Get - 管理者情報取得
func (h *adminHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &pb.GetAdminRequest{
		UserId: userID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.GetAdmin(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAdminResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Create - 管理者登録
func (h *adminHandler) Create(ctx *gin.Context) {
	req := &request.CreateAdminRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.CreateAdminRequest{
		LastName:             req.LastName,
		LastNameKana:         req.LastNameKana,
		FirstName:            req.FirstName,
		FirstNameKana:        req.FirstNameKana,
		Username:             req.Username,
		Role:                 req.Role.Proto(),
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.CreateAdmin(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAdminResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// UpdateContact - 連絡先更新
func (h *adminHandler) UpdateContact(ctx *gin.Context) {
	req := &request.UpdateAdminContactRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.UpdateAdminContactRequest{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.UpdateAdminContact(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAdminResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Update Profile - プロフィール更新
func (h *adminHandler) UpdateProfile(ctx *gin.Context) {
	req := &request.UpdateAdminProfileRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.UpdateAdminProfileRequest{
		LastName:      req.LastName,
		LastNameKana:  req.LastNameKana,
		FirstName:     req.FirstName,
		FirstNameKana: req.FirstNameKana,
		Username:      req.Username,
		Role:          req.Role.Proto(),
		ThumbnailUrl:  req.ThumbnailURL,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.UpdateAdminProfile(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAdminResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// UpdatePassword - パスワード更新
func (h *adminHandler) UpdatePassword(ctx *gin.Context) {
	req := &request.UpdateAdminPasswordRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.UpdateAdminPasswordRequest{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.UpdateAdminPassword(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAdminResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// UploadThumbnail - サムネイルアップロード
func (h *adminHandler) UploadThumbnail(ctx *gin.Context) {
	userID := ctx.Param("userID")

	file, _, err := ctx.Request.FormFile("thumbnail")
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}
	defer file.Close()

	c := util.SetMetadata(ctx)
	stream, err := h.adminClient.UploadAdminThumbnail(c)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrInternalServerError.New(err))
		return
	}

	var count int64           // 読み込み回数
	buf := make([]byte, 1024) // 1リクエストの上限設定

	in := &pb.UploadAdminThumbnailRequest{
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

		in := &pb.UploadAdminThumbnailRequest{
			Thumbnail: buf,
			Position:  count,
		}

		err = stream.Send(in)
		if err != nil {
			util.ErrorHandling(ctx, entity.ErrInternalServerError.New(err))
			return
		}

		count++
	}

	out, err := stream.CloseAndRecv()
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAdminThumbnailResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Delete - 管理者権限削除
func (h *adminHandler) Delete(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &pb.DeleteAdminRequest{
		UserId: userID,
	}

	c := util.SetMetadata(ctx)
	_, err := h.adminClient.DeleteAdmin(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *adminHandler) getAdminResponse(out *pb.AdminResponse) *response.AdminResponse {
	return &response.AdminResponse{
		ID:               out.GetId(),
		Username:         out.GetUsername(),
		Email:            out.GetEmail(),
		PhoneNumber:      out.GetPhoneNumber(),
		Role:             entity.Role(out.GetRole()),
		ThumbnailURL:     out.GetThumbnailUrl(),
		SelfIntroduction: out.GetSelfIntroduction(),
		LastName:         out.GetLastName(),
		FirstName:        out.GetFirstName(),
		LastNameKana:     out.GetLastNameKana(),
		FirstNameKana:    out.GetFirstNameKana(),
		CreatedAt:        out.GetCreatedAt(),
		UpdatedAt:        out.GetUpdatedAt(),
	}
}

func (h *adminHandler) getAdminListResponse(out *pb.AdminListResponse) *response.AdminListResponse {
	users := make([]*response.AdminListUser, len(out.GetUsers()))
	for i, u := range out.GetUsers() {
		user := &response.AdminListUser{
			ID:               u.GetId(),
			Username:         u.GetUsername(),
			Email:            u.GetEmail(),
			PhoneNumber:      u.GetPhoneNumber(),
			ThumbnailURL:     u.GetThumbnailUrl(),
			SelfIntroduction: u.GetSelfIntroduction(),
			LastName:         u.GetLastName(),
			FirstName:        u.GetFirstName(),
			LastNameKana:     u.GetLastNameKana(),
			FirstNameKana:    u.GetFirstNameKana(),
			CreatedAt:        u.GetCreatedAt(),
			UpdatedAt:        u.GetUpdatedAt(),
		}

		users[i] = user
	}

	return &response.AdminListResponse{
		Users:  users,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}

func (h *adminHandler) getAdminThumbnailResponse(out *pb.AdminThumbnailResponse) *response.AdminThumbnailResponse {
	return &response.AdminThumbnailResponse{
		ThumbnailURL: out.GetThumbnailUrl(),
	}
}
