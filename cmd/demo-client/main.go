package main

import (
	"context"
	"grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Auth client
	authConn, err := grpc.DialContext(
		ctx,
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to dial auth service: %v", err)
	}
	defer authConn.Close()

	authClient := proto.NewAuthServiceClient(authConn)

	loginRes, err := authClient.Login(ctx, &proto.LoginRequest{
		Email:    "fefi@example.com",
		Password: "password",
	})
	if err != nil {
		log.Fatalf("failed to login: %v", err)
	}
	log.Printf("login successful: token=%s, user_id=%d", loginRes.Token, loginRes.UserId)

	// Optional: Validate token
	validateRes, err := authClient.ValidateToken(ctx, &proto.ValidateTokenRequest{
		Token: loginRes.Token,
	})
	if err != nil {
		log.Fatalf("failed to validate token: %v", err)
	}
	log.Printf("token validation successful: valid=%t, user_id=%d", validateRes.Valid, validateRes.UserId)
	userID := validateRes.UserId

	// User client
	userConn, err := grpc.DialContext(
		ctx,
		"localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to dial user service: %v", err)
	}
	defer userConn.Close()

	userClient := proto.NewUserServiceClient(userConn)

	userRes, err := userClient.GetUser(ctx, &proto.GetUserRequest{
		UserId: userID,
	})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}
	log.Printf("user retrieved: id=%d, name=%s, email=%s", userRes.User.Id, userRes.User.Name, userRes.User.Email)

	// Billing client
	billingConn, err := grpc.DialContext(
		ctx,
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to dial billing service: %v", err)
	}
	defer billingConn.Close()

	billingClient := proto.NewBillingServiceClient(billingConn)
	invoicesRes, err := billingClient.ListInvoices(ctx, &proto.ListInvoicesRequest{
		UserId: userID,
	})
	if err != nil {
		log.Fatalf("failed to list invoices: %v", err)
	}
	log.Printf("invoices retrieved: %d", len(invoicesRes.Invoices))
	for _, invoice := range invoicesRes.Invoices {
		log.Printf("invoice: id=%d, user_id=%d, amount=%f, currency=%s", invoice.Id, invoice.UserId, invoice.Amount, invoice.Currency)
	}

	log.Println("demo completed successfully")
}
