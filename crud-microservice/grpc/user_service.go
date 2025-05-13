package grpc

import (
	"context"

	"crud-microservice/crud-microservice/pb"
	"crud-microservice/models"
	"crud-microservice/store"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	Store store.UserStore
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := models.User{Name: req.Name, Email: req.Email}
	s.Store.Create(user)
	return &pb.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

// Implement GetUser, GetAllUsers, UpdateUser, DeleteUser similarly.
