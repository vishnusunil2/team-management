package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"team-management/internal/service/role_service"
	"team-management/internal/service/user_service"
)

type PermissionMiddlewares struct {
	userService *user_service.UserService
	roleService *role_service.RoleService
}

func NewPermissionMiddlewares(userService *user_service.UserService, roleService *role_service.RoleService) *PermissionMiddlewares {
	return &PermissionMiddlewares{
		userService: userService,
		roleService: roleService,
	}
}

func (u *PermissionMiddlewares) CheckPermissions(c echo.Context, userId string, permissionsId int) bool {
	userObj, err := u.userService.GetUserById(c, userId)
	if err != nil {
		log.Fatalf("error while rettrieving userId")
		return false
	}
	permissions, err := u.roleService.GetRolePermissions(c, userObj.RoleId)
	if err != nil {
		log.Fatalf("error while rettrieving permissions")
		return false
	}
	for _, permission := range permissions {
		if permission.PermissionId == permissionsId {
			return true
		}
	}
	return false
}
func (u *PermissionMiddlewares) CheckGroupPermissions(c echo.Context, userId string, permissionsId int) bool {
	return false
}
