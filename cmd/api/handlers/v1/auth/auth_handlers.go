package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"team-management/cmd/api/response"
	"team-management/internal/service/user_service"
)

type AuthHandler struct {
	userService *user_service.UserService
}

func NewAuthHandlers(userService *user_service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}
func (h *AuthHandler) RegisterUser(c echo.Context) error {
	var req user_service.UserSignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	userObj, err := h.userService.UserSignup(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "there was an error signing up",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "user signed up successfully",
		Data:       userObj,
		Error:      nil,
	})
}
