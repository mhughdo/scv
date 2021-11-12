package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app *application) listLanguagesHandler(c echo.Context) error {
	languages, err := app.models.Languages.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, languages)
}
