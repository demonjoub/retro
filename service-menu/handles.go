package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/resto/service-menu/db"
	"github.com/resto/service-menu/schema"
	"github.com/resto/util"
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

func HandlerCreateMenu(h db.MenuHandler, c echo.Context) error {
	menu := schema.Menu{}
	if err := c.Bind(&menu); err != nil {
		code := http.StatusBadRequest
		response := util.Response(c, code, err.Error(), schema.Menu{})
		return c.JSON(code, response)
	}
	if err := h.DB.Save(&menu).Error; err != nil {
		code := http.StatusBadRequest
		response := util.Response(c, code, err.Error(), schema.Menu{})
		return c.JSON(code, response)
	}
	// save multipart
	form, err := c.MultipartForm()
	if err != nil {
		code := http.StatusBadRequest
		response := util.Response(c, code, err.Error(), schema.Menu{})
		return c.JSON(code, response)
	}
	files := form.File["files"]
	path, err := util.WriteFile(files)
	if err != nil {
		code := http.StatusBadRequest
		response := util.Response(c, code, err.Error(), schema.Menu{})
		return c.JSON(code, response)
	}
	// save to images db by menu.id
	for _, uri := range path {
		imagesData := schema.MenuImage{MenuId: menu.Id, Image: uri}
		if err := h.DB.Save(&imagesData).Error; err != nil {
			code := http.StatusBadRequest
			response := util.Response(c, code, err.Error(), schema.Menu{})
			return c.JSON(code, response)
		}
	}
	code := http.StatusAccepted
	menu.Path = path
	response := util.Response(c, code, "created", menu)
	return c.JSON(code, response)
}

// query path file in schema Images
func findImagePath(images []schema.MenuImage) []string {
	var path []string
	for _, image := range images {
		path = append(path, image.Image)
	}
	return path
}

func HandlerGetMenu(h db.MenuHandler, c echo.Context) error {
	category := c.QueryParam("type")
	id := c.QueryParam("id")
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")

	menu := []schema.Menu{}
	if len(category) != 0 { // get by category
		if err := h.DB.Where("category_id = ?", category).Find(&menu).Error; err != nil {
			code := http.StatusBadRequest
			response := util.Response(c, code, err.Error(), schema.Menu{})
			return c.JSON(code, response)
		}
	} else if len(id) > 0 { // get by id
		if err := h.DB.Where("id = ?", id).Find(&menu).Error; err != nil {
			code := http.StatusBadRequest
			response := util.Response(c, code, err.Error(), schema.Menu{})
			return c.JSON(code, response)
		}
	} else if len(offset) > 0 && len(limit) > 0 {
		if err := h.DB.Limit(limit).Offset(offset).Find(&menu).Error; err != nil {
			code := http.StatusBadRequest
			response := util.Response(c, code, err.Error(), schema.Menu{})
			return c.JSON(code, response)
		}
	} else { // get all
		if err := h.DB.Find(&menu).Error; err != nil {
			code := http.StatusBadRequest
			response := util.Response(c, code, err.Error(), schema.Menu{})
			return c.JSON(code, response)
		}
	}
	images := []schema.MenuImage{}
	for i, item := range menu {
		var path []string
		if err := h.DB.Select("image").Where("menu_id = ?", item.Id).Find(&images).Error; err != nil {
			code := http.StatusNotFound
			response := util.Response(c, code, err.Error(), schema.Menu{})
			return c.JSON(code, response)
		}
		for _, image := range images {
			path = append(path, image.Image)
		}
		menu[i].Path = path
	}
	code := http.StatusOK
	response := util.Response(c, code, "ok", menu)
	return c.JSON(code, response)
}
