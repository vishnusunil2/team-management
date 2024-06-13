package role_repo

type AddRoleRequest struct {
	Name string
}
type AddPermissionRequest struct {
	Name        string
	Description string
}
type RolePermissionRequest struct {
	RoleId       int
	PermissionId int
}
type AddGroupRequest struct {
	Name        string
	Description string
}
type UserGroupRequest struct {
	GroupId int
	UserId  string
}
type GroupPermissionRequest struct {
	GroupId      int
	PermissionId int
}
type CheckGroupPermission struct {
	UserId       int
	PermissionId int
}
