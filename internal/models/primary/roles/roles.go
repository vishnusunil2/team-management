package roles

import "team-management/internal/models/primary/user"

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name string `json:"name" gorm:"unique;not null"`
}

type Permission struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
}
type Group struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `gorm:"not null"`
}
type UserGroup struct {
	UserId  string    `json:"user_id" gorm:"not null"`
	GroupId int       `json:"group_id" gorm:"not null"`
	User    user.User `gorm:"foreignKey:UserId"`
	Group   Group     `gorm:"foreignKey:GroupId"`
}
type RolePermission struct {
	RoleId       int        `gorm:"primaryKey"`
	PermissionId int        `gorm:"primaryKey"`
	Role         Role       `gorm:"foreignKey:RoleId"`
	Permission   Permission `gorm:"foreignKey:PermissionId"`
}
type GroupPermission struct {
	GroupId      int        `gorm:"primaryKey"`
	PermissionId int        `gorm:"primaryKey"`
	Group        Group      `gorm:"foreignKey:GroupId"`
	Permission   Permission `gorm:"foreignKey:PermissionId"`
}

func NewRole(name string) *Role {
	return &Role{Name: name}
}
func NewPermission(name string, description string) *Permission {
	return &Permission{Name: name, Description: description}
}
func NewRolePermission(roleId int, permissionId int) *RolePermission {
	return &RolePermission{RoleId: roleId, PermissionId: permissionId}
}
func NewGroup(name, description string) *Group {
	return &Group{
		Name:        name,
		Description: description,
	}
}
func NewGroupRole(groupId int, userId string) *UserGroup {
	return &UserGroup{GroupId: groupId, UserId: userId}
}
func NewGroupPermission(groupId int, permissionId int) *GroupPermission {
	return &GroupPermission{GroupId: groupId, PermissionId: permissionId}
}
