package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	v1 "team-management/cmd/api/handlers/v1"
	"team-management/common/config"
	"team-management/common/database"
	"team-management/common/logger"
	"team-management/di"
	"team-management/internal/models"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err.Error())
	}
	if cfg == nil {
		logrus.Fatalf("failed to load configurations")
		os.Exit(1)
	}
	logger.Init(cfg.LogLevel)
	db, err := database.Init(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialise the datbase")
		os.Exit(2)
	}
	if err := models.MigrateModels(db); err != nil {
		logrus.Fatalf("failed to migrate models: %v", err.Error())
		os.Exit(3)
	}
	deps := di.Initialize(db)
	e := echo.New()
	apiGroup := e.Group("/api/v1")
	v1.RegisterHandlers(apiGroup, deps)

	e.Logger.Fatal(e.Start(cfg.ServerAddress))
}
