package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"primarKey;unique"`
	Email     string    `gorm:"not null"`
	Phone     string    `gorm:"null"`
	Password  string    `gorm:"null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email, phone, firstName, lastName, password string) *User {
	return &User{
		Id:        uuid.New(),
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		CreatedAt: time.Now(),
	}
}
