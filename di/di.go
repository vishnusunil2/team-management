package di

import (
	"gorm.io/gorm"
	"team-management/cmd/api/handlers/v1/auth"
	"team-management/internal/repo/user_repo"
	"team-management/internal/service/user_service"
)

type Dependencies struct {
	AuthHandler *auth.AuthHandler
}

func Initialize(db *gorm.DB) *Dependencies {
	userRepo := user_repo.NewUserRepo(db)
	userService := user_service.NewUserService(userRepo)
	authHandler := auth.NewAuthHandlers(userService)

	return &Dependencies{
		AuthHandler: authHandler,
	}
}
