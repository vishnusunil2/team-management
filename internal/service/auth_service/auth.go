package auth_service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"team-management/common/config"
	"team-management/common/utils"
	"team-management/internal/models/primary/user"
	"team-management/internal/repo/user_repo"
	"time"
)

type AuthService struct {
	userRepo *user_repo.Repo
}

func NewAuthService(userRepo *user_repo.Repo) *AuthService {
	return &AuthService{userRepo: userRepo}
}
func (a *AuthService) UserSignup(ctx echo.Context, req UserSignupRequest) (*UserResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	userObj, err := a.userRepo.CreateUser(ctx, &user_repo.CreateUserRequest{
		Email:     req.Email,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  hashedPassword,
	})
	accessToken, err := a.generateAuthToken(userObj)
	if err != nil {
		return nil, err
	}
	userResponse := &UserResponse{
		Id:          userObj.Id.String(),
		FirstName:   userObj.FirstName,
		LastName:    userObj.LastName,
		Email:       userObj.Email,
		AccessToken: accessToken,
	}
	return userResponse, nil
}
func (a *AuthService) Login(ctx echo.Context, loginRequest *AuthRequest) (*UserResponse, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return nil, err
	}
	if err := utils.CompareHashedPassword(user.Password, loginRequest.Password); err != nil {
		return nil, fmt.Errorf("invalid password")
	}
	accessToken, err := a.generateAuthToken(user)
	if err != nil {
		return nil, err
	}
	authResponse := &UserResponse{
		Id:          user.Id.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		AccessToken: accessToken,
	}
	return authResponse, nil
}
func (a *AuthService) generateAuthToken(userObj *user.User) (string, error) {
	claims := &CustomClaims{
		UserID: userObj.Id.String(),
		Name:   userObj.FirstName,
		Email:  userObj.Email,
		StandardClaims: jwt.StandardClaims{
			Subject:   userObj.Id.String(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetApiConfig().AuthJwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
