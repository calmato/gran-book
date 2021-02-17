package v1

import (
	"context"

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
		Username:             req.Username,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
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
		Email: req.Email,
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
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
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
		Username:         req.Username,
		Gender:           req.Gender,
		Thumbnail:        req.Thumbnail,
		SelfIntroduction: req.SelfIntroduction,
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

	err = s.AuthApplication.UpdateAddress(ctx, in, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

func getAuthResponse(u *user.User) *pb.AuthResponse {
	return &pb.AuthResponse{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           u.Gender,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             u.Role,
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
