package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/resto/service-upload/schema"
	"github.com/resto/util"
)

func checkError(_error error) (bool, string) {
	errorMessage := ""
	flag := false
	if _error != nil {
		errorMessage = _error.Error()
		flag = true
	}
	return flag, errorMessage
}

func HandlerUploadFiles(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		code := http.StatusBadRequest
		response := util.Response(c, code, err.Error(), schema.ResponsePath{})
		return c.JSON(code, response)
	}
	files := form.File["files"]
	path, err := util.WriteFile(files)
	if err != nil {
		code := http.StatusBadRequest
		response := util.Response(c, code, err.Error(), schema.ResponsePath{})
		return c.JSON(code, response)
	}
	data := schema.ResponsePath{Path: path}
	code := http.StatusCreated
	response := util.Response(c, code, "success", data)
	return c.JSON(code, response)
}
