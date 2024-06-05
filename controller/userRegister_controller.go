package controller

import (
	"mikti-depublic/model/domain"
	"mikti-depublic/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{UserService: service}
}

func (c *UserController) Register(ctx echo.Context) error {
	var user domain.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := c.UserService.Register(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"Message": "User has successfully registered, can now log in."})
}
