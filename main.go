package main

import (
	"flag"
	"log"
	"mikti-depublic/app"
	"mikti-depublic/controller"
	"mikti-depublic/db/seeds"
	"mikti-depublic/helper"
	"mikti-depublic/middleware"
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
	db := app.GetDB()
	// user register
	userRegisterRepo := repository.NewUserRegisterRepository(db)
	userRegisterService := service.NewUserRegisterService(userRegisterRepo, db)
	userRegisterController := controller.NewUserRegisterController(userRegisterService)

	// admin register
	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo, db)
	adminController := controller.NewAdminController(adminService)

	// user login
	userRepo := repository.NewUserRepository(db)
	tokenUseCase := helper.NewTokenUseCase()
	userService := service.NewUserService(userRepo, tokenUseCase)
	userController := controller.NewUserController(userService)

	// event
	eventRepositoryImpl := repository.NewEventRepository(db)
	eventServiceImpl := service.NewEventService(eventRepositoryImpl)
	eventControllerImpl := controller.NewEventController(eventServiceImpl)
	// history
	historyRepo := repository.NewHistoryRepositoryImpl(app.DB)
	historyService := service.NewHistoryServiceImpl(historyRepo)
	historyController := controller.NewHistoryControllerImpl(historyService)
	// transaction
	transactionRepo := repository.NewTransactionRepository(app.DB)
	transactionService := service.NewTransactionServiceImpl(transactionRepo, eventRepositoryImpl)
	transactionController := controller.NewTransactionControllerImpl(transactionService)

	// Seed database if the seed flag is provided
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

	// e.Use(common.LoggingMiddleware)
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = helper.BindValidate

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// user register
	e.POST("/user/register", userRegisterController.Register)
	e.POST("/admin/register", adminController.Register)

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

	// transaction
	e.POST("/transaction", transactionController.CreateTransaction, middleware.JwtTokenValidator)
	e.GET("/transaction/:id", transactionController.GetTransactionById, middleware.JwtTokenValidator)
	e.GET("/transactions", transactionController.GetAllTransactions, middleware.JwtTokenValidator)
	e.PUT("/transaction/:id", transactionController.ConfirmPayment, middleware.JwtTokenValidator)

	e.Logger.Fatal(e.Start(":8080"))
}
