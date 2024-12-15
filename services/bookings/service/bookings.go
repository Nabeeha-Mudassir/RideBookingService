package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	
	"time"

	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/bookings"
)

type BookingService struct {
	db *sql.DB
}

func NewBookingService(db *sql.DB) *BookingService {
	return &BookingService{db: db}
}

func (s *BookingService) CreateBooking(ctx context.Context, booking *bookings.CreateBookingRequest) (*bookings.CreateBookingResponse, error) {
	result, err := s.db.Exec(`
		INSERT INTO bookings (user_id, ride_id, timestamp)
		VALUES (?, ?, ?)`,
		booking.UserId, booking.Ride.RideId, time.Now(),
	)
	if err != nil {
		return nil, err
	}
	// Scan the returned values into the booking struct
	bookingID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
		
	}

	new_booking := &bookings.Booking{
		BookingId: int32(bookingID),
		UserId:    booking.UserId,
		RideId:    booking.Ride.RideId,
	}
	return &bookings.CreateBookingResponse{
		Booking: new_booking,
	}, nil
}

func (s *BookingService) GetBooking(ctx context.Context, bookingID int32) (*bookings.Booking, error) {
	
	row := s.db.QueryRow(`
		SELECT b.id, b.user_id, b.ride_id, b.timestamp, r.source, r.destination, r.distance, r.cost
		FROM bookings b
		JOIN rides r ON b.ride_id = r.id
		WHERE b.id = ?`, bookingID)

	var booking bookings.Booking
	var source, destination string
	var distance, cost int32

	err := row.Scan(
		&booking.BookingId, &booking.UserId, &booking.RideId,
		&booking.RideTime, &source, &destination, &distance, &cost,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("booking with ID %d not found", bookingID)
	} else if err != nil {
		return nil, fmt.Errorf("error fetching booking: %w", err)
	}

	bookingRide := &bookings.Booking{
		BookingId: bookingID,
		RideId:     booking.RideId,
		UserId:    booking.UserId,
		RideTime: booking.RideTime,
	}

	return bookingRide, nil
}

func (s *BookingService) UpdateRide(ctx context.Context, ride *bookings.Ride) error {
	result, err := s.db.Exec(`
		UPDATE rides
		SET source = ?, destination = ?, distance = ?, cost = ?
		WHERE id = ?`,
		ride.Source, ride.Destination, ride.Distance, ride.Cost, ride.RideId,
	)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("ride not found")
	}
	return nil
}
