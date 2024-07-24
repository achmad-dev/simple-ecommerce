package middleware

import (
	"net/http"
	"strings"

	"github.com/achmad-dev/simple-ecommerce/gateway/internal/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtHelper utils.JwtHelper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "missing or malformed JWT",
				})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "missing or malformed JWT",
				})
			}

			email, err := jwtHelper.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "invalid or expired JWT",
				})
			}
			c.Set("email", email)
			return next(c)
		}
	}
}
