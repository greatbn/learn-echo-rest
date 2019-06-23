package main

import (
	"github.com/greatbn/learn-echo-rest/controllers"
	"github.com/greatbn/learn-echo-rest/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Validator = &models.UserValidator{Validator: validator.New()}
	// Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS

	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		},
	))

	e.POST("/users", controllers.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}