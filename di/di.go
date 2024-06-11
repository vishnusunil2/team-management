package di

import (
	"gorm.io/gorm"
	"team-management/cmd/api/handlers/v1/auth"
	"team-management/cmd/api/handlers/v1/team"
	"team-management/internal/repo/team_repo"
	"team-management/internal/repo/user_repo"
	"team-management/internal/service/auth_service"
	"team-management/internal/service/team_service"
)

type Dependencies struct {
	AuthHandler *auth.AuthHandler
	TeamHandler *team.TeamHandler
}

func Initialize(db *gorm.DB) *Dependencies {
	userRepo := user_repo.NewUserRepo(db)
	teamRepo := team_repo.NewTeamRepo(db)

	authService := auth_service.NewAuthService(userRepo)
	teamService := team_service.NewTeamService(teamRepo)

	authHandler := auth.NewAuthHandlers(authService)
	teamHandler := team.NewInternalHandler(teamService)

	return &Dependencies{
		AuthHandler: authHandler,
		TeamHandler: teamHandler,
	}
}
