package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	LoginUser(c echo.Context) error
	LoginAdmin(c echo.Context) error
}