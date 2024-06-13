package di

import (
	"gorm.io/gorm"
	"team-management/cmd/api/handlers/v1/auth"
	"team-management/cmd/api/handlers/v1/roles"
	"team-management/cmd/api/handlers/v1/team"
	"team-management/cmd/api/middlewares"
	"team-management/internal/repo/role_repo"
	"team-management/internal/repo/team_repo"
	"team-management/internal/repo/user_repo"
	"team-management/internal/service/auth_service"
	"team-management/internal/service/role_service"
	"team-management/internal/service/team_service"
	"team-management/internal/service/user_service"
)

type Dependencies struct {
	AuthHandler *auth.AuthHandler
	TeamHandler *team.TeamHandler
	RoleHandler *roles.RoleHandler
}

func Initialize(db *gorm.DB) *Dependencies {
	userRepo := user_repo.NewUserRepo(db)
	teamRepo := team_repo.NewTeamRepo(db)
	roleRepo := role_repo.NewRoleRepo(db)

	authService := auth_service.NewAuthService(userRepo)
	teamService := team_service.NewTeamService(teamRepo)
	roleService := role_service.NewRoleService(roleRepo)
	userService := user_service.NewUserService(userRepo)

	permissionMiddleware := middlewares.NewPermissionMiddlewares(userService, roleService)

	authHandler := auth.NewAuthHandlers(authService)
	teamHandler := team.NewInternalHandler(teamService, permissionMiddleware, roleService)
	roleHandler := roles.NewRoleHandler(roleService)

	return &Dependencies{
		AuthHandler: authHandler,
		TeamHandler: teamHandler,
		RoleHandler: roleHandler,
	}
}
