package v1

import (
	"context"

	pb "github.com/calmato/gran-book/api/user/proto"
)

// UserServer - Userインターフェースの構造体
type UserServer struct {
	pb.UnimplementedUserServiceServer
}

// CreateUser - ユーザ登録
func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}
