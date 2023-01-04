package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"training-api/models"

	"github.com/labstack/echo/v4"
)

// this is to test sending data to a page
// template render and data parsing is in main
// api searching for list of users

func FetchAllUsers(c echo.Context) error {
	result, err := models.FetchUsers()

	if err != nil {
		fmt.Println("Checking error(users.controller line14)")

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	tmp, _ := template.ParseFiles("views/template.html")
	tmp.Execute(c.Response(), result.Data)
	return c.JSON(http.StatusOK, result)
}
