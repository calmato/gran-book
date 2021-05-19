package v1

import (
	"context"
	"io"

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
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	}

	if o := req.GetOrder(); o != nil {
		in.By = o.GetBy()
		in.Direction = o.GetDirection()
	}

	us, out, err := s.AdminApplication.List(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminListResponse(us, out)
	return res, nil
}

// SearchAdmin - 管理者一覧取得
func (s *AdminServer) SearchAdmin(ctx context.Context, req *pb.SearchAdminRequest) (*pb.AdminListResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.SearchAdmin{
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	}

	if o := req.GetOrder(); o != nil {
		in.By = o.GetBy()
		in.Direction = o.GetDirection()
	}

	if s := req.GetSearch(); s != nil {
		in.Field = s.GetField()
		in.Value = s.GetValue()
	}

	us, out, err := s.AdminApplication.Search(ctx, in)
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

	u, err := s.AdminApplication.Show(ctx, req.GetId())
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
		Username:             req.GetUsername(),
		Email:                req.GetEmail(),
		Password:             req.GetPassword(),
		PasswordConfirmation: req.GetPasswordConfirmation(),
		Role:                 int(req.GetRole()),
		LastName:             req.GetLastName(),
		FirstName:            req.GetFirstName(),
		LastNameKana:         req.GetLastNameKana(),
		FirstNameKana:        req.GetFirstNameKana(),
	}

	u, err := s.AdminApplication.Create(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminContact - 管理者連絡先更新
func (s *AdminServer) UpdateAdminContact(
	ctx context.Context, req *pb.UpdateAdminContactRequest,
) (*pb.AdminResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = hasAdminRole(cu, req.GetId())
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAdminContact{
		Email:       req.GetEmail(),
		PhoneNumber: req.GetPhoneNumber(),
	}

	u, err := s.AdminApplication.UpdateContact(ctx, in, req.GetId())
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

	err = hasAdminRole(cu, req.GetId())
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAdminPassword{
		Password:             req.GetPassword(),
		PasswordConfirmation: req.GetPasswordConfirmation(),
	}

	u, err := s.AdminApplication.UpdatePassword(ctx, in, req.GetId())
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

	err = hasAdminRole(cu, req.GetId())
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.UpdateAdminProfile{
		Username:      req.GetUsername(),
		LastName:      req.GetLastName(),
		FirstName:     req.GetFirstName(),
		LastNameKana:  req.GetLastNameKana(),
		FirstNameKana: req.GetFirstNameKana(),
		Role:          int(req.GetRole()),
		ThumbnailURL:  req.GetThumbnailUrl(),
	}

	u, err := s.AdminApplication.UpdateProfile(ctx, in, req.GetId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UploadAdminThumbnail - サムネイルアップロード
func (s *AdminServer) UploadAdminThumbnail(stream pb.AdminService_UploadAdminThumbnailServer) error {
	ctx := stream.Context()
	thumbnailBytes := map[int][]byte{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			_, err := s.AuthApplication.Authentication(ctx)
			if err != nil {
				return errorHandling(err)
			}

			// 分割して送信されてきたサムネイルのバイナリをまとめる
			thumbnail := []byte{}
			for i := 0; i < len(thumbnailBytes); i++ {
				thumbnail = append(thumbnail, thumbnailBytes[i]...)
			}

			in := &input.UploadAdminThumbnail{
				Thumbnail: thumbnail,
			}

			thumbnailURL, err := s.AdminApplication.UploadThumbnail(ctx, in, req.GetUserId())
			if err != nil {
				return errorHandling(err)
			}

			res := &pb.AdminThumbnailResponse{
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

// DeleteAdmin - 管理者権限削除
func (s *AdminServer) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminRequest) (*pb.EmptyUser, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = hasAdminRole(cu, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AdminApplication.Delete(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.EmptyUser{}, nil
}

func getAdminResponse(u *user.User) *pb.AdminResponse {
	return &pb.AdminResponse{
		Id:               u.ID,
		Username:         u.Username,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             int32(u.Role),
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
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
			Role:             int32(u.Role),
			ThumbnailUrl:     u.ThumbnailURL,
			SelfIntroduction: u.SelfIntroduction,
			LastName:         u.LastName,
			FirstName:        u.FirstName,
			LastNameKana:     u.LastNameKana,
			FirstNameKana:    u.FirstNameKana,
			CreatedAt:        datetime.TimeToString(u.CreatedAt),
			UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
		}

		users[i] = user
	}

	res := &pb.AdminListResponse{
		Users:  users,
		Limit:  int64(out.Limit),
		Offset: int64(out.Offset),
		Total:  int64(out.Total),
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
