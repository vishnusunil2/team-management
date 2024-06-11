package auth_service

import "github.com/golang-jwt/jwt"

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserSignupRequest struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
type CustomClaims struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
