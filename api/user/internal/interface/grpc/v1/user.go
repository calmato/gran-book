package v1

import (
	"context"

	"github.com/calmato/gran-book/api/user/internal/application"
	"github.com/calmato/gran-book/api/user/internal/application/input"
	"github.com/calmato/gran-book/api/user/lib/datetime"
	pb "github.com/calmato/gran-book/api/user/proto"
)

// UserServer - Userインターフェースの構造体
type UserServer struct {
	pb.UnimplementedUserServiceServer
	UserApplication application.UserApplication
}

// CreateUser - ユーザ登録
func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	in := &input.CreateUser{
		Username:             req.Username,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	u, err := s.UserApplication.Create(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := &pb.UserResponse{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           u.Gender,
		Email:            u.Email,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		Lastname:         u.Lastname,
		Firstname:        u.Firstname,
		LastnameKana:     u.LastnameKana,
		FirstnameKana:    u.FirstnameKana,
		PostalCode:       u.PostalCode,
		Prefecture:       u.Prefecture,
		City:             u.City,
		AddressLine1:     u.AddressLine1,
		AddressLine2:     u.AddressLine2,
		PhoneNumber:      u.PhoneNumber,
		Role:             u.Role,
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}

	return res, nil
}
