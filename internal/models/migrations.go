package models

import (
	"gorm.io/gorm"
	"team-management/internal/models/primary/team"
	"team-management/internal/models/primary/user"
)

func MigrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&team.Team{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&team.TeamMember{}); err != nil {
		return err
	}
	return nil
}
