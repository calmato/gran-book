package server

import (
	"context"
	"io"
	"strings"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"golang.org/x/xerrors"
)

// AuthServer - Authサービスインターフェースの構造体
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	AuthRequestValidation validation.AuthRequestValidation
	UserApplication       application.UserApplication
}

// GetAuth - 認証情報取得
func (s *AuthServer) GetAuth(ctx context.Context, req *pb.Empty) (*pb.AuthResponse, error) {
	u, err := s.UserApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// CreateAuth - ユーザー登録
func (s *AuthServer) CreateAuth(ctx context.Context, req *pb.CreateAuthRequest) (*pb.AuthResponse, error) {
	err := s.AuthRequestValidation.CreateAuth(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u := &user.User{
		Username: req.GetUsername(),
		Email:    strings.ToLower(req.GetEmail()),
		Password: req.GetPassword(),
	}

	err = s.UserApplication.Create(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthEmail - メールアドレス更新
func (s *AuthServer) UpdateAuthEmail(ctx context.Context, req *pb.UpdateAuthEmailRequest) (*pb.AuthResponse, error) {
	u, err := s.UserApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AuthRequestValidation.UpdateAuthEmail(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Email = strings.ToLower(req.GetEmail())

	err = s.UserApplication.Update(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthPassword - パスワード更新
func (s *AuthServer) UpdateAuthPassword(
	ctx context.Context, req *pb.UpdateAuthPasswordRequest,
) (*pb.AuthResponse, error) {
	u, err := s.UserApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AuthRequestValidation.UpdateAuthPassword(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Password = req.GetPassword()

	err = s.UserApplication.UpdatePassword(ctx, u)
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
	u, err := s.UserApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AuthRequestValidation.UpdateAuthProfile(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Username = req.GetUsername()
	u.Gender = int(req.GetGender())
	u.ThumbnailURL = req.GetThumbnailUrl()
	u.SelfIntroduction = req.GetSelfIntroduction()

	err = s.UserApplication.Update(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAuthResponse(u)
	return res, nil
}

// UpdateAuthAddress - 住所更新
func (s *AuthServer) UpdateAuthAddress(
	ctx context.Context, req *pb.UpdateAuthAddressRequest,
) (*pb.AuthResponse, error) {
	u, err := s.UserApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AuthRequestValidation.UpdateAuthAddress(req)
	if err != nil {
		return nil, errorHandling(err)
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

	err = s.UserApplication.Update(ctx, u)
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
			u, err := s.UserApplication.Authentication(ctx)
			if err != nil {
				return errorHandling(err)
			}

			// 分割して送信されてきたサムネイルのバイナリをまとめる
			thumbnail := []byte{}
			for i := 0; i < len(thumbnailBytes); i++ {
				thumbnail = append(thumbnail, thumbnailBytes[i]...)
			}

			thumbnailURL, err := s.UserApplication.UploadThumbnail(ctx, u.ID, thumbnail)
			if err != nil {
				return errorHandling(err)
			}

			res := getAuthThumbnailResponse(thumbnailURL)
			return stream.SendAndClose(res)
		}

		if err != nil {
			return errorHandling(err)
		}

		err = s.AuthRequestValidation.UploadAuthThumbnail(req)
		if err != nil {
			return errorHandling(err)
		}

		num := int(req.GetPosition())
		if thumbnailBytes[num] != nil {
			err = xerrors.New("Position is duplicated")
			return errorHandling(exception.InvalidRequestValidation.New(err))
		}

		thumbnailBytes[num] = req.GetThumbnail()
	}
}

// DeleteAuth - ユーザー退会
func (s *AuthServer) DeleteAuth(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	u, err := s.UserApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.UserApplication.Delete(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.Empty{}, nil
}

// RegisterAuthDevice - デバイス情報登録
func (s *AuthServer) RegisterAuthDevice(
	ctx context.Context, req *pb.RegisterAuthDeviceRequest,
) (*pb.AuthResponse, error) {
	res := getAuthResponse(&user.User{})
	return res, nil
}

func getAuthResponse(u *user.User) *pb.AuthResponse {
	return &pb.AuthResponse{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           pb.Gender(u.Gender),
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             pb.Role(u.Role),
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
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}
}

func getAuthThumbnailResponse(thumbnailURL string) *pb.AuthThumbnailResponse {
	return &pb.AuthThumbnailResponse{
		ThumbnailUrl: thumbnailURL,
	}
}
