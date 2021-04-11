package v1

import (
	"context"
	"io"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

// AuthServer - Authインターフェースの構造体
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	AuthApplication application.AuthApplication
}

// GetAuth - 認証情報取得
func (s *AuthServer) GetAuth(ctx context.Context, req *pb.EmptyUser) (*pb.AuthResponse, error) {
	u, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// CreateAuth - ユーザ登録
func (s *AuthServer) CreateAuth(ctx context.Context, req *pb.CreateAuthRequest) (*pb.AuthResponse, error) {
	in := &input.CreateAuth{
		Username:             req.GetUsername(),
		Email:                req.GetEmail(),
		Password:             req.GetPassword(),
		PasswordConfirmation: req.GetPasswordConfirmation(),
	}

	u, err := s.AuthApplication.Create(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthEmail - ログイン用メールアドレスの更新
func (s *AuthServer) UpdateAuthEmail(ctx context.Context, req *pb.UpdateAuthEmailRequest) (*pb.AuthResponse, error) {
	u, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAuthEmail{
		Email: req.GetEmail(),
	}

	err = s.AuthApplication.UpdateEmail(ctx, in, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthPassword - ログイン用パスワードの更新
func (s *AuthServer) UpdateAuthPassword(
	ctx context.Context, req *pb.UpdateAuthPasswordRequest,
) (*pb.AuthResponse, error) {
	u, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAuthPassword{
		Password:             req.GetPassword(),
		PasswordConfirmation: req.GetPasswordConfirmation(),
	}

	err = s.AuthApplication.UpdatePassword(ctx, in, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthProfile - プロフィール更新
func (s *AuthServer) UpdateAuthProfile(
	ctx context.Context, req *pb.UpdateAuthProfileRequest,
) (*pb.AuthResponse, error) {
	u, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAuthProfile{
		Username:         req.GetUsername(),
		Gender:           int(req.GetGender()),
		ThumbnailURL:     req.GetThumbnailUrl(),
		SelfIntroduction: req.GetSelfIntroduction(),
	}

	err = s.AuthApplication.UpdateProfile(ctx, in, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, err
}

// UpdateAuthAddress - 住所更新
func (s *AuthServer) UpdateAuthAddress(
	ctx context.Context, req *pb.UpdateAuthAddressRequest,
) (*pb.AuthResponse, error) {
	u, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAuthAddress{
		LastName:      req.GetLastName(),
		FirstName:     req.GetFirstName(),
		LastNameKana:  req.GetLastNameKana(),
		FirstNameKana: req.GetFirstNameKana(),
		PhoneNumber:   req.GetPhoneNumber(),
		PostalCode:    req.GetPostalCode(),
		Prefecture:    req.GetPrefecture(),
		City:          req.GetCity(),
		AddressLine1:  req.GetAddressLine1(),
		AddressLine2:  req.GetAddressLine2(),
	}

	err = s.AuthApplication.UpdateAddress(ctx, in, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UploadAuthThumbnail - サムネイルアップロード
func (s *AuthServer) UploadAuthThumbnail(stream pb.AuthService_UploadAuthThumbnailServer) error {
	ctx := stream.Context()
	thumbnailBytes := map[int][]byte{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			u, err := s.AuthApplication.Authentication(ctx)
			if err != nil {
				return errorHandling(err)
			}

			// 分割して送信されてきたサムネイルのバイナリをまとめる
			thumbnail := []byte{}
			for i := 0; i < len(thumbnailBytes); i++ {
				thumbnail = append(thumbnail, thumbnailBytes[i]...)
			}

			in := &input.UploadAuthThumbnail{
				Thumbnail: thumbnail,
			}

			thumbnailURL, err := s.AuthApplication.UploadThumbnail(ctx, in, u)
			if err != nil {
				return errorHandling(err)
			}

			res := &pb.AuthThumbnailResponse{
				ThumbnailUrl: thumbnailURL,
			}

			return stream.SendAndClose(res)
		}

		if err != nil {
			return errorHandling(err)
		}

		num := int(req.GetPosition())
		thumbnailBytes[num] = req.GetThumbnail()
	}
}

func getAuthResponse(u *user.User) *pb.AuthResponse {
	return &pb.AuthResponse{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           int32(u.Gender),
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             int32(u.Role),
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		PostalCode:       u.PostalCode,
		Prefecture:       u.Prefecture,
		City:             u.City,
		AddressLine1:     u.AddressLine1,
		AddressLine2:     u.AddressLine2,
		Activated:        u.Activated,
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}
}
