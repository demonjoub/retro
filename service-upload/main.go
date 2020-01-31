package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	setupRoutes()
}

func setupRoutes() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/upload", UploadFile)
	e.Logger.Fatal(e.Start(":1101"))
}
