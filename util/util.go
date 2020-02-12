package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/resto/schema"
)

func Response(c echo.Context, code int, message string, data interface{}) interface{} {
	c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(code)

	// msg := schema.Body{Message: message}
	body := schema.Body{Message: message, Data: data}
	response := schema.Response{Code: code, Body: body}
	return response
}

func ReadFile() {

}

func WriteFile(files []*multipart.FileHeader) ([]string, error) {
	var path []string
	var err error
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			fmt.Println("-* OPEN FILE Error *-")
			fmt.Println("Open File Error", err.Error())
			fmt.Println("-* --------------- *-")
			return path, err
		}
		defer src.Close()
		// Destination
		now := time.Now()
		sec := now.Unix()
		data := strconv.FormatInt(sec, 10)
		pathFile := "../images/" + data + file.Filename
		dst, err := os.Create(pathFile)
		if err != nil {
			fmt.Println("-* CREACTE FILE Error *-")
			fmt.Println("Create File Error", err.Error())
			fmt.Println("-* --------------- *-")
			return path, err
		}
		defer dst.Close()
		// Copy File
		if _, err := io.Copy(dst, src); err != nil {
			fmt.Println("-* COPY FILE Error *-")
			fmt.Println("Copy File Error", err.Error())
			fmt.Println("-* --------------- *-")
			return path, err
		}
		path = append(path, pathFile)
	}
	return path, err
}
