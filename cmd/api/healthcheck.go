package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)


func (app *application) healthcheckHandler(c echo.Context) error {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
		},
	}

	return c.JSON(http.StatusOK, env)
}
