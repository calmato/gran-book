package v1

import (
	"io"
	"net/http"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/admin/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/admin/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/admin/proto/service/user"
	"github.com/gin-gonic/gin"
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
	adminClient user.AdminServiceClient
}

func NewAdminHandler(ac user.AdminServiceClient) AdminHandler {
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

	in := &user.ListAdminRequest{
		Limit:  limit,
		Offset: offset,
	}

	if field != "" {
		search := &user.Search{
			Field: field,
			Value: value,
		}

		in.Search = search
	}

	// TODO: CamelCase -> snake_case に変換する関数作成したい..
	if by != "" {
		order := &user.Order{
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

	as := entity.NewAdmins(out.Admins)

	res := h.getAdminListResponse(as, out.Limit, out.Offset, out.Total)
	ctx.JSON(http.StatusOK, res)
}

// Get - 管理者情報取得
func (h *adminHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &user.GetAdminRequest{
		UserId: userID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.GetAdmin(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAdmin(out.Admin)

	res := h.getAdminResponse(a)
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

	in := &user.CreateAdminRequest{
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

	a := entity.NewAdmin(out.Admin)

	res := h.getAdminResponse(a)
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

	in := &user.UpdateAdminContactRequest{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.UpdateAdminContact(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAdmin(out.Admin)

	res := h.getAdminResponse(a)
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

	in := &user.UpdateAdminProfileRequest{
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

	a := entity.NewAdmin(out.Admin)

	res := h.getAdminResponse(a)
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

	in := &user.UpdateAdminPasswordRequest{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	c := util.SetMetadata(ctx)
	out, err := h.adminClient.UpdateAdminPassword(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAdmin(out.Admin)

	res := h.getAdminResponse(a)
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

	in := &user.UploadAdminThumbnailRequest{
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

		in := &user.UploadAdminThumbnailRequest{
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

	res := h.getAdminThumbnailResponse(out.ThumbnailUrl)
	ctx.JSON(http.StatusOK, res)
}

// Delete - 管理者権限削除
func (h *adminHandler) Delete(ctx *gin.Context) {
	userID := ctx.Param("userID")

	in := &user.DeleteAdminRequest{
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

func (h *adminHandler) getAdminResponse(a *entity.Admin) *response.AdminResponse {
	return &response.AdminResponse{
		ID:               a.Id,
		Username:         a.Username,
		Email:            a.Email,
		PhoneNumber:      a.PhoneNumber,
		Role:             a.Role(),
		ThumbnailURL:     a.ThumbnailUrl,
		SelfIntroduction: a.SelfIntroduction,
		LastName:         a.LastName,
		FirstName:        a.FirstName,
		LastNameKana:     a.LastNameKana,
		FirstNameKana:    a.FirstNameKana,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
	}
}

func (h *adminHandler) getAdminListResponse(as entity.Admins, limit, offset, total int64) *response.AdminListResponse {
	users := make([]*response.AdminListUser, len(as))
	for i, a := range as {
		user := &response.AdminListUser{
			ID:               a.Id,
			Username:         a.Username,
			Email:            a.Email,
			Role:             a.Role(),
			PhoneNumber:      a.PhoneNumber,
			ThumbnailURL:     a.ThumbnailUrl,
			SelfIntroduction: a.SelfIntroduction,
			LastName:         a.LastName,
			FirstName:        a.FirstName,
			LastNameKana:     a.LastNameKana,
			FirstNameKana:    a.FirstNameKana,
			CreatedAt:        a.CreatedAt,
			UpdatedAt:        a.UpdatedAt,
		}

		users[i] = user
	}

	return &response.AdminListResponse{
		Users:  users,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

func (h *adminHandler) getAdminThumbnailResponse(thumbnailURL string) *response.AdminThumbnailResponse {
	return &response.AdminThumbnailResponse{
		ThumbnailURL: thumbnailURL,
	}
}
