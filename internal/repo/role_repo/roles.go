package role_repo

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"team-management/internal/models/primary/roles"
)

type RoleRepo struct {
	DB *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		DB: db,
	}
}
func (r *RoleRepo) AddRole(ctx echo.Context, req AddRoleRequest) (*roles.Role, error) {
	roleObj := roles.NewRole(req.Name)
	if err := r.getBaseRoleDbQuery(ctx).Create(roleObj).Error; err != nil {
		return nil, err
	}
	return roleObj, nil
}
func (r *RoleRepo) AddPermission(ctx echo.Context, req AddPermissionRequest) (*roles.Permission, error) {
	permissionObj := roles.NewPermission(req.Name, req.Description)
	if err := r.getBasePermissionDbQuery(ctx).Create(permissionObj).Error; err != nil {
		return nil, err
	}
	return permissionObj, nil
}
func (r *RoleRepo) AddRolePermission(ctx echo.Context, req RolePermissionRequest) (*RolePermissionResponse, error) {
	rolePermissionObj := roles.NewRolePermission(req.RoleId, req.PermissionId)
	if err := r.getBaseRolePermissionDbQuery(ctx).Create(rolePermissionObj).Error; err != nil {
		return nil, err
	}
	return &RolePermissionResponse{
		RoleId:       rolePermissionObj.RoleId,
		PermissionId: rolePermissionObj.PermissionId,
	}, nil
}
func (r *RoleRepo) RemoveRolePermission(ctx echo.Context, req RolePermissionRequest) error {
	if err := r.getBaseRolePermissionDbQuery(ctx).
		Where("role_id=? AND permission_id=?", req.RoleId, req.PermissionId).
		Delete(&roles.Permission{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *RoleRepo) GetRolePermissions(ctx echo.Context, roleId int) ([]*roles.RolePermission, error) {
	var res []*roles.RolePermission
	if err := r.getBaseRolePermissionDbQuery(ctx).
		Where("role_id=?", roleId).
		Scan(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func (r *RoleRepo) AddGroup(ctx echo.Context, req AddGroupRequest) (*roles.Group, error) {
	groupObj := roles.NewGroup(req.Name, req.Description)
	if err := r.getBaseGroupDbQuery(ctx).Create(groupObj).Error; err != nil {
		return nil, err
	}
	return groupObj, nil
}
func (r *RoleRepo) AddUserGroup(ctx echo.Context, req UserGroupRequest) (UserGroupResponse, error) {
	groupRoleObj := roles.NewGroupRole(req.GroupId, req.UserId)
	if err := r.getBaseUserGroupDbQuery(ctx).Create(groupRoleObj).Error; err != nil {
		return UserGroupResponse{}, err
	}
	return UserGroupResponse{
		GroupId: req.GroupId,
		UserId:  req.UserId,
	}, nil
}
func (r *RoleRepo) RemoveUserGroup(ctx echo.Context, req UserGroupRequest) error {
	if err := r.getBaseUserGroupDbQuery(ctx).
		Where("user_id=? AND group_id=?", req.UserId, req.GroupId).
		Delete(&roles.UserGroup{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *RoleRepo) AddGroupPermission(ctx echo.Context, req GroupPermissionRequest) (*GroupPermissionResponse, error) {
	groupPermissionObj := roles.NewGroupPermission(req.GroupId, req.PermissionId)
	if err := r.getBaseGroupPermissionDbQuery(ctx).Create(groupPermissionObj).Error; err != nil {
		return nil, err
	}
	return &GroupPermissionResponse{
		GroupId:      groupPermissionObj.GroupId,
		PermissionId: groupPermissionObj.PermissionId,
	}, nil
}
func (r *RoleRepo) RemoveGroupPermission(ctx echo.Context, req GroupPermissionRequest) error {
	if err := r.getBaseGroupPermissionDbQuery(ctx).
		Where("group_id=? AND permission_id=?", req.GroupId, req.PermissionId).
		Delete(&roles.GroupPermission{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *RoleRepo) getBaseRoleDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&roles.Role{}).
		WithContext(ctx.Request().Context())
}
func (r *RoleRepo) getBasePermissionDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&roles.Permission{}).
		WithContext(ctx.Request().Context())
}
func (r *RoleRepo) getBaseRolePermissionDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&roles.RolePermission{}).
		WithContext(ctx.Request().Context())
}
func (r *RoleRepo) getBaseGroupDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&roles.Group{}).
		WithContext(ctx.Request().Context())
}
func (r *RoleRepo) getBaseUserGroupDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&roles.UserGroup{}).
		WithContext(ctx.Request().Context())
}
func (r *RoleRepo) getBaseGroupPermissionDbQuery(ctx echo.Context) *gorm.DB {
	return r.DB.Model(&roles.GroupPermission{}).
		WithContext(ctx.Request().Context())
}

func (r *RoleRepo) CheckGroupPermissions(ctx echo.Context, userID string, permissionId int) bool {
	var exists bool
	checkGroupPermissionQuery := `SELECT EXISTS (SELECT * FROM user_groups AS ug JOIN groups AS g ON g.id = ug.group_id 
    JOIN group_permissions AS gp ON g.id = gp.group_id WHERE gp.permission_id = $1 AND ug.user_id = $2)`
	if err := r.DB.Raw(checkGroupPermissionQuery, permissionId, userID).Scan(&exists).
		Error; err != nil {
		log.Println("error checking permissions", err)
		return false
	}
	return exists
}
