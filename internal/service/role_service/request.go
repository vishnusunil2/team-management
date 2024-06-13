package role_service

type AddRoleRequest struct {
	Name string `json:"name" validate:"required"`
}
type AddPermissionRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type RolePermissionRequest struct {
	RoleId       int `json:"role_id" validate:"required"`
	PermissionId int `json:"permission_id" validate:"required"`
}
type AddGroupRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type UserGroupRequest struct {
	GroupId int    `json:"group_id" validate:"required"`
	UserId  string `json:"user_id" validate:"required"`
}
type GroupPermissionRequest struct {
	GroupId      int `json:"group_id" validate:"required"`
	PermissionId int `json:"permission_id" validate:"required"`
}
type CheckGroupPermission struct {
	UserId       string `json:"user_id" validate:"required"`
	PermissionId int    `json:"permission_id" validate:"required"`
}
