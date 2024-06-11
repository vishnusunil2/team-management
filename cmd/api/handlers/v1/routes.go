package v1

import (
	"github.com/labstack/echo/v4"
	"team-management/cmd/api/middlewares"
	"team-management/di"
)

func RegisterHandlers(v1 *echo.Group, deps *di.Dependencies) {
	registerPublicHandlers(v1, deps)
	registerProtectedHandlers(v1, deps)
}
func registerPublicHandlers(v1 *echo.Group, deps *di.Dependencies) {
	registerAuthHandlers(v1, deps)
}
func registerAuthHandlers(v1 *echo.Group, deps *di.Dependencies) {
	auth := v1.Group("/auth")
	auth.POST("/signup", deps.AuthHandler.RegisterUser)
	auth.POST("/login", deps.AuthHandler.Login)
}
func registerProtectedHandlers(v1 *echo.Group, deps *di.Dependencies) {
	protectedGroup := v1.Group("")
	protectedGroup.Use(middlewares.BearerTokenMiddleware)
	registerTeamHandlers(protectedGroup, deps)

}
func registerTeamHandlers(v1 *echo.Group, deps *di.Dependencies) {
	team := v1.Group("/teams")
	team.POST("/", deps.TeamHandler.CreateTeam)
	team.POST("/:teamId/members", deps.TeamHandler.AddMember)
	team.DELETE("/:teamId/members/:memberId", deps.TeamHandler.RemoveMember)
	team.PATCH("/:teamId/members/:memberId", deps.TeamHandler.MakeAdmin)
}
