package main

import (
	"database/sql"
	"log"
	"net"

	handler "github.com/Nabeeha-Mudassir/RideBookingService/services/users/handler/users"
	"github.com/Nabeeha-Mudassir/RideBookingService/services/users/service"
	"google.golang.org/grpc"

	_ "github.com/mattn/go-sqlite3"
)

type gRPCServer struct {
	addr string
	db   *sql.DB
}

func NewgRPCServer(addr string, db *sql.DB) *gRPCServer {
	return &gRPCServer{addr: addr, db: db}
}

func initDB() (*sql.DB, error) {
	// SQLite database file
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}

	// Create `users` table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	// Seed initial data
	_, err = db.Exec(`
		INSERT INTO users (name) 
		VALUES ('Alice'), ('Bob')
		ON CONFLICT DO NOTHING;
	`)
	if err != nil {
	
		log.Println("Seed data might already exist.")
	}

	return db, nil
}
func (s *gRPCServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen for user service: %v", err)
	}

	gRPCServer := grpc.NewServer()

	//Register the service 
	userService := service.NewUserService(s.db)
	handler.NewUsersGrpService(gRPCServer, userService) 


	log.Println("Users gRPC server is running on port", s.addr)
	return gRPCServer.Serve(listener)
}