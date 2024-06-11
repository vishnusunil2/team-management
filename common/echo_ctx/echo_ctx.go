package echo_ctx

import "github.com/labstack/echo/v4"

const (
	User = "user"
)

func SetUserId(ctx echo.Context, userId string) {
	ctx.Set(User, userId)
}

func GetUserId(ctx echo.Context) string { return ctx.Get(User).(string) }
