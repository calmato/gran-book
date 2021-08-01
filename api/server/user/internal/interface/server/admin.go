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
	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"golang.org/x/xerrors"
)

// AdminServer - Adminサービスインターフェースの構造体
type AdminServer struct {
	pb.UnimplementedAdminServiceServer
	AdminRequestValidation validation.AdminRequestValidation
	UserApplication        application.UserApplication
}

// ListAdmin - 管理者一覧取得
func (s *AdminServer) ListAdmin(ctx context.Context, req *pb.ListAdminRequest) (*pb.AdminListResponse, error) {
	_, err := s.UserApplication.Authorization(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AdminRequestValidation.ListAdmin(req)
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

	us, total, err := s.UserApplication.ListAdmin(ctx, q)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminListResponse(us, limit, offset, total)
	return res, nil
}

// GetAdmin - 管理者取得
func (s *AdminServer) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.AdminResponse, error) {
	_, err := s.UserApplication.Authorization(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AdminRequestValidation.GetAdmin(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.UserApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// CreateAdmin - 管理者登録
func (s *AdminServer) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.AdminResponse, error) {
	role, err := s.UserApplication.Authorization(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.UserApplication.HasAdminRole(role)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AdminRequestValidation.CreateAdmin(req)
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

	err = s.UserApplication.Create(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminContact - 連絡先更新
func (s *AdminServer) UpdateAdminContact(
	ctx context.Context, req *pb.UpdateAdminContactRequest,
) (*pb.AdminResponse, error) {
	role, err := s.UserApplication.Authorization(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.UserApplication.HasAdminRole(role)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AdminRequestValidation.UpdateAdminContact(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.UserApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Email = strings.ToLower(req.GetEmail())
	u.PhoneNumber = req.GetPhoneNumber()

	err = s.UserApplication.Update(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminPassword - パスワード更新
func (s *AdminServer) UpdateAdminPassword(
	ctx context.Context, req *pb.UpdateAdminPasswordRequest,
) (*pb.AdminResponse, error) {
	role, err := s.UserApplication.Authorization(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.UserApplication.HasAdminRole(role)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AdminRequestValidation.UpdateAdminPassword(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.UserApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	u.Password = req.GetPassword()

	err = s.UserApplication.UpdatePassword(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getAdminResponse(u)
	return res, nil
}

// UpdateAdminProfile - プロフィール更新
func (s *AdminServer) UpdateAdminProfile(
	ctx context.Context, req *pb.UpdateAdminProfileRequest,
) (*pb.AdminResponse, error) {
	role, err := s.UserApplication.Authorization(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.UserApplication.HasAdminRole(role)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.AdminRequestValidation.UpdateAdminProfile(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.UserApplication.GetAdmin(ctx, req.GetUserId())
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

	err = s.UserApplication.Update(ctx, u)
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
	userID := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			role, err := s.UserApplication.Authorization(ctx)
			if err != nil {
				return errorHandling(err)
			}

			err = s.UserApplication.HasAdminRole(role)
			if err != nil {
				return errorHandling(err)
			}

			u, err := s.UserApplication.GetAdmin(ctx, userID)
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

			res := getAdminThumbnailResponse(thumbnailURL)
			return stream.SendAndClose(res)
		}

		if err != nil {
			return errorHandling(err)
		}

		err = s.AdminRequestValidation.UploadAdminThumbnail(req)
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
func (s *AdminServer) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminRequest) (*pb.Empty, error) {
	role, err := s.UserApplication.Authorization(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.UserApplication.HasAdminRole(role)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.UserApplication.GetAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	err = s.UserApplication.DeleteAdmin(ctx, u)
	if err != nil {
		return nil, errorHandling(err)
	}

	return &pb.Empty{}, nil
}

func getAdminListResponse(us []*user.User, limit, offset, total int) *pb.AdminListResponse {
	users := make([]*pb.AdminListResponse_User, len(us))
	for i, u := range us {
		user := &pb.AdminListResponse_User{
			Id:               u.ID,
			Username:         u.Username,
			Email:            u.Email,
			PhoneNumber:      u.PhoneNumber,
			Role:             pb.Role(u.Role),
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

	return &pb.AdminListResponse{
		Users:  users,
		Limit:  int64(limit),
		Offset: int64(offset),
		Total:  int64(total),
	}
}

func getAdminResponse(u *user.User) *pb.AdminResponse {
	return &pb.AdminResponse{
		Id:               u.ID,
		Username:         u.Username,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             pb.Role(u.Role),
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

func getAdminThumbnailResponse(thumbnailURL string) *pb.AdminThumbnailResponse {
	return &pb.AdminThumbnailResponse{
		ThumbnailUrl: thumbnailURL,
	}
}
