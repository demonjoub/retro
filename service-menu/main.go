package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/labstack/echo"
	"github.com/resto/service-menu/db"
	"github.com/resto/service-menu/schema"
)

var h = db.MenuHandler{}

func main() {
	h.Initialize(&schema.Category{})
	h.Initialize(&schema.Menu{})
	h.Initialize(&schema.MenuImage{})
	// hhandler.DB.Model(&schema.Menu{}).Related(&schema.MenuImage{})
	// hhandler.Model(&game).Related(&gameImages)

	setupRoutes()
}

func setupRoutes() {
	e := echo.New()

	e.POST("/category", func(c echo.Context) error {
		return HandlerCreateCategory(h, c)
	})

	e.PUT("/category/:id", func(c echo.Context) error {
		return HandlerUpdateCategory(h, c)
	})

	e.DELETE("/category/:id", func(c echo.Context) error {
		return HandlerDeleteCategory(h, c)
	})

	// get by id
	e.GET("/category/:id", func(c echo.Context) error {
		return HandlerGetCategory(h, c)
	})

	// get all
	e.GET("/category", func(c echo.Context) error {
		return HandlerGetCategory(h, c)
	})

	// create menu
	e.POST("/menu/create", func(c echo.Context) error {
		return HandlerCreateMenu(h, c)
	})

	// Get all menu
	e.GET("/menu", func(c echo.Context) error {
		return HandlerGetMenu(h, c)
	})

	e.Logger.Fatal(e.Start(":1102"))
}
