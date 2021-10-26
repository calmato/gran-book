package server

import (
	"context"
	"io"
	"strings"

	"github.com/calmato/gran-book/api/internal/user/application"
	"github.com/calmato/gran-book/api/internal/user/domain/user"
	"github.com/calmato/gran-book/api/internal/user/interface/validation"
	"github.com/calmato/gran-book/api/pkg/exception"
	pb "github.com/calmato/gran-book/api/proto/user"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
	authRequestValidation validation.AuthRequestValidation
	userApplication       application.UserApplication
}

func NewAuthServer(arv validation.AuthRequestValidation, ua application.UserApplication) pb.AuthServiceServer {
	return &authServer{
		authRequestValidation: arv,
		userApplication:       ua,
	}
}

// GetAuth - 認証情報取得
func (s *authServer) GetAuth(ctx context.Context, req *pb.Empty) (*pb.AuthResponse, error) {
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// CreateAuth - ユーザー登録
func (s *authServer) CreateAuth(ctx context.Context, req *pb.CreateAuthRequest) (*pb.AuthResponse, error) {
	err := s.authRequestValidation.CreateAuth(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u := &user.User{
		Username: req.GetUsername(),
		Email:    strings.ToLower(req.GetEmail()),
		Password: req.GetPassword(),
	}

	err = s.userApplication.Create(ctx, u)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthEmail - メールアドレス更新
func (s *authServer) UpdateAuthEmail(ctx context.Context, req *pb.UpdateAuthEmailRequest) (*pb.AuthResponse, error) {
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.authRequestValidation.UpdateAuthEmail(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u.Email = strings.ToLower(req.GetEmail())

	err = s.userApplication.Update(ctx, u)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthPassword - パスワード更新
func (s *authServer) UpdateAuthPassword(
	ctx context.Context, req *pb.UpdateAuthPasswordRequest,
) (*pb.AuthResponse, error) {
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.authRequestValidation.UpdateAuthPassword(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u.Password = req.GetPassword()

	err = s.userApplication.UpdatePassword(ctx, u)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthProfile - プロフィール更新
func (s *authServer) UpdateAuthProfile(
	ctx context.Context, req *pb.UpdateAuthProfileRequest,
) (*pb.AuthResponse, error) {
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.authRequestValidation.UpdateAuthProfile(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u.Username = req.GetUsername()
	u.Gender = int(req.GetGender())
	u.ThumbnailURL = req.GetThumbnailUrl()
	u.SelfIntroduction = req.GetSelfIntroduction()

	err = s.userApplication.Update(ctx, u)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthAddress - 住所更新
func (s *authServer) UpdateAuthAddress(
	ctx context.Context, req *pb.UpdateAuthAddressRequest,
) (*pb.AuthResponse, error) {
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.authRequestValidation.UpdateAuthAddress(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u.LastName = req.GetLastName()
	u.FirstName = req.GetFirstName()
	u.LastNameKana = req.GetLastNameKana()
	u.FirstNameKana = req.GetFirstNameKana()
	u.PhoneNumber = req.GetPhoneNumber()
	u.PostalCode = req.GetPostalCode()
	u.Prefecture = req.GetPrefecture()
	u.City = req.GetCity()
	u.AddressLine1 = req.GetAddressLine1()
	u.AddressLine2 = req.GetAddressLine2()

	err = s.userApplication.Update(ctx, u)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UploadAuthThumbnail - サムネイルアップロード
func (s *authServer) UploadAuthThumbnail(stream pb.AuthService_UploadAuthThumbnailServer) error {
	ctx := stream.Context()
	thumbnailBytes := map[int][]byte{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			u, err := s.userApplication.Authentication(ctx)
			if err != nil {
				return toGRPCError(err)
			}

			// 分割して送信されてきたサムネイルのバイナリをまとめる
			thumbnail := []byte{}
			for i := 0; i < len(thumbnailBytes); i++ {
				thumbnail = append(thumbnail, thumbnailBytes[i]...)
			}

			thumbnailURL, err := s.userApplication.UploadThumbnail(ctx, u.ID, thumbnail)
			if err != nil {
				return toGRPCError(err)
			}

			res := getAuthThumbnailResponse(thumbnailURL)
			return stream.SendAndClose(res)
		}

		if err != nil {
			return toGRPCError(err)
		}

		err = s.authRequestValidation.UploadAuthThumbnail(req)
		if err != nil {
			return toGRPCError(err)
		}

		num := int(req.GetPosition())
		if thumbnailBytes[num] != nil {
			return toGRPCError(exception.ErrInvalidRequestValidation.New(errInvalidUploadRequest))
		}

		thumbnailBytes[num] = req.GetThumbnail()
	}
}

// DeleteAuth - ユーザー退会
func (s *authServer) DeleteAuth(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.userApplication.Delete(ctx, u)
	if err != nil {
		return nil, toGRPCError(err)
	}

	return &pb.Empty{}, nil
}

// RegisterAuthDevice - デバイス情報登録
func (s *authServer) RegisterAuthDevice(
	ctx context.Context, req *pb.RegisterAuthDeviceRequest,
) (*pb.AuthResponse, error) {
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.authRequestValidation.RegisterAuthDevice(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u.InstanceID = req.GetInstanceId()

	err = s.userApplication.Update(ctx, u)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

func getAuthResponse(u *user.User) *pb.AuthResponse {
	return &pb.AuthResponse{
		Auth: u.Auth(),
	}
}

func getAuthThumbnailResponse(thumbnailURL string) *pb.AuthThumbnailResponse {
	return &pb.AuthThumbnailResponse{
		ThumbnailUrl: thumbnailURL,
	}
}
