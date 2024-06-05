package main

import (
	"flag"
	"log"
	"mikti-depublic/app"
	"mikti-depublic/controller"
	"mikti-depublic/db/seeds"
	"mikti-depublic/helper"
	"mikti-depublic/repository"
	"mikti-depublic/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	app.LoadEnv()
	app.DBConnection()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	
	db := app.GetDB()
	// user
	userRepo := repository.NewUserRepository(db)
	tokenUseCase := helper.NewTokenUseCase()
	userService := service.NewUserService(userRepo, tokenUseCase)
	userController := controller.NewUserController(userService)
	// event
	eventRepositoryImpl := repository.NewEventRepository(db)
	eventServieImpl := service.NewEventService(eventRepositoryImpl)
	eventControllerImpl := controller.NewEventController(eventServieImpl)
	// history
	historyRepo := repository.NewHistoryRepositoryImpl(app.DB)
	historyService := service.NewHistoryServiceImpl(historyRepo)
	historyController := controller.NewHistoryControllerImpl(historyService)

	seedFlag := flag.Bool("seed", false, "seed database")
	flag.Parse()
	if *seedFlag {
		log.Println("Seeding database")
		seeds.Run()
		log.Println("Seeding database done")
		return
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = helper.BindAndValidate

	// common.NewLogger()
	// e.Use(common.LoggingMiddleware)
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = helper.BindValidate

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// user login
	e.POST("/login-user", userController.LoginUser)
	e.POST("/login-admin", userController.LoginAdmin)

	// event
	e.POST("/event/createEvent", eventControllerImpl.CreateEvent)
	e.GET("/event/:id", eventControllerImpl.GetEvent)
	e.GET("/event/list", eventControllerImpl.GetListEvent)
	e.PUT("/event/:id", eventControllerImpl.UpdateEvent)
	e.DELETE("/event/:id", eventControllerImpl.DeleteEvent)

	// history
	e.GET("/history", historyController.GetHistory)
	e.GET("/history/:id", historyController.GetHistoryByID)
	e.GET("/history/status/:status", historyController.GetHistoryByStatus)

	e.Logger.Fatal(e.Start(":8080"))
}
