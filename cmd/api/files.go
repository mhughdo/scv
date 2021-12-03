package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go/format"
	"net/http"
	"net/url"
	"scv/internal/data"
	sanbox "scv/internal/sandbox"
	"scv/models"
	"strings"
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
	language, err := app.models.Languages.Get(file.LanguageID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, envelope{"message": "the requested resource could not be found"})
		default:
			return err
		}
	}

	result := &struct {
		Hash     string          `json:"hash"`
		Language models.Language `json:"language"`
		Content  string          `json:"content"`
	}{
		Hash:     file.Hash,
		Language: *language,
		Content:  file.Content,
	}

	return c.JSON(http.StatusOK, result)
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

	_, err := app.models.Languages.Get(input.LanguageID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, envelope{"message": "the requested resource could not be found"})
		default:
			return err
		}
	}

	referer := c.Request().Header.Get("Referer")
	refererStrings := strings.Split(referer, "/")
	if refererStrings[3] != "" {
		file, err := app.models.Files.Get(refererStrings[3])
		if err == nil && file.Hash == refererStrings[3] {
			return c.String(http.StatusOK, refererStrings[3])
		}
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

func (app *application) formatFileHandler(c echo.Context) error {
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

	unescapedQuery, err := url.QueryUnescape(input.Content)
	if err != nil {
		return c.JSON(http.StatusOK, envelope{
			"content": "",
			"message": err.Error(),
		})
	}

	formatted, err := format.Source([]byte(unescapedQuery))
	if err != nil {
		return c.JSON(http.StatusOK, envelope{
			"content": "",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, envelope{
		"content": string(formatted),
		"message": "",
	})
}

func (app *application) compileAndRunHandler(c echo.Context) error {
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
	//unescapedQuery, err := url.QueryUnescape(input.Content)
	//if err != nil {
	//	return c.JSON(http.StatusOK, envelope{
	//		"content": "",
	//		"message": err.Error(),
	//	})
	//}

	res, err := sanbox.CompileAndRun(input.Content)
	fmt.Println(err)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error occured")
	}

	return c.JSON(http.StatusCreated, res)
}
