package handler

import (
	"context"

	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/users"
	"github.com/Nabeeha-Mudassir/RideBookingService/services/users/types"
	"google.golang.org/grpc"
)


type UsersGrpcHandler struct{
	//service injection
	userService types.UserService
	users.UnimplementedUserServiceServer
}

func NewUsersGrpService(grpc *grpc.Server,userService types.UserService)  {
	gRPCHandler := &UsersGrpcHandler{
		userService: userService, 
	}

	//register the UserServiceServer 
	users.RegisterUserServiceServer(grpc, gRPCHandler)
}

func (h* UsersGrpcHandler) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	user := &users.User{
		Name: req.Name,
	}

	err := h.userService.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &users.CreateUserResponse{
		UserId: user.UserId,
	}
	return res, nil
}	
func (h* UsersGrpcHandler) GetUserID(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	user, err := h.userService.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &users.GetUserResponse{
		Name: user.Name,
	}
	return res, nil
}
func (h* UsersGrpcHandler) DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (*users.DeleteUserResponse, error) {
	err := h.userService.DeleteUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &users.DeleteUserResponse{
		DeletionMessage: "User successfully deleted",
	}
	return res, nil
}