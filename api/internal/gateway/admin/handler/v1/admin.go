package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/calmato/gran-book/api/internal/gateway/admin/entity"
	request "github.com/calmato/gran-book/api/internal/gateway/admin/request/v1"
	response "github.com/calmato/gran-book/api/internal/gateway/admin/response/v1"
	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/conv"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/gin-gonic/gin"
)

// listAdmin - 管理者一覧取得
func (h *apiV1Handler) listAdmin(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	field := ctx.DefaultQuery("field", "")
	value := ctx.DefaultQuery("value", "")
	by := ctx.DefaultQuery("by", "")
	direction := ctx.DefaultQuery("direction", "")

	as, total, err := h.adminListAdmin(c, limit, offset, field, value, by, direction)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.AdminListResponse{
		Users:  entity.NewAdmins(as),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	ctx.JSON(http.StatusOK, res)
}

// getAdmin - 管理者情報取得
func (h *apiV1Handler) getAdmin(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")

	a, err := h.adminGetAdmin(c, userID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.AdminResponse{
		Admin: entity.NewAdmin(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// createAdmin - 管理者登録
func (h *apiV1Handler) createAdmin(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateAdminRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
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
	out, err := h.Admin.CreateAdmin(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAdmin(out.Admin)

	res := &response.AdminResponse{
		Admin: entity.NewAdmin(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// updateAdminContact - 連絡先更新
func (h *apiV1Handler) updateAdminContact(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAdminContactRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.UpdateAdminContactRequest{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	out, err := h.Admin.UpdateAdminContact(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAdmin(out.Admin)

	res := &response.AdminResponse{
		Admin: entity.NewAdmin(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// updateAdminProfile - プロフィール更新
func (h *apiV1Handler) updateAdminProfile(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAdminProfileRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
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
	out, err := h.Admin.UpdateAdminProfile(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAdmin(out.Admin)

	res := &response.AdminResponse{
		Admin: entity.NewAdmin(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// updateAdminPassword - パスワード更新
func (h *apiV1Handler) updateAdminPassword(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAdminPasswordRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.UpdateAdminPasswordRequest{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
	out, err := h.Admin.UpdateAdminPassword(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAdmin(out.Admin)

	res := &response.AdminResponse{
		Admin: entity.NewAdmin(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// uploadAdminThumbnail - サムネイルアップロード
func (h *apiV1Handler) uploadAdminThumbnail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")

	file, _, err := ctx.Request.FormFile("thumbnail")
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	defer file.Close()

	stream, err := h.Admin.UploadAdminThumbnail(c)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInternal.New(err))
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
		util.ErrorHandling(ctx, exception.ErrInternal.New(err))
		return
	}

	count++

	for {
		if _, err = file.Read(buf); err != nil {
			break
		}

		in := &user.UploadAdminThumbnailRequest{
			Thumbnail: buf,
			Position:  count,
		}
		err = stream.Send(in)
		if err != nil {
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

	res := &response.AdminThumbnailResponse{
		ThumbnailURL: out.ThumbnailUrl,
	}
	ctx.JSON(http.StatusOK, res)
}

// deleteAdmin - 管理者権限削除
func (h *apiV1Handler) deleteAdmin(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")

	in := &user.DeleteAdminRequest{
		UserId: userID,
	}
	_, err := h.Admin.DeleteAdmin(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *apiV1Handler) adminListAdmin(
	ctx context.Context,
	limit, offset int64,
	field, value, by, direction string,
) (gentity.Admins, int64, error) {
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

	if by != "" {
		orderBy, err := conv.CamelToSnake(by)
		if err != nil {
			err := exception.ErrInvalidArgument.New(err)
			return nil, 0, err
		}

		order := &user.Order{
			Field:   orderBy,
			OrderBy: gentity.NewOrderByByValue(direction).Proto(),
		}
		in.Order = order
	}

	out, err := h.Admin.ListAdmin(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewAdmins(out.Admins), out.Total, nil
}

func (h *apiV1Handler) adminGetAdmin(ctx context.Context, userID string) (*gentity.Admin, error) {
	in := &user.GetAdminRequest{
		UserId: userID,
	}

	out, err := h.Admin.GetAdmin(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewAdmin(out.Admin), nil
}
