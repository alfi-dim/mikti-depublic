package controller

import (
	"mikti-depublic/model/domain"
	"mikti-depublic/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserRegisterController struct {
	UserService *service.UserRegisterService
}

func NewUserRegisterController(service *service.UserRegisterService) *UserRegisterController {
	return &UserRegisterController{UserService: service}
}

func (c *UserRegisterController) Register(ctx echo.Context) error {
	var user domain.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := c.UserService.Register(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"Message": "You have successfully registered as User, now you can log in."})
}
