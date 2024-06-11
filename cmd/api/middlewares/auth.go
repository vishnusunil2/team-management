package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"team-management/cmd/api/response"
	"team-management/common/config"
	"team-management/common/echo_ctx"
	"team-management/internal/service/auth_service"
	"time"
)

const (
	AuthorizationHeader = "Authorization"
)

func BearerTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(AuthorizationHeader)
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "error validating token",
				Data:       nil,
				Error:      "missing authorization header",
			})
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "error validating token",
				Data:       nil,
				Error:      "invalid token type",
			})
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "error validating token",
				Data:       nil,
				Error:      "missing token header",
			})
		}
		claims := &auth_service.CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetApiConfig().AuthJwtSecret), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "error parsing token",
				Data:       nil,
				Error:      err.Error(),
			})
		}
		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "error validating token",
				Data:       nil,
				Error:      "invalid token",
			})
		}
		if claims.ExpiresAt < time.Now().Unix() {
			return c.JSON(http.StatusUnauthorized, response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "session expired",
				Data:       nil,
				Error:      "token expired",
			})
		}
		echo_ctx.SetUserId(c, claims.UserID)
		return next(c)

	}
}
