package v1

import (
	"github.com/labstack/echo/v4"
	"team-management/di"
)

func RegisterHandlers(v1 *echo.Group, deps *di.Dependencies) {
	registerPublicHandlers(v1, deps)
}
func registerPublicHandlers(v1 *echo.Group, deps *di.Dependencies) {
	registerAuthHandlers(v1, deps)
}
func registerAuthHandlers(v1 *echo.Group, deps *di.Dependencies) {
	auth := v1.Group("/auth")
	auth.POST("/signup", deps.AuthHandler.RegisterUser)
}
