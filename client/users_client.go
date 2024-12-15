package main

import (
	"context"
	"log"
	"time"

	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main_1() {
	// Use context for gRPC dialing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to the gRPC server using insecure credentials
	conn, err := grpc.DialContext(ctx, ":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Users gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a UserService client
	client := users.NewUserServiceClient(conn)

	// Create a new user
	userName := "Jane Doe"
	createResp, err := createUser(client, userName)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	log.Printf("User created: ID=%d, Name=%s", createResp.UserId, userName)

	// Get user by ID
	getResp, err := getUser(client, createResp.UserId) 
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	log.Printf("Fetched user: ID=%d, Name=%s", createResp.UserId, getResp.Name)

	// // Delete the user
	deleteResp, err := deleteUser(client, createResp.UserId)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	log.Printf("User deleted: %s", deleteResp.DeletionMessage)
}

// createUser calls the CreateUser gRPC method
func createUser(client users.UserServiceClient, name string) (*users.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &users.CreateUserRequest{
		Name: name,
	}
	return client.CreateUser(ctx, req)
}

// getUser calls the GetUserID gRPC method
func getUser(client users.UserServiceClient, userID int32) (*users.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &users.GetUserRequest{
		UserId: userID,
	}
	return client.GetUserID(ctx, req)
}

// deleteUser calls the DeleteUser gRPC method
func deleteUser(client users.UserServiceClient, userID int32) (*users.DeleteUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &users.DeleteUserRequest{
		UserId: userID,
	}
	return client.DeleteUser(ctx, req)
}
