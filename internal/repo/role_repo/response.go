package role_repo

type RolePermissionResponse struct {
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}
type UserGroupResponse struct {
	GroupId int    `json:"group_id"`
	UserId  string `json:"user_id"`
}
type GroupPermissionResponse struct {
	GroupId      int `json:"group_id"`
	PermissionId int `json:"permission_id"`
}
