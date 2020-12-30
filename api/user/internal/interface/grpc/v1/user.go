package v1

import (
	"context"
	"log"

	"github.com/calmato/gran-book/api/user/internal/application"
	"github.com/calmato/gran-book/api/user/internal/application/input"
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

	log.Println("user", u)
	log.Println("err", err)

	return &pb.UserResponse{}, nil
}
