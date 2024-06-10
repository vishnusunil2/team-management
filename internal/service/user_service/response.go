package user_service

import "github.com/google/uuid"

type UserResponse struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}
