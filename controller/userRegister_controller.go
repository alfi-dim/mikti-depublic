package controller

import (
	"mikti-depublic/model"
	"mikti-depublic/model/web"
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
	// var user domain.User
	user := new(web.UserRegisterRequest)
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := ctx.Validate(user); err != nil {
		return err
	}

	newUser, err := c.UserService.Register(*user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "You have successfully registered as User", newUser))
}
