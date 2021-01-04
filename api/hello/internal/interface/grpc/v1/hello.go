package v1

import (
	"context"

	pb "github.com/calmato/gran-book/api/hello/proto"
)

type HelloServer struct {
	pb.UnimplementedGreeterServer
}

// type user struct {
// 	ID   string
// 	Name string
// }

// var (
// 	users = []*user{
// 		&user{
// 			ID:   "0",
// 			Name: "Hamada",
// 		},
// 		&user{
// 			ID:   "1",
// 			Name: "Inatomi",
// 		},
// 		&user{
// 			ID:   "2",
// 			Name: "Nishikawa",
// 		},
// 	}
// )

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	res := &pb.HelloReply{
		Message: req.Name + " world!",
	}

	return res, nil
}

// func (s *HelloServer) GetUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.UsersResponse, error) {
// 	urs := make([]*pb.User, len(users))
// 	for i, u := range users {
// 		ur := &pb.User{
// 			Id:   u.ID,
// 			Name: u.Name,
// 		}

// 		urs[i] = ur
// 	}

// 	res := &pb.UsersResponse{
// 		Users: urs,
// 	}

// 	return res, nil
// }

// func (s *HelloServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
// 	id, err := strconv.Atoi(req.Id)
// 	if err != nil {
// 		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	if id >= len(users) {
// 		msg := "Not found"
// 		return &pb.UserResponse{}, status.Error(codes.NotFound, msg)
// 	}

// 	u := users[id]
// 	if u == nil {
// 		msg := "Not found"
// 		return &pb.UserResponse{}, status.Error(codes.NotFound, msg)
// 	}

// 	res := &pb.UserResponse{
// 		Id:   u.ID,
// 		Name: u.Name,
// 	}

// 	return res, nil
// }

// func (s *HelloServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
// 	if req.Name == "" {
// 		msg := "Invalid argument: name is required"
// 		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, msg)
// 	}

// 	id := strconv.Itoa(len(users))
// 	u := &user{ID: id, Name: req.Name}

// 	users = append(users, u)

// 	res := &pb.UserResponse{
// 		Id:   u.ID,
// 		Name: u.Name,
// 	}

// 	return res, nil
// }

// func (s *HelloServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
// 	if req.Name == "" {
// 		msg := "Invalid argument: name is required"
// 		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, msg)
// 	}

// 	id, err := strconv.Atoi(req.Id)
// 	if err != nil {
// 		return &pb.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	u := users[id]
// 	if u == nil {
// 		msg := "Not found"
// 		return &pb.UserResponse{}, status.Error(codes.NotFound, msg)
// 	}

// 	u.Name = req.Name

// 	res := &pb.UserResponse{
// 		Id:   u.ID,
// 		Name: u.Name,
// 	}

// 	return res, nil
// }

// func (s *HelloServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.EmptyResponse, error) {
// 	id, err := strconv.Atoi(req.Id)
// 	if err != nil {
// 		return &pb.EmptyResponse{}, err
// 	}

// 	if id >= len(users) {
// 		msg := "Not found"
// 		return &pb.EmptyResponse{}, status.Error(codes.NotFound, msg)
// 	}

// 	users = append(users[:id], users[id+1:]...)

// 	return &pb.EmptyResponse{}, nil
// }
