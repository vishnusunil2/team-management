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
	registerRoleHandlers(protectedGroup, deps)
	registerPermissionHandlers(protectedGroup, deps)
	registerGroupHandlers(protectedGroup, deps)
}
func registerTeamHandlers(v1 *echo.Group, deps *di.Dependencies) {
	team := v1.Group("/teams")
	team.POST("/", deps.TeamHandler.CreateTeam)
	team.POST("/:teamId/members", deps.TeamHandler.AddMember)
	team.DELETE("/:teamId/members/:memberId", deps.TeamHandler.RemoveMember)
	team.PATCH("/:teamId/members/:memberId", deps.TeamHandler.MakeAdmin)
}
func registerRoleHandlers(v1 *echo.Group, deps *di.Dependencies) {
	roles := v1.Group("/roles")
	roles.POST("/", deps.RoleHandler.AddRole)
	roles.POST("/:roleId/permissions/:permissionId", deps.RoleHandler.AddRolePermission)
	roles.DELETE("/:roleId/permissions/:permissionId", deps.RoleHandler.RemoveRolePermission)
}
func registerPermissionHandlers(v1 *echo.Group, deps *di.Dependencies) {
	permissions := v1.Group("/permissions")
	permissions.POST("/", deps.RoleHandler.AddPermission)
}
func registerGroupHandlers(v1 *echo.Group, deps *di.Dependencies) {
	group := v1.Group("/groups")
	group.POST("/", deps.RoleHandler.AddGroup)
	group.POST("/:groupId/users/:userId", deps.RoleHandler.AddGroupRole)
	group.DELETE("/:groupId/users/:userId", deps.RoleHandler.RemoveGroupRole)
	group.POST("/:groupId/permissions/:permissionId", deps.RoleHandler.AddGroupPermission)
	group.DELETE("/:groupId/permissions/:permissionId", deps.RoleHandler.RemoveGroupPermission)
}
