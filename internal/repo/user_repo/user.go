package user_repo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"team-management/internal/models/primary/user"
)

var (
	userRepo *Repo
)

type Repo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *Repo {
	return &Repo{
		DB: db,
	}
}
func (r *Repo) CreateUser(ctx echo.Context, request *CreateUserRequest) (*user.User, error) {
	userObj := user.NewUser(request.Email, request.Phone, request.FirstName, request.LastName, request.Password)
	fmt.Println("hii", userObj)
	res := r.getBaseDbQuery(ctx).Create(&userObj)
	if res.Error != nil {
		return nil, res.Error
	}
	return userObj, nil
}
func (r *Repo) getBaseDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&user.User{}).
		WithContext(ctx.Request().Context())
}
