package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const (
	DbUser = "admin"
	DbPass = "3425"
	DbHost = "127.0.0.1"
	DbPort = 3307
	DbName = "auth"
)

func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPass, DbHost, DbPort, DbName)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
