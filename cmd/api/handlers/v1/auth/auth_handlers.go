package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"team-management/cmd/api/response"
	"team-management/internal/service/auth_service"
)

type AuthHandler struct {
	authService *auth_service.AuthService
}

func NewAuthHandlers(authService *auth_service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}
func (h *AuthHandler) RegisterUser(c echo.Context) error {
	var req auth_service.UserSignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	userObj, err := h.authService.UserSignup(c, req)
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
func (h *AuthHandler) Login(c echo.Context) error {
	var req auth_service.AuthRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	userObj, err := h.authService.Login(c, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error logging in",
			Data:       nil,
			Error:      err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "logged in successfully",
		Data:       userObj,
		Error:      nil,
	})
}
