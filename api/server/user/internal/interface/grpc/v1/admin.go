package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

// AdminServer - Authインターフェースの構造体
type AdminServer struct {
	pb.UnimplementedAdminServiceServer
	AuthApplication  application.AuthApplication
	AdminApplication application.AdminApplication
}

// ListAdmin - 管理者一覧取得
func (s *AdminServer) ListAdmin(ctx context.Context, req *pb.ListAdminRequest) (*pb.AdminListResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.ListAdmin{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	if req.Order != nil {
		in.By = req.Order.By
		in.Direction = req.Order.Direction
	}

	us, out, err := s.AdminApplication.List(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminListResponse(us, out)
	return res, nil
}

// GetAdmin - 管理者情報取得
func (s *AdminServer) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.AdminResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.AdminApplication.Show(ctx, req.Id)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// CreateAdmin - 管理者登録
func (s *AdminServer) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.AdminResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
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

	err = authorization(cu)
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

	err = authorization(cu)
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

	err = authorization(cu)
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

func getAdminListResponse(us []*user.User, out *output.ListQuery) *pb.AdminListResponse {
	users := make([]*pb.AdminListResponse_User, len(us))
	for i, u := range us {
		user := &pb.AdminListResponse_User{
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

		users[i] = user
	}

	res := &pb.AdminListResponse{
		Users:  users,
		Limit:  out.Limit,
		Offset: out.Offset,
		Total:  out.Total,
	}

	if out.Order != nil {
		order := &pb.AdminListResponse_Order{
			By:        out.Order.By,
			Direction: out.Order.Direction,
		}

		res.Order = order
	}

	return res
}
