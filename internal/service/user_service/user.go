package user_service

import (
	"github.com/labstack/echo/v4"
	"team-management/internal/models/primary/user"
	"team-management/internal/repo/user_repo"
)

type UserService struct {
	userRepo *user_repo.Repo
}

func NewUserService(userRepo *user_repo.Repo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetUserById(ctx echo.Context, id string) (*user.User, error) {
	userObj, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return userObj, nil
}
