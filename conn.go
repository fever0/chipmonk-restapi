package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitDB() {
	var err error
	conn := "root:@tcp(localhost:3306)/"
	db, err = gorm.Open("mysql", conn)

	//uncomment the log mod for sql queries in console
	// db.LogMode(true)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.Exec("CREATE DATABASE chipmonk_test")
	db.Exec("USE chipmonk_test")

	// Migration to create tables for Users schema
	db.AutoMigrate(&User{})
}
