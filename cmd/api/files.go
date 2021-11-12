package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"scv/internal/data"
)

func (app *application) getFileHandler(c echo.Context) error {
	hash := c.Param("hash")

	file, err := app.models.Files.Get(hash)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, envelope{"error": "the requested resource could not be found"})
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, file)
}
