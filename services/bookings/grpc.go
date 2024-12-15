package main

import (
	"database/sql"
	"log"
	"net"

	handler "github.com/Nabeeha-Mudassir/RideBookingService/services/bookings/handler"
	"github.com/Nabeeha-Mudassir/RideBookingService/services/bookings/service"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewgRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func initDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Create Users table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)`)
	if err != nil {
		return nil, err
	}

	// Create Rides table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS rides (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			source TEXT NOT NULL,
			destination TEXT NOT NULL,
			distance INTEGER NOT NULL,
			cost INTEGER NOT NULL
		)`)
	if err != nil {
		return nil, err
	}

	// Create Bookings table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS bookings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			ride_id INTEGER NOT NULL,
			timestamp DATETIME NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(ride_id) REFERENCES rides(id)
		)`)
	if err != nil {
		return nil, err
	}

	log.Println("Database tables created or already exist.")
	return db, nil
}

func SeedRides(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO rides (id, source, destination, distance, cost)
		VALUES
			(123, 'Lahore', 'Islamabad', 300, 500000)
		ON CONFLICT(id) DO NOTHING
	`)
	return err
}

func SeedUsers(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO users (id, name)
		VALUES
			(1, 'Alice'),
			(2, 'Bob'),
			(3, 'Charlie')
		ON CONFLICT(id) DO NOTHING
	`)
	return err
}

func SeedBookings(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO bookings (id, user_id, ride_id, timestamp)
		VALUES
			(1, 1, 123, '2024-12-15 10:00:00')
		ON CONFLICT(id) DO NOTHING
	`)
	return err
}

func (s *gRPCServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Initialize SQLite DB
	db, err := sql.Open("sqlite3", "./ridebooking.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}
	defer db.Close()

	_db, _err := initDB("./ridebooking.db")

	if _err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Seed data
	if err := SeedRides(_db); err != nil {
		log.Fatalf("Failed to seed rides: %v", err)
	}
	if err := SeedUsers(_db); err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
	if err := SeedBookings(_db); err != nil {
		log.Fatalf("Failed to seed bookings: %v", err)
	}

	gRPCServer := grpc.NewServer()

	// Register the service
	bookingService := service.NewBookingService(_db) // Pass DB here
	handler.NewBookingsGrpService(gRPCServer, bookingService)

	log.Println("Bookings gRPC server is running on port", s.addr)
	return gRPCServer.Serve(listener)
}
