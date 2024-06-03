package main

import (
	"mikti-depublic/app"
	"mikti-depublic/controller"
	"mikti-depublic/repository"
	"mikti-depublic/service"

	"github.com/labstack/echo/v4"
)

func init() {
	app.LoadEnv()
	app.DBConnection()
}

func main() {
	historyRepo := repository.NewHistoryRepositoryImpl(app.DB)
	historyService := service.NewHistoryServiceImpl(historyRepo)
	historyController := controller.NewHistoryControllerImpl(historyService)

	e := echo.New()

	e.GET("/history", historyController.GetHistory)
	e.GET("/history/:id", historyController.GetHistoryByID)
	e.GET("/history/status/:status", historyController.GetHistoryByStatus)

	e.Logger.Fatal(e.Start(":8080"))
}
