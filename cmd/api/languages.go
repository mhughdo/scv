package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *application) listLanguagesHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{}{})
}