package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"scv/internal/data"
	"strconv"
)

func (app *application) listLanguagesHandler(c echo.Context) error {
	languages, err := app.models.Languages.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, languages)
}

func (app *application) getLanguageHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, envelope{"error": "the requested resource could not be found"})
	}

	language, err := app.models.Languages.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, envelope{"error": "the requested resource could not be found"})
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, language)
}
