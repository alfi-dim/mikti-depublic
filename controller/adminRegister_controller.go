package controller

import (
	"mikti-depublic/model/domain"
	"mikti-depublic/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminService *service.AdminService
}

func NewAdminController(service *service.AdminService) *AdminController {
	return &AdminController{AdminService: service}
}

func (c *AdminController) Register(ctx echo.Context) error {
	var admin domain.Admin
	if err := ctx.Bind(&admin); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := c.AdminService.Register(admin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"Message": "Admin has successfully registered, can now log in."})
}
