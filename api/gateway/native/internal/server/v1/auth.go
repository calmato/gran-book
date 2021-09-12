package v1

import (
	"io"
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
	UpdateEmail(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	UploadThumbnail(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterDevice(ctx *gin.Context)
}

type authHandler struct {
	authClient user.AuthServiceClient
}

func NewAuthHandler(ac user.AuthServiceClient) AuthHandler {
	return &authHandler{
		authClient: ac,
	}
}

// Get - 認証情報取得
func (h *authHandler) Get(ctx *gin.Context) {
	in := &user.Empty{}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.GetAuth(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAuth(out.Auth)
	res := h.getAuthResponse(a)
	ctx.JSON(http.StatusOK, res)
}

// Create - ユーザー登録
func (h *authHandler) Create(ctx *gin.Context) {
	req := &request.CreateAuthRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &user.CreateAuthRequest{
		Username:             req.Username,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.CreateAuth(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAuth(out.Auth)
	res := h.getAuthResponse(a)
	ctx.JSON(http.StatusOK, res)
}

// UpdateProfile - プロフィール情報更新
func (h *authHandler) UpdateProfile(ctx *gin.Context) {
	req := &request.UpdateAuthProfileRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &user.UpdateAuthProfileRequest{
		Username:         req.Username,
		Gender:           user.Gender(req.Gender),
		ThumbnailUrl:     req.ThumbnailURL,
		SelfIntroduction: req.SelfIntroduction,
	}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.UpdateAuthProfile(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAuth(out.Auth)
	res := h.getAuthResponse(a)
	ctx.JSON(http.StatusOK, res)
}

// UpdateAddress - 住所情報更新
func (h *authHandler) UpdateAddress(ctx *gin.Context) {
	req := &request.UpdateAuthAddressRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
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

	c := util.SetMetadata(ctx)
	out, err := h.authClient.UpdateAuthAddress(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAuth(out.Auth)
	res := h.getAuthResponse(a)
	ctx.JSON(http.StatusOK, res)
}

// UpdateEmail - メールアドレス更新
func (h *authHandler) UpdateEmail(ctx *gin.Context) {
	req := &request.UpdateAuthEmailRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &user.UpdateAuthEmailRequest{
		Email: req.Email,
	}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.UpdateAuthEmail(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAuth(out.Auth)
	res := h.getAuthResponse(a)
	ctx.JSON(http.StatusOK, res)
}

// UpdatePassword - パスワード更新
func (h *authHandler) UpdatePassword(ctx *gin.Context) {
	req := &request.UpdateAuthPasswordRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &user.UpdateAuthPasswordRequest{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.UpdateAuthPassword(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	a := entity.NewAuth(out.Auth)
	res := h.getAuthResponse(a)
	ctx.JSON(http.StatusOK, res)
}

// UploadThumbnail サムネイルアップロード
func (h *authHandler) UploadThumbnail(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("thumbnail")
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}
	defer file.Close()

	c := util.SetMetadata(ctx)
	stream, err := h.authClient.UploadAuthThumbnail(c)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrInternalServerError.New(err))
		return
	}

	var count int64           // 読み込み回数
	buf := make([]byte, 1024) // 1リクエストの上限設定
	for {
		if _, err = file.Read(buf); err != nil {
			if err == io.EOF {
				break
			}

			util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
			return
		}

		in := &user.UploadAuthThumbnailRequest{
			Thumbnail: buf,
			Position:  count,
		}

		if err = stream.Send(in); err != nil {
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

	res := h.getAuthThumbnailResponse(out.ThumbnailUrl)
	ctx.JSON(http.StatusOK, res)
}

// Delete - ユーザー退会
func (h *authHandler) Delete(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	_, err := h.authClient.DeleteAuth(c, &user.Empty{})
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// RegisterDevice - デバイス情報登録
func (h *authHandler) RegisterDevice(ctx *gin.Context) {
	req := &request.RegisterAuthDeviceRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &user.RegisterAuthDeviceRequest{
		InstanceId: req.InstanceID,
	}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.RegisterAuthDevice(c, in)
	if err != nil {
		util.ErrorHandling(ctx, nil)
		return
	}

	a := entity.NewAuth(out.Auth)
	res := h.getAuthResponse(a)
	ctx.JSON(http.StatusOK, res)
}

func (h *authHandler) getAuthResponse(a *entity.Auth) *response.AuthResponse {
	return &response.AuthResponse{
		ID:               a.Id,
		Username:         a.Username,
		Gender:           a.Gender(),
		Email:            a.Email,
		PhoneNumber:      a.PhoneNumber,
		ThumbnailURL:     a.ThumbnailUrl,
		SelfIntroduction: a.SelfIntroduction,
		LastName:         a.LastName,
		FirstName:        a.FirstName,
		LastNameKana:     a.LastNameKana,
		FirstNameKana:    a.FirstNameKana,
		PostalCode:       a.PostalCode,
		Prefecture:       a.Prefecture,
		City:             a.City,
		AddressLine1:     a.AddressLine1,
		AddressLine2:     a.AddressLine2,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
	}
}

func (h *authHandler) getAuthThumbnailResponse(thumbnailURL string) *response.AuthThumbnailResponse {
	return &response.AuthThumbnailResponse{
		ThumbnailURL: thumbnailURL,
	}
}
