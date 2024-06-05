package controller

import (
	"mikti-depublic/model"
	"mikti-depublic/model/web"
	"mikti-depublic/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: service,
	}
}

func (controller *UserControllerImpl) LoginUser(c echo.Context) error {
	user := new(web.UserLoginRequest)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	userRes, errLogin := controller.UserService.LoginUser(user.Email, user.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "login berhasil", userRes))
}

func (controller *UserControllerImpl) LoginAdmin(c echo.Context) error {
	admin := new(web.UserLoginRequest)

	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	adminRes, errLogin := controller.UserService.LoginAdmin(admin.Email, admin.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "login berhasil", adminRes))
}