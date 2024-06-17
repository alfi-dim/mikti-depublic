package middleware

import (
	"github.com/labstack/echo/v4"
	"mikti-depublic/helper"
	"net/http"
)

func JwtTokenValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Missing or invalid token",
			})
		}

		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Missing or invalid token",
			})
		}
		token := authHeader[len(bearerPrefix):]

		tokenUseCase := helper.NewTokenUseCase()
		claims, err := tokenUseCase.DecodeTokenPayload(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid or expired token",
			})
		}

		// Add claims to context
		c.Set("claims", claims)
		return next(c)
	}
}
