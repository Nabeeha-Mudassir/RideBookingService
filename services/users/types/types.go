package types

import (
	"context"

	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/users"
)

type UserService interface {
	CreateUser(context.Context, *users.User) error
	GetUser(context.Context, int32) (*users.User, error)
	DeleteUser(context.Context, int32) error
	
}