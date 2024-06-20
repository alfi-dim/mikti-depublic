package controller

import (
	"mikti-depublic/model"
	"mikti-depublic/model/web"
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
	// var admin domain.Admin
	admin := new(web.AdminRegisterRequest)
	if err := ctx.Bind(&admin); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := ctx.Validate(admin); err != nil {
		return err
	}

	newAdmin, err := c.AdminService.Register(*admin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// return ctx.JSON(http.StatusOK, map[string]string{"Message": "You have successfully registered as Admin, now you can log in."})
	return ctx.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "You have successfully registered as Admin", newAdmin))
}
