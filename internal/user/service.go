package user

import (
	"context"

	"grpc/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedUserServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	if req.UserId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user_id must be > 0")
	}

	// Fake user:
	u := &proto.User{
		Id:    req.UserId,
		Name:  "Fefi",
		Email: "fefi@example.com",
	}

	return &proto.GetUserResponse{
		User: u,
	}, nil
}
