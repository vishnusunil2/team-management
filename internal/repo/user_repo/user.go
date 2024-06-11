package user_repo

import (
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
	res := r.getBaseDbQuery(ctx).Create(&userObj)
	if res.Error != nil {
		return nil, res.Error
	}

	return userObj, nil
}
func (r *Repo) GetUserByEmail(ctx echo.Context, email string) (*user.User, error) {
	var userObj user.User
	if err := r.getBaseDbQuery(ctx).Where("email=?", email).First(&userObj).Error; err != nil {
		return nil, err
	}
	return &userObj, nil
}
func (r *Repo) GetUserById(ctx echo.Context, userId string) (*user.User, error) {
	var userObj user.User
	if err := r.getBaseDbQuery(ctx).Where("id=?", userId).First(&userObj).Error; err != nil {
		return nil, err
	}
	return &userObj, nil
}
func (r *Repo) getBaseDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&user.User{}).
		WithContext(ctx.Request().Context())
}
