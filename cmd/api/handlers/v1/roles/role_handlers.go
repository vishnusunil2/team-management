package roles

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"team-management/cmd/api/response"
	"team-management/internal/service/role_service"
)

type RoleHandler struct {
	roleService *role_service.RoleService
}

func NewRoleHandler(roleService *role_service.RoleService) *RoleHandler {
	return &RoleHandler{roleService}
}
func (r *RoleHandler) AddRole(c echo.Context) error {
	var req role_service.AddRoleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      nil,
		})
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Error:      "please provide a valid name",
		})
	}

	roleObj, err := r.roleService.AddRole(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding role",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "role added successfully",
		Data:       roleObj,
		Error:      nil,
	})
}
func (r *RoleHandler) AddPermission(c echo.Context) error {
	var req role_service.AddPermissionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Error:      "please provide a valid name",
		})
	}
	if req.Description == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Error:      "please provide a valid description",
		})
	}
	permissionObj, err := r.roleService.AddPermission(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding permission",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "permission added successfully",
		Data:       permissionObj,
		Error:      nil,
	})
}
func (r *RoleHandler) AddRolePermission(c echo.Context) error {
	roleId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid roleId",
			Error:      err.Error(),
		})
	}
	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid permissionId",
			Error:      err.Error(),
		})
	}
	req := role_service.RolePermissionRequest{
		RoleId:       roleId,
		PermissionId: permissionId,
	}
	permissionObj, err := r.roleService.AddRolePermission(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding permission to role",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "permission successfully added to role",
		Data:       permissionObj,
		Error:      nil,
	})
}
func (r *RoleHandler) RemoveRolePermission(c echo.Context) error {
	roleId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid role_id",
			Error:      err.Error(),
		})
	}
	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid permission_id",
			Error:      err.Error(),
		})
	}
	req := role_service.RolePermissionRequest{
		RoleId:       roleId,
		PermissionId: permissionId,
	}
	if err := r.roleService.RemoveRolePermission(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error removing permission for this role",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "permission removed successfully",
		Data:       nil,
		Error:      nil,
	})
}
func (r *RoleHandler) AddGroup(c echo.Context) error {
	var req role_service.AddGroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Error:      "please provide a valid name",
		})
	}
	if req.Description == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Error:      "please provide a valid description",
		})
	}
	groupObj, err := r.roleService.AddGroup(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding group",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "group added successfully",
		Data:       groupObj,
	})
}
func (r *RoleHandler) AddUserGroup(c echo.Context) error {
	userId := c.Param("userId")

	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid permission_id",
			Error:      err.Error(),
		})
	}
	groupRoleRequest := role_service.UserGroupRequest{
		GroupId: groupId,
		UserId:  userId,
	}
	groupRoleObj, err := r.roleService.AddUserGroup(c, groupRoleRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding group role",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "user added to the group successfully",
		Data:       groupRoleObj,
	})
}
func (r *RoleHandler) RemoveUserGroup(c echo.Context) error {
	userId := c.Param("userId")

	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid permission_id",
			Error:      err.Error(),
		})
	}
	groupRoleRequest := role_service.UserGroupRequest{
		GroupId: groupId,
		UserId:  userId,
	}
	if err := r.roleService.RemoveUserGroup(c, groupRoleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error removing group role",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "user removed from the group successfully",
		Data:       nil,
		Error:      nil,
	})
}
func (r *RoleHandler) AddGroupPermission(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid group_id",
			Error:      err.Error(),
		})
	}
	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid permission_id",
			Error:      err.Error(),
		})
	}
	req := role_service.GroupPermissionRequest{
		GroupId:      groupId,
		PermissionId: permissionId,
	}
	groupPermissionObj, err := r.roleService.AddGroupPermission(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error adding group permission",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "group permission added successfully",
		Data:       groupPermissionObj,
	})
}
func (r *RoleHandler) RemoveGroupPermission(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid group_id",
			Error:      err.Error(),
		})
	}
	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid permission_id",
			Error:      err.Error(),
		})
	}
	req := role_service.GroupPermissionRequest{
		GroupId:      groupId,
		PermissionId: permissionId,
	}
	if err := r.roleService.RemoveGroupPermission(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error removing group permission",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "group permission removed successfully",
		Data:       nil,
		Error:      nil,
	})
}
