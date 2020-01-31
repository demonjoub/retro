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
	if err := h.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	c.Response().WriteHeader(http.StatusCreated)
	return c.JSON(http.StatusCreated, schema.Response{Message: "created success", Data: category})
}

func HandlerUpdateCategory(h db.MenuHandler, c echo.Context) error {
	id := c.Param("id")
	category := schema.Category{}
	if err := h.DB.Find(&category, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	if err := h.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	return c.JSON(http.StatusAccepted, schema.Response{Message: "updated success", Data: category})
}

func HandlerDeleteCategory(h db.MenuHandler, c echo.Context) error {
	id := c.Param("id")
	category := schema.Category{}
	if err := h.DB.Find(&category, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	if err := h.DB.Delete(&category).Error; err != nil {
		return c.JSON(http.StatusBadRequest, schema.Error{Error: err.Error()})
	}
	return c.JSON(http.StatusAccepted, schema.Response{Message: "deleted success"})
}

func HandlerGetCategory(h db.MenuHandler, c echo.Context) error {
	id := c.Param("id")
	category := []schema.Category{}
	if err := h.DB.Find(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, schema.Error{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, schema.Response{Message: "ok", Data: category})
}
