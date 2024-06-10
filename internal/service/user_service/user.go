package user_service

import (
	"github.com/labstack/echo/v4"
	"team-management/common/utils"
	"team-management/internal/repo/user_repo"
)

type UserService struct {
	userRepo *user_repo.Repo
}

func NewUserService(repo *user_repo.Repo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}
func (u *UserService) UserSignup(ctx echo.Context, req UserSignupRequest) (*UserResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	userObj, err := u.userRepo.CreateUser(ctx, &user_repo.CreateUserRequest{
		Email:     req.Email,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	userResponse := &UserResponse{
		ID:        userObj.Id,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
		Email:     userObj.Email,
	}
	return userResponse, nil
}
