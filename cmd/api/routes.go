package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *application) routes() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/v1/languages", app.listLanguagesHandler)
	router.GET("/v1/healthcheck", app.healthcheckHandler)
	router.GET("v1/file/:hash", app.getFileHandler)

	return router
}
