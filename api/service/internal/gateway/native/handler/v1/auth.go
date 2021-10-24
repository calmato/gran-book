package v1

import (
	"context"
	"io"
	"net/http"

	gentity "github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/internal/gateway/native/entity"
	request "github.com/calmato/gran-book/api/service/internal/gateway/native/request/v1"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/gin-gonic/gin"
)

// getAuth - 認証情報取得
func (h *apiV1Handler) getAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	a, err := h.authGetAuth(c)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.AuthResponse{
		Auth: entity.NewAuth(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// createAuth - ユーザー登録
func (h *apiV1Handler) createAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateAuthRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.CreateAuthRequest{
		Username:             req.Username,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
	out, err := h.Auth.CreateAuth(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAuth(out.Auth)

	res := &response.AuthResponse{
		Auth: entity.NewAuth(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// updateAuthProfile - プロフィール情報更新
func (h *apiV1Handler) updateAuthProfile(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAuthProfileRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.UpdateAuthProfileRequest{
		Username:         req.Username,
		Gender:           user.Gender(req.Gender),
		ThumbnailUrl:     req.ThumbnailURL,
		SelfIntroduction: req.SelfIntroduction,
	}
	out, err := h.Auth.UpdateAuthProfile(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAuth(out.Auth)

	res := &response.AuthResponse{
		Auth: entity.NewAuth(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// updateAuthAddress - 住所情報更新
func (h *apiV1Handler) updateAuthAddress(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAuthAddressRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.UpdateAuthAddressRequest{
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		LastNameKana:  req.LastNameKana,
		FirstNameKana: req.FirstNameKana,
		PhoneNumber:   req.PhoneNumber,
		PostalCode:    req.PostalCode,
		Prefecture:    req.Prefecture,
		City:          req.City,
		AddressLine1:  req.AddressLine1,
		AddressLine2:  req.AddressLine2,
	}
	out, err := h.Auth.UpdateAuthAddress(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAuth(out.Auth)

	res := &response.AuthResponse{
		Auth: entity.NewAuth(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// updateAuthEmail - メールアドレス更新
func (h *apiV1Handler) updateAuthEmail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAuthEmailRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.UpdateAuthEmailRequest{
		Email: req.Email,
	}
	out, err := h.Auth.UpdateAuthEmail(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAuth(out.Auth)

	res := &response.AuthResponse{
		Auth: entity.NewAuth(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// updatePassword - パスワード更新
func (h *apiV1Handler) updateAuthPassword(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAuthPasswordRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.UpdateAuthPasswordRequest{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
	out, err := h.Auth.UpdateAuthPassword(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAuth(out.Auth)

	res := &response.AuthResponse{
		Auth: entity.NewAuth(a),
	}
	ctx.JSON(http.StatusOK, res)
}

// uploadAuthThumbnail サムネイルアップロード
func (h *apiV1Handler) uploadAuthThumbnail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	file, _, err := ctx.Request.FormFile("thumbnail")
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	defer file.Close()

	stream, err := h.Auth.UploadAuthThumbnail(c)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInternal.New(err))
		return
	}

	var count int64           // 読み込み回数
	buf := make([]byte, 1024) // 1リクエストの上限設定
	for {
		if _, err = file.Read(buf); err != nil {
			break
		}

		in := &user.UploadAuthThumbnailRequest{
			Thumbnail: buf,
			Position:  count,
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

	res := &response.AuthThumbnailResponse{
		ThumbnailURL: out.ThumbnailUrl,
	}
	ctx.JSON(http.StatusOK, res)
}

// deleteAuth - ユーザー退会
func (h *apiV1Handler) deleteAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	_, err := h.Auth.DeleteAuth(c, &user.Empty{})
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// registerAuthDevice - デバイス情報登録
func (h *apiV1Handler) registerAuthDevice(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.RegisterAuthDeviceRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &user.RegisterAuthDeviceRequest{
		InstanceId: req.InstanceID,
	}
	out, err := h.Auth.RegisterAuthDevice(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	a := gentity.NewAuth(out.Auth)

	res := &response.AuthResponse{
		Auth: entity.NewAuth(a),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) authGetAuth(ctx context.Context) (*gentity.Auth, error) {
	in := &user.Empty{}
	out, err := h.Auth.GetAuth(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewAuth(out.Auth), nil
}
