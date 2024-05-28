package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"mikti-depublic/app"
	"mikti-depublic/common"
	"mikti-depublic/db/seeds"
	"net/http"
)

func main() {
	err := godotenv.Load()
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	common.Logger.LogInfo().Msg(e.Start(":8080").Error())
}
