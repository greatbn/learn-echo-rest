package controllers

import (
	"net/http"

	"github.com/greatbn/learn-echo-rest/models"
	"github.com/labstack/echo"
	"gopkg.in/validator.v2"
)

// create user

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := validator.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusCreated, user, " ")
}
