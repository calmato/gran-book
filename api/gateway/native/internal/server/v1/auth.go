package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
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
	authClient pb.AuthServiceClient
}

func NewAuthHandler(authConn *grpc.ClientConn) AuthHandler {
	ac := pb.NewAuthServiceClient(authConn)

	return &authHandler{
		authClient: ac,
	}
}

// Get - 認証情報取得
func (h *authHandler) Get(ctx *gin.Context) {
	in := &pb.Empty{}

	out, err := h.authClient.GetAuth(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAuthResponse(out)
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

	in := &pb.CreateAuthRequest{
		Username:             req.Username,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	out, err := h.authClient.CreateAuth(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAuthResponse(out)
	ctx.JSON(http.StatusOK, res)
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

func (h *authHandler) getAuthResponse(out *pb.AuthResponse) *response.AuthResponse {
	return &response.AuthResponse{
		ID:               out.GetId(),
		Username:         out.GetUsername(),
		Gender:           int(out.GetGender()),
		Email:            out.GetEmail(),
		PhoneNumber:      out.GetPhoneNumber(),
		Role:             int(out.GetRole()),
		ThumbnailURL:     out.GetThumbnailUrl(),
		SelfIntroduction: out.GetSelfIntroduction(),
		LastName:         out.GetLastName(),
		FirstName:        out.GetFirstName(),
		LastNameKana:     out.GetLastNameKana(),
		FirstNameKana:    out.GetFirstNameKana(),
		PostalCode:       out.GetPostalCode(),
		Prefecture:       out.GetPrefecture(),
		City:             out.GetCity(),
		AddressLine1:     out.GetAddressLine1(),
		AddressLine2:     out.GetAddressLine2(),
		CreatedAt:        out.GetCreatedAt(),
		UpdatedAt:        out.GetUpdatedAt(),
	}
}
