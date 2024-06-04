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

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {
	err := godotenv.Load()
	db := app.DBConnection()
	eventRepositoryImpl := repository.NewEventRepository(db)
	eventServieImpl := service.NewEventService(eventRepositoryImpl)
	eventControllerImpl := controller.NewEventController(eventServieImpl)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.DBConnection()

	seedFlag := flag.Bool("seed", false, "seed database")
	flag.Parse()
	if *seedFlag {
		log.Println("Seeding database")
		seeds.Run()
		log.Println("Seeding database done")
		return
	}
	e := echo.New()

	// logger
	common.NewLogger()
	e.Use(common.LoggingMiddleware)
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = helper.BindValidate

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/event/createEvent", eventControllerImpl.CreateEvent)
	e.GET("/event/:id", eventControllerImpl.GetEvent)
	e.GET("/event/list", eventControllerImpl.GetListEvent)
	e.PUT("/event/:id", eventControllerImpl.UpdateEvent)
	e.DELETE("/event/:id", eventControllerImpl.DeleteEvent)

	common.Logger.LogInfo().Msg(e.Start(":8080").Error())
}
