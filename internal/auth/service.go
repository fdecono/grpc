package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"grpc/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedAuthServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email and password are required")
	}

	// Fake auth: always succeed and return user_id = 1
	token := fmt.Sprintf("token-%d", time.Now().UnixNano())

	return &proto.LoginResponse{
		Token:  token,
		UserId: 1,
	}, nil
}

func (s *Service) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	if strings.HasPrefix(req.Token, "token-") {
		return &proto.ValidateTokenResponse{
			Valid:  true,
			UserId: 1,
		}, nil
	}

	return &proto.ValidateTokenResponse{
		Valid:  false,
		UserId: 0,
	}, nil
}
