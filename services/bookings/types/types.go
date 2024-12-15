package types

import (
	"context"

	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/bookings"
)

type BookingService interface {
	CreateBooking(context.Context, *bookings.CreateBookingRequest) (*bookings.CreateBookingResponse, error)
	GetBooking(context.Context, int32) (*bookings.Booking, error)
	UpdateRide(context.Context, *bookings.Ride) error
}