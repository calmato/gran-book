package v1

import (
	"io"
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
	req := &request.UpdateAuthProfileRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &pb.UpdateAuthProfileRequest{
		Username:         req.Username,
		Gender:           pb.Gender(req.Gender),
		ThumbnailUrl:     req.ThumbnailURL,
		SelfIntroduction: req.SelfIntroduction,
	}

	out, err := h.authClient.UpdateAuthProfile(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAuthResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// UpdateAddress - 住所情報更新
func (h *authHandler) UpdateAddress(ctx *gin.Context) {
	req := &request.UpdateAuthAddressRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &pb.UpdateAuthAddressRequest{
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

	out, err := h.authClient.UpdateAuthAddress(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAuthResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// UpdateEmail - メールアドレス更新
func (h *authHandler) UpdateEmail(ctx *gin.Context) {
	req := &request.UpdateAuthEmailRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &pb.UpdateAuthEmailRequest{
		Email: req.Email,
	}

	out, err := h.authClient.UpdateAuthEmail(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAuthResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// UpdatePassword - パスワード更新
func (h *authHandler) UpdatePassword(ctx *gin.Context) {
	req := &request.UpdateAuthPasswordRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
	}

	in := &pb.UpdateAuthPasswordRequest{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	out, err := h.authClient.UpdateAuthPassword(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getAuthResponse(out)
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

	stream, err := h.authClient.UploadAuthThumbnail(ctx)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrInternalServerError.New(err))
		return
	}

	var count int64           // 読み込み回数
	buf := make([]byte, 1024) // 1リクエストの上限設定
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
			return
		}

		in := &pb.UploadAuthThumbnailRequest{
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

	res := h.getAuthThumbnailResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Delete - ユーザー退会
func (h *authHandler) Delete(ctx *gin.Context) {
	_, err := h.authClient.DeleteAuth(ctx, &pb.Empty{})
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

	in := &pb.RegisterAuthDeviceRequest{
		InstanceId: req.InstanceID,
	}

	out, err := h.authClient.RegisterAuthDevice(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, nil)
		return
	}

	res := h.getAuthResponse(out)
	ctx.JSON(http.StatusOK, res)
}

func (h *authHandler) getAuthResponse(out *pb.AuthResponse) *response.AuthResponse {
	return &response.AuthResponse{
		ID:               out.GetId(),
		Username:         out.GetUsername(),
		Gender:           entity.Gender(out.GetGender()),
		Email:            out.GetEmail(),
		PhoneNumber:      out.GetPhoneNumber(),
		Role:             entity.Role(out.GetRole()),
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

func (h *authHandler) getAuthThumbnailResponse(out *pb.AuthThumbnailResponse) *response.AuthThumbnailResponse {
	return &response.AuthThumbnailResponse{
		ThumbnailURL: out.GetThumbnailUrl(),
	}
}
