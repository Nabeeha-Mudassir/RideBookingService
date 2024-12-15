package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/bookings"
)

func main() {
	// Set up the connection to the server using NewClient
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials())) // Replace with your server address
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new BookingServiceClient
	client := bookings.NewBookingServiceClient(conn)


	bookingID := createBooking(client)
	if bookingID != 0 {
		getBooking(client, bookingID)
	} else {
		log.Println("Failed to create booking; skipping GetBooking.")
	}

	// // Example: Update Ride
	// updateRide(client)
}

func getBooking(client bookings.BookingServiceClient, bookingID int32) {
	req := &bookings.GetBookingRequest{
		BookingId: bookingID,
	}

	resp, err := client.GetBooking(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting booking: %v", err)
	}

	fmt.Printf("Booking details: %+v\n", resp)
}

func createBooking(client bookings.BookingServiceClient) int32 {
	req := &bookings.CreateBookingRequest{
		UserId: 7,
		Ride: &bookings.Ride{
			RideId:     123,
			Source:     "Lahore",
			Destination: "Islamabad",
			Distance:   300,
			Cost:       500000,
		},
	}

	resp, err := client.CreateBooking(context.Background(), req)
	if err != nil {
		log.Fatalf("Error creating booking: %v", err)
		return 0
	}

	fmt.Printf("Created Booking: %+v\n", resp.Booking.BookingId)
	return resp.Booking.BookingId
}

func updateRide(client bookings.BookingServiceClient) {
	// Example RideRequest
	req := &bookings.RideRequest{
		RideId: 123, 
	}

	// Invoke the UpdateRide method
	resp, err := client.UpdateRide(context.Background(), req)
	if err != nil {
		log.Fatalf("Error updating ride: %v", err)
	}

	// Handle the response
	fmt.Printf("Updated Ride: %+v\n", resp)
}
