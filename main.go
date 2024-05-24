package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"mikti-depublic/app"
	"mikti-depublic/common"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.DBConnection()
	e := echo.New()

	// logger
	common.NewLogger()
	e.Use(common.LoggingMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	common.Logger.LogInfo().Msg(e.Start(":8080").Error())
}
