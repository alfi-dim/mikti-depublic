package main

import (
	"flag"
	"log"
	"mikti-depublic/app"
	"mikti-depublic/controller"
	"mikti-depublic/repository"
	"mikti-depublic/service"
	"mikti-depublic/common"
	"mikti-depublic/db/seeds"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
)

func init() {
	app.LoadEnv()
	app.DBConnection()
}

func main() {
	seedFlag := flag.Bool("seed", false, "seed database")
	flag.Parse()
	if *seedFlag {
		log.Println("Seeding database")
		seeds.Run()
		log.Println("Seeding database done")
		return
	}

	e := echo.New()

	historyRepo := repository.NewHistoryRepositoryImpl(app.DB)
	historyService := service.NewHistoryServiceImpl(historyRepo)
	historyController := controller.NewHistoryControllerImpl(historyService)

	e.GET("/history", historyController.GetHistory)
	e.GET("/history/:id", historyController.GetHistoryByID)
	e.GET("/history/status/:status", historyController.GetHistoryByStatus)

	e.Logger.Fatal(e.Start(":8080"))
}
