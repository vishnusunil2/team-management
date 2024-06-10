package user_repo

import "github.com/golang-jwt/jwt"

type CreateUserRequest struct {
	Email     string
	Phone     string
	FirstName string
	LastName  string
	Password  string
}
type CustomClaims struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
