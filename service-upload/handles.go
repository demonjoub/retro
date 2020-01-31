package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/resto/service-upload/schema"
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

func UploadFile(c echo.Context) error {
	form, err := c.MultipartForm()
	isError, message := checkError(err)
	if isError {
		return c.JSON(http.StatusNoContent, message)
	}
	files := form.File["files"]
	var paths []string
	for _, file := range files {
		src, err := file.Open()
		isError, message = checkError(err)
		if isError {
			return c.JSON(http.StatusNoContent, message)
		}
		defer src.Close()
		// Destination
		now := time.Now() // current local time
		sec := now.Unix()
		data := strconv.FormatInt(sec, 10)
		pathfile := "../images/" + data + file.Filename
		dst, err := os.Create(pathfile)
		isError, message = checkError(err)
		if isError {
			return c.JSON(http.StatusNoContent, message)
		}
		defer dst.Close()
		// copy file
		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusNoContent, err.Error())
		}
		paths = append(paths, pathfile)
	}
	response := &schema.Message{
		Code:    http.StatusCreated,
		Path:    paths,
		Message: "created",
	}
	c.Response().WriteHeader(http.StatusCreated)
	return c.JSONPretty(http.StatusCreated, response, "")
}
