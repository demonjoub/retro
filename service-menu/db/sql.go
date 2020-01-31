package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type MenuHandler struct {
	DB *gorm.DB
}

func (h *MenuHandler) Initialize(value interface{}) {
	db, err := gorm.Open("mysql", "admin:1234@tcp(127.0.0.1:3306)/resto?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("connecting fail")
		log.Fatal(err)
	}
	fmt.Println("connecting success")

	db.AutoMigrate(value)
	h.DB = db
}
