package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/labstack/echo"
	"github.com/resto/service-menu/db"
	"github.com/resto/service-menu/schema"
	"net/http"
)

var mhandler = db.MenuHandler{}

func main() {
	mhandler.Initialize(&schema.Category{})
	setupRoutes()
}

func setupRoutes() {
	e := echo.New()

	e.POST("/category", func(c echo.Context) error {
		return HandlerCreateCategory(mhandler, c)
	})

	e.PUT("/category/:id", func(c echo.Context) error {
		return HandlerUpdateCategory(mhandler, c)
	})

	e.DELETE("/category/:id", func(c echo.Context) error {
		return HandlerDeleteCategory(mhandler, c)
	})

	e.GET("/category/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	e.Logger.Fatal(e.Start(":1102"))
}
