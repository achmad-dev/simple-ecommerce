package middleware

import (
	"net/http"

	"github.com/achmad-dev/simple-ecommerce/gateway/internal/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtHelper utils.JwtHelper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "missing or malformed JWT",
				})
			}

			token, err := jwtHelper.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "invalid or expired JWT",
				})
			}

			c.Set("email", token)

			return next(c)
		}
	}
}
