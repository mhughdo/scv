package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"scv/internal/data"
	"scv/models"
	"time"
)

func (app *application) getFileHandler(c echo.Context) error {
	hash := c.Param("hash")

	file, err := app.models.Files.Get(hash)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, envelope{"message": "the requested resource could not be found"})
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, file)
}

func (app *application) shareFileHandler(c echo.Context) error {
	var input struct {
		LanguageID int    `json:"language_id"`
		Content    string `json:"content"`
	}

	if err := c.Bind(&input); err != nil {
		return err
	}

	if input.LanguageID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Language ID must greater than 0")
	}

	if input.Content == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "File content must be provided")
	}

	file := models.File{
		LanguageID: input.LanguageID,
		Content:    input.Content,
	}

	for {
		hash, err := app.hashids.Encode([]int{int(time.Now().UnixMicro())})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error creating hash for file.")
		}
		file.Hash = hash
		err = app.models.Files.Insert(file)
		if err == nil {
			return c.String(http.StatusOK, hash)
		}
		if !errors.Is(err, data.ErrUniqueViolation) {
			return err
		}
	}
}
