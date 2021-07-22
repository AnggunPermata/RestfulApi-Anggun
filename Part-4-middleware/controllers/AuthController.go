package controllers

import (
	"net/http"
	"project/anggun/lib/database"
	"project/anggun/models"

	"github.com/labstack/echo"
)

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, err := database.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "login success",
		"users":  users,
	})
}
