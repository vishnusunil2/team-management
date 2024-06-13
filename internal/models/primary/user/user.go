package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;unique"`
	Email     string    `gorm:"not null"`
	Phone     string    `gorm:"null"`
	Password  string    `gorm:"null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"null"`
	RoleId    int
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
		RoleId:    2,
		Phone:     phone,
		CreatedAt: time.Now(),
	}
}
