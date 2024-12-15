package service

import (
	"context"
	"database/sql"

	"github.com/Nabeeha-Mudassir/RideBookingService/services/common/genproto/users"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, user *users.User) error {
	err := s.db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&user.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUser(ctx context.Context, userID int32) (*users.User, error) {
	user := &users.User{}
	err := s.db.QueryRow("SELECT id, name FROM users WHERE id = $1", userID).
		Scan(&user.UserId, &user.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, userID int32) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}
