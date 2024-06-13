package role_service

import (
	"github.com/labstack/echo/v4"
	"team-management/internal/models/primary/roles"
	"team-management/internal/repo/role_repo"
)

type RoleService struct {
	roleRepo *role_repo.RoleRepo
}

func NewRoleService(roleRepo *role_repo.RoleRepo) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}
func (r *RoleService) AddRole(ctx echo.Context, req AddRoleRequest) (*roles.Role, error) {
	roleRequest := role_repo.AddRoleRequest{Name: req.Name}
	roleObj, err := r.roleRepo.AddRole(ctx, roleRequest)
	if err != nil {
		return nil, err
	}
	return roleObj, nil
}
func (r *RoleService) AddPermission(ctx echo.Context, req AddPermissionRequest) (*roles.Permission, error) {
	permissionRequest := role_repo.AddPermissionRequest{Name: req.Name, Description: req.Description}
	permissionObj, err := r.roleRepo.AddPermission(ctx, permissionRequest)
	if err != nil {
		return nil, err
	}
	return permissionObj, nil
}
func (r *RoleService) AddRolePermission(ctx echo.Context, req RolePermissionRequest) (*role_repo.RolePermissionResponse, error) {
	rolePermissionRequest := role_repo.RolePermissionRequest{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	}
	rolePermissionObj, err := r.roleRepo.AddRolePermission(ctx, rolePermissionRequest)
	if err != nil {
		return nil, err
	}
	return rolePermissionObj, err
}
func (r *RoleService) RemoveRolePermission(ctx echo.Context, req RolePermissionRequest) error {
	rolePermissionRequest := role_repo.RolePermissionRequest{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	}
	err := r.roleRepo.RemoveRolePermission(ctx, rolePermissionRequest)
	if err != nil {
		return err
	}
	return nil
}
func (r *RoleService) GetRolePermissions(ctx echo.Context, roleId int) ([]*roles.RolePermission, error) {
	return r.roleRepo.GetRolePermissions(ctx, roleId)
}
func (r *RoleService) AddGroup(ctx echo.Context, req AddGroupRequest) (*roles.Group, error) {
	addGroupRequest := role_repo.AddGroupRequest{Name: req.Name, Description: req.Description}
	groupObj, err := r.roleRepo.AddGroup(ctx, addGroupRequest)
	if err != nil {
		return nil, err
	}
	return groupObj, nil
}
func (r *RoleService) AddUserGroup(ctx echo.Context, req UserGroupRequest) (role_repo.UserGroupResponse, error) {
	groupRoleRequest := role_repo.UserGroupRequest{
		GroupId: req.GroupId,
		UserId:  req.UserId,
	}
	groupRoleResponse, err := r.roleRepo.AddUserGroup(ctx, groupRoleRequest)
	if err != nil {
		return role_repo.UserGroupResponse{}, err
	}
	return groupRoleResponse, nil
}
func (r *RoleService) RemoveUserGroup(ctx echo.Context, req UserGroupRequest) error {
	groupRoleRequest := role_repo.UserGroupRequest{
		GroupId: req.GroupId,
		UserId:  req.UserId,
	}
	if err := r.roleRepo.RemoveUserGroup(ctx, groupRoleRequest); err != nil {
		return err
	}
	return nil
}
func (r *RoleService) AddGroupPermission(ctx echo.Context, req GroupPermissionRequest) (*role_repo.GroupPermissionResponse, error) {
	groupPermissionRequest := role_repo.GroupPermissionRequest{
		GroupId:      req.GroupId,
		PermissionId: req.PermissionId,
	}
	groupPermissionObj, err := r.roleRepo.AddGroupPermission(ctx, groupPermissionRequest)
	if err != nil {
		return nil, err
	}
	return groupPermissionObj, nil
}
func (r *RoleService) RemoveGroupPermission(ctx echo.Context, req GroupPermissionRequest) error {
	groupPermissionRequest := role_repo.GroupPermissionRequest{
		GroupId:      req.GroupId,
		PermissionId: req.PermissionId,
	}
	if err := r.roleRepo.RemoveGroupPermission(ctx, groupPermissionRequest); err != nil {
		return err
	}
	return nil
}
func (r *RoleService) CheckGroupPermission(ctx echo.Context, req CheckGroupPermission) bool {
	return r.roleRepo.CheckGroupPermissions(ctx, req.UserId, req.PermissionId)
}
