package billing

import (
	"context"

	"grpc/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedBillingServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListInvoices(ctx context.Context, req *proto.ListInvoicesRequest) (*proto.ListInvoicesResponse, error) {
	if req.UserId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user_id must be > 0")
	}

	invoices := []*proto.Invoice{
		{Id: 1, UserId: req.UserId, Amount: 100, Currency: "USD"},
		{Id: 2, UserId: req.UserId, Amount: 200, Currency: "EUR"},
		{Id: 3, UserId: req.UserId, Amount: 300, Currency: "GBP"},
		{Id: 4, UserId: req.UserId, Amount: 400, Currency: "JPY"},
		{Id: 5, UserId: req.UserId, Amount: 500, Currency: "KRW"},
		{Id: 6, UserId: req.UserId, Amount: 600, Currency: "CNY"},
		{Id: 7, UserId: req.UserId, Amount: 700, Currency: "HKD"},
		{Id: 8, UserId: req.UserId, Amount: 800, Currency: "AUD"},
		{Id: 9, UserId: req.UserId, Amount: 900, Currency: "CAD"},
	}

	return &proto.ListInvoicesResponse{
		Invoices: invoices,
	}, nil
}
