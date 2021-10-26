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
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/user"
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

// updateAuthProfile - プロフィール情報更新
func (h *apiV1Handler) updateAuthProfile(ctx *gin.Context) {
	req := &request.UpdateAuthProfileRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	// TODO: 実装...

	ctx.JSON(http.StatusNotImplemented, nil)
}

// updateAuthEmail - メールアドレス更新
func (h *apiV1Handler) updateAuthEmail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAuthEmailRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
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

// updateAuthPassword - パスワード更新
func (h *apiV1Handler) updateAuthPassword(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateAuthPasswordRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
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

	res := &response.AuthThumbnailResponse{
		ThumbnailURL: out.ThumbnailUrl,
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
