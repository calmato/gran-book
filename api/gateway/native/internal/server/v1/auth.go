package v1

import (
	"net/http"

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

type authHandler struct{}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

// Get - 認証情報取得
func (h *authHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// Create - ユーザー登録
func (h *authHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// UpdateProfile - プロフィール情報更新
func (h *authHandler) UpdateProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// UpdateAddress - 住所情報更新
func (h *authHandler) UpdateAddress(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// UpdateEmail - メールアドレス更新
func (h *authHandler) UpdateEmail(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// UpdatePassword - パスワード更新
func (h *authHandler) UpdatePassword(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// UploadThumbnail サムネイルアップロード
func (h *authHandler) UploadThumbnail(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// Delete - ユーザー退会
func (h *authHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}

// RegisterDevice - デバイス情報登録
func (h *authHandler) RegisterDevice(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}
