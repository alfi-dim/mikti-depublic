package main

import (
	"flag"
	"log"
	"mikti-depublic/app"
	"mikti-depublic/common"
	"mikti-depublic/controller"
	"mikti-depublic/db/seeds"
	"mikti-depublic/helper"
	"mikti-depublic/repository"
	"mikti-depublic/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomeValidator struct {
	validator *validator.Validate
}

func (cv *CustomeValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := app.DBConnection()
	userRepo := repository.NewUserRepository(db)
	tokenUseCase := helper.NewTokenUseCase()
	userService := service.NewUserService(userRepo, tokenUseCase)
	controller := controller.NewUserController(userService)

	seedFlag := flag.Bool("seed", false, "seed database")
	flag.Parse()
	if *seedFlag {
		log.Println("Seeding database")
		seeds.Run()
		log.Println("Seeding database done")
		return
	}
	e := echo.New()
	e.Validator = &CustomeValidator{validator: validator.New()}
	e.HTTPErrorHandler = helper.BindAndValidate

	e.POST("/login-user", controller.LoginUser)
	e.POST("/login-admin", controller.LoginAdmin)

	// logger
	common.NewLogger()
	e.Use(common.LoggingMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	common.Logger.LogInfo().Msg(e.Start(":8080").Error())
}
