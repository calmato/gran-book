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

type AuthHandler interface {
	Get(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	UpdateEmail(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	UploadThumbnail(ctx *gin.Context)
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

	c := util.SetMetadata(ctx)
	out, err := h.authClient.GetAuth(c, in)
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
		return
	}

	// TODO: 実装...

	ctx.JSON(http.StatusNotImplemented, nil)
}

// UpdateEmail - メールアドレス更新
func (h *authHandler) UpdateEmail(ctx *gin.Context) {
	req := &request.UpdateAuthEmailRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &pb.UpdateAuthEmailRequest{
		Email: req.Email,
	}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.UpdateAuthEmail(c, in)
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
		return
	}

	in := &pb.UpdateAuthPasswordRequest{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	c := util.SetMetadata(ctx)
	out, err := h.authClient.UpdateAuthPassword(c, in)
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

	c := util.SetMetadata(ctx)
	stream, err := h.authClient.UploadAuthThumbnail(c)
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

func (h *authHandler) getAuthResponse(out *pb.AuthResponse) *response.AuthResponse {
	return &response.AuthResponse{
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

func (h *authHandler) getAuthThumbnailResponse(out *pb.AuthThumbnailResponse) *response.AuthThumbnailResponse {
	return &response.AuthThumbnailResponse{
		ThumbnailURL: out.GetThumbnailUrl(),
	}
}
