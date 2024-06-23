package middleware

import (
	"mikti-depublic/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
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

		// Periksa peran admin
		switch c.Path() {
		case "/event/createEvent":
			// Memeriksa admin untuk createEvent
			id := claims.ID
			if !strings.HasPrefix(id, "ADMIN") {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Akses ditolak, endpoint ini hanya untuk Admin!",
				})
			}
		case "/event/:id":
			// Memeriksa admin untuk update atau delete event
			id := claims.ID
			if !strings.HasPrefix(id, "ADMIN") {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Akses ditolak, endpoint ini hanya untuk Admin!",
				})
			}
		case "/history":
			// Memeriksa admin untuk semua history
			id := claims.ID
			if !strings.HasPrefix(id, "ADMIN") {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Akses ditolak, endpoint ini hanya untuk Admin!",
				})
			}
		case "/history/:id":
			// Memeriksa admin untuk melihat history sesuai ID
			id := claims.ID
			requestedID := c.Param("id")
			if !strings.HasPrefix(id, "ADMIN") && id != requestedID {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Akses ditolak, Anda hanya bisa melihat history Anda sendiri!",
				})
			}
		case "/history/status/:status":
			// Memeriksa admin untuk melihat history sesuai status
			id := claims.ID
			if !strings.HasPrefix(id, "ADMIN") {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Akses ditolak, endpoint ini hanya untuk Admin!",
				})
			}
		}
		return next(c)
	}
}
