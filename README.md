# gRPC Demo

## What is gRPC?

gRPC is a high-performance, open-source RPC (Remote Procedure Call) framework developed by Google. It uses Protocol Buffers (protobuf) as its interface definition language and message format, enabling efficient communication between services. Key benefits include:

- **Performance**: Uses HTTP/2 and binary serialization for faster communication
- **Type Safety**: Strongly typed contracts defined in `.proto` files
- **Language Agnostic**: Generate client/server code for multiple languages
- **Streaming**: Supports unary, server streaming, client streaming, and bidirectional streaming

## Project Structure

This project demonstrates a microservices architecture with three gRPC services:

- **Auth Service** (port 50051): Handles user authentication and token validation
- **Billing Service** (port 50052): Manages invoices and billing data
- **User Service** (port 50053): Provides user information

The demo client connects to all three services and demonstrates a typical workflow: login → validate token → get user → list invoices.

## Prerequisites

- Go 1.25.1 or later
- Protocol Buffer compiler (protoc) - optional, only needed if regenerating proto files

## Protocol Buffer Contract

The service contracts are defined in `proto/app.proto` using Protocol Buffers. This file defines the service interfaces and message types that both clients and servers use to communicate.

### Example: User Service Contract

Here's how the `User` message and `UserService` are defined:

```protobuf
service UserService {
    rpc GetUser (GetUserRequest) returns (GetUserResponse);
}

message GetUserRequest {
    int64 user_id = 1;
}

message GetUserResponse {
    User user = 1;
}

message User {
    int64 id = 1;
    string name = 2;
    string email = 3;
}
```

This contract defines:
- **Service**: `UserService` with a `GetUser` RPC method
- **Request**: `GetUserRequest` takes a `user_id` (int64)
- **Response**: `GetUserResponse` returns a `User` message
- **Message**: `User` contains `id`, `name`, and `email` fields

The numbers (1, 2, 3) are field tags used for binary encoding and must be unique within each message.

### Generating gRPC Code

To generate Go code from the proto file, use the Protocol Buffer compiler:

```bash
protoc --go_out=. --go-grpc_out=. proto/app.proto
```

This command:
- `--go_out=.`: Generates Go message types (e.g., `app.pb.go`)
- `--go-grpc_out=.`: Generates Go gRPC client/server code (e.g., `app_grpc.pb.go`)

**Note**: The generated files (`app.pb.go` and `app_grpc.pb.go`) are already included in this repository, so you only need to run this command if you modify `proto/app.proto`.

## Running the Demo

### 1. Install Dependencies

```bash
go mod download
```

### 2. Start the Services

Open three separate terminal windows and run each service:

**Terminal 1 - Auth Service:**
```bash
go run cmd/auth/main.go
```

**Terminal 2 - Billing Service:**
```bash
go run cmd/billing/main.go
```

**Terminal 3 - User Service:**
```bash
go run cmd/user/main.go
```

### 3. Run the Demo Client

Once all three services are running, execute the demo client:

```bash
go run cmd/demo-client/main.go
```

The client will:
1. Connect to the Auth service and perform a login
2. Validate the received token
3. Retrieve user information from the User service
4. List invoices from the Billing service

You should see output showing the successful execution of each step.

```bash
login successful: token=token-1763429986495767300, user_id=1
token validation successful: valid=true, user_id=1
user retrieved: id=1, name=Fefi, email=fefi@example.com
invoices retrieved: 9
invoice: id=1, user_id=1, amount=100.000000, currency=USD
invoice: id=2, user_id=1, amount=200.000000, currency=EUR
invoice: id=3, user_id=1, amount=300.000000, currency=GBP
invoice: id=4, user_id=1, amount=400.000000, currency=JPY
invoice: id=5, user_id=1, amount=500.000000, currency=KRW
invoice: id=6, user_id=1, amount=600.000000, currency=CNY
invoice: id=7, user_id=1, amount=700.000000, currency=HKD
invoice: id=8, user_id=1, amount=800.000000, currency=AUD
invoice: id=9, user_id=1, amount=900.000000, currency=CAD
demo completed successfully
```

## Service Ports

- Auth Service: `localhost:50051`
- Billing Service: `localhost:50052`
- User Service: `localhost:50053`

