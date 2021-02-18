package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"golang.org/x/xerrors"
)

// AdminServer - Authインターフェースの構造体
type AdminServer struct {
	pb.UnimplementedAdminServiceServer
	AuthApplication  application.AuthApplication
	AdminApplication application.AdminApplication
}

// CreateAdmin - 管理者登録
func (s *AdminServer) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.AdminResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = hasAdminRole(cu, "")
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.CreateAdmin{
		Username:             req.Username,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		Role:                 req.Role,
		LastName:             req.LastName,
		FirstName:            req.FirstName,
		LastNameKana:         req.LastNameKana,
		FirstNameKana:        req.FirstNameKana,
	}

	u, err := s.AdminApplication.Create(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminRole - 管理者権限更新
func (s *AdminServer) UpdateAdminRole(ctx context.Context, req *pb.UpdateAdminRoleRequest) (*pb.AdminResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = hasAdminRole(cu, req.Id)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAdminRole{
		Role: req.Role,
	}

	u, err := s.AdminApplication.UpdateRole(ctx, in, req.Id)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminPassword - 管理者パスワード更新
func (s *AdminServer) UpdateAdminPassword(
	ctx context.Context, req *pb.UpdateAdminPasswordRequest,
) (*pb.AdminResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = hasAdminRole(cu, req.Id)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAdminPassword{
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	u, err := s.AdminApplication.UpdatePassword(ctx, in, req.Id)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminProfile - 管理者プロフィール更新
func (s *AdminServer) UpdateAdminProfile(
	ctx context.Context, req *pb.UpdateAdminProfileRequest,
) (*pb.AdminResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = hasAdminRole(cu, req.Id)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAdminProfile{
		Username:      req.Username,
		Email:         req.Email,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		LastNameKana:  req.LastNameKana,
		FirstNameKana: req.FirstNameKana,
	}

	u, err := s.AdminApplication.UpdateProfile(ctx, in, req.Id)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

func hasAdminRole(u *user.User, uid string) error {
	if u == nil || u.ID == uid {
		err := xerrors.New("This ID belongs to the current user")
		return exception.Forbidden.New(err)
	}

	if u.Role != user.AdminRole {
		err := xerrors.New("This account doesn't have administrator privileges")
		return exception.Forbidden.New(err)
	}

	return nil
}

func getAdminResponse(u *user.User) *pb.AdminResponse {
	return &pb.AdminResponse{
		Id:               u.ID,
		Username:         u.Username,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             u.Role,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		Activated:        u.Activated,
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}
}
