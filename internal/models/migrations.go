package models

import (
	"gorm.io/gorm"
	"team-management/internal/models/primary/roles"
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
	if err := db.AutoMigrate(&roles.Role{}, &roles.Permission{}, &roles.RolePermission{}, &roles.Group{}, &roles.UserGroup{}, &roles.GroupPermission{}); err != nil {
		return err
	}
	return nil
}
