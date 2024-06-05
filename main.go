package main

import (
	"flag"
	"log"
	"mikti-depublic/app"
	"mikti-depublic/common"
	"mikti-depublic/controller"
	"mikti-depublic/db/seeds"
	"mikti-depublic/repository"
	"mikti-depublic/service"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	db := app.DBConnection()

	// Seed database if the seed flag is provided
	seedFlag := flag.Bool("seed", false, "seed database")
	flag.Parse()
	if *seedFlag {
		log.Println("Seeding database")
		seeds.Run()
		log.Println("Seeding database done")
		return
	}

	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(common.LoggingMiddleware)

	// Initialize logger
	common.NewLogger()

	// Initialize repository, service, and controller
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo)
	adminController := controller.NewAdminController(adminService)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/user/register", userController.Register)
	e.POST("/admin/register", adminController.Register)

	// Start server
	common.Logger.LogInfo().Msg(e.Start(":8080").Error())
}
