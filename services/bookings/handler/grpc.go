package handler

import (
	"context"
	

	"github.com/Nabeeha-Mudassir/RideBookingService/services/bookings/types"
	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/bookings"
	"google.golang.org/grpc"
)


type BookingsGrpcHandler struct{
	//service injection
	bookingService types.BookingService
	bookings.UnimplementedBookingServiceServer
}

func NewBookingsGrpService(grpc *grpc.Server,bookingService types.BookingService)  {
	gRPCHandler := &BookingsGrpcHandler{
		bookingService: bookingService, 
	}

	//register the UserServiceServer 
	bookings.RegisterBookingServiceServer(grpc, gRPCHandler)
}

func (h *BookingsGrpcHandler) CreateBooking(ctx context.Context, req *bookings.CreateBookingRequest) (*bookings.CreateBookingResponse, error) {
	

	resp, err := h.bookingService.CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	
	return resp, nil
}

func (h *BookingsGrpcHandler) GetBooking(ctx context.Context, req *bookings.GetBookingRequest) (*bookings.GetBookingResponse, error) {
	booking,err := h.bookingService.GetBooking(ctx, req.BookingId)
	if err != nil {
		return nil, err
	}

	return &bookings.GetBookingResponse{Booking: booking}, nil
}

func (h *BookingsGrpcHandler) UpdateRide(ctx context.Context, req *bookings.RideRequest) (*bookings.RideResponse, error) {
	err := h.bookingService.UpdateRide(ctx, req.Ride)
	if err != nil {
		return nil, err
	}

	return &bookings.RideResponse{
		Message: "Ride updated successfully",
	}, nil
}
