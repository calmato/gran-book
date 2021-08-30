package server

import (
	"context"
	"io"
	"strings"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"golang.org/x/xerrors"
)

type adminServer struct {
	pb.UnimplementedAdminServiceServer
	adminRequestValidation validation.AdminRequestValidation
	userApplication        application.UserApplication
}

func NewAdminServer(arv validation.AdminRequestValidation, ua application.UserApplication) pb.AdminServiceServer {
	return &adminServer{
		adminRequestValidation: arv,
		userApplication:        ua,
	}
}

// ListAdmin - 管理者一覧取得
func (s *adminServer) ListAdmin(ctx context.Context, req *pb.ListAdminRequest) (*pb.AdminListResponse, error) {
	err := s.adminRequestValidation.ListAdmin(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	q := &database.ListQuery{
		Limit:      limit,
		Offset:     offset,
		Conditions: []*database.ConditionQuery{},
	}

	if req.GetSearch() != nil {
		c := &database.ConditionQuery{
			Field:    req.GetSearch().GetField(),
			Operator: "LIKE",
			Value:    req.GetSearch().GetValue(),
		}

		q.Conditions = append(q.Conditions, c)
	}

	if req.GetOrder() != nil {
		o := &database.OrderQuery{
			Field:   req.GetOrder().GetField(),
			OrderBy: int(req.GetOrder().GetOrderBy()),
		}

		q.Order = o
	}

	us, total, err := s.userApplication.ListAdmin(ctx, q)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminListResponse(us, limit, offset, total)
	return res, nil
}

// GetAdmin - 管理者取得
func (s *adminServer) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.AdminResponse, error) {
	err := s.adminRequestValidation.GetAdmin(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.userApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// CreateAdmin - 管理者登録
func (s *adminServer) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.AdminResponse, error) {
	err := s.adminRequestValidation.CreateAdmin(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u := &user.User{
		Username:      req.GetUsername(),
		Email:         strings.ToLower(req.GetEmail()),
		PhoneNumber:   req.GetPhoneNumber(),
		Password:      req.GetPassword(),
		Role:          int(req.GetRole()),
		LastName:      req.GetLastName(),
		FirstName:     req.GetFirstName(),
		LastNameKana:  req.GetLastNameKana(),
		FirstNameKana: req.GetFirstNameKana(),
	}

	err = s.userApplication.Create(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminContact - 連絡先更新
func (s *adminServer) UpdateAdminContact(
	ctx context.Context, req *pb.UpdateAdminContactRequest,
) (*pb.AdminResponse, error) {
	err := s.adminRequestValidation.UpdateAdminContact(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.userApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Email = strings.ToLower(req.GetEmail())
	u.PhoneNumber = req.GetPhoneNumber()

	err = s.userApplication.Update(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminPassword - パスワード更新
func (s *adminServer) UpdateAdminPassword(
	ctx context.Context, req *pb.UpdateAdminPasswordRequest,
) (*pb.AdminResponse, error) {
	err := s.adminRequestValidation.UpdateAdminPassword(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.userApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Password = req.GetPassword()

	err = s.userApplication.UpdatePassword(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminProfile - プロフィール更新
func (s *adminServer) UpdateAdminProfile(
	ctx context.Context, req *pb.UpdateAdminProfileRequest,
) (*pb.AdminResponse, error) {
	err := s.adminRequestValidation.UpdateAdminProfile(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.userApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Username = req.GetUsername()
	u.Role = int(req.GetRole())
	u.LastName = req.GetLastName()
	u.FirstName = req.GetFirstName()
	u.LastNameKana = req.GetLastNameKana()
	u.FirstNameKana = req.GetFirstNameKana()
	u.ThumbnailURL = req.GetThumbnailUrl()

	err = s.userApplication.Update(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UploadAdminThumbnail - サムネイルアップロード
func (s *adminServer) UploadAdminThumbnail(stream pb.AdminService_UploadAdminThumbnailServer) error {
	ctx := stream.Context()
	thumbnailBytes := map[int][]byte{}
	userID := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			u, err := s.userApplication.GetAdmin(ctx, userID)
			if err != nil {
				return errorHandling(err)
			}

			// 分割して送信されてきたサムネイルのバイナリをまとめる
			thumbnail := []byte{}
			for i := 0; i < len(thumbnailBytes); i++ {
				thumbnail = append(thumbnail, thumbnailBytes[i]...)
			}

			thumbnailURL, err := s.userApplication.UploadThumbnail(ctx, u.ID, thumbnail)
			if err != nil {
				return errorHandling(err)
			}

			res := getAdminThumbnailResponse(thumbnailURL)
			return stream.SendAndClose(res)
		}

		if err != nil {
			return errorHandling(err)
		}

		err = s.adminRequestValidation.UploadAdminThumbnail(req)
		if err != nil {
			return errorHandling(err)
		}

		if userID == "" {
			userID = req.GetUserId()
		}

		num := int(req.GetPosition())
		if thumbnailBytes[num] != nil {
			err = xerrors.New("Position is duplicated")
			return errorHandling(exception.InvalidRequestValidation.New(err))
		}

		thumbnailBytes[num] = req.GetThumbnail()
	}
}

// DeleteAdmin - 管理者権限削除
func (s *adminServer) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminRequest) (*pb.Empty, error) {
	err := s.adminRequestValidation.DeleteAdmin(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.userApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.userApplication.DeleteAdmin(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.Empty{}, nil
}

func getAdminListResponse(us user.Users, limit, offset, total int) *pb.AdminListResponse {
	return &pb.AdminListResponse{
		Admins: us.Admin(),
		Limit:  int64(limit),
		Offset: int64(offset),
		Total:  int64(total),
	}
}

func getAdminResponse(u *user.User) *pb.AdminResponse {
	return &pb.AdminResponse{
		Admin: u.Admin(),
	}
}

func getAdminThumbnailResponse(thumbnailURL string) *pb.AdminThumbnailResponse {
	return &pb.AdminThumbnailResponse{
		ThumbnailUrl: thumbnailURL,
	}
}
