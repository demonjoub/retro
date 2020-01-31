package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/resto/service-menu/db"
	"github.com/resto/service-menu/schema"
)

func HandlerCreateCategory(h db.MenuHandler, c echo.Context) error {
	category := schema.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	if (schema.Category{} == category) {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: "It is an empty structure"})
	}
	if err := h.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	c.Response().WriteHeader(http.StatusCreated)
	return c.JSON(http.StatusCreated, category)
}
