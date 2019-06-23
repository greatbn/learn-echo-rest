package controllers

import (
	"net/http"

	"github.com/greatbn/learn-echo-rest/models"
	"github.com/labstack/echo"
)

// create user

func CreateUser(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return
	}
	if err = c.Validate(user); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, user)
}
