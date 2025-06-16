package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // DB is a global variable that holds the database connection

func InitDB() {
	dsn := "aminorsh:gray@2minor4.me@tcp(127.0.0.1:3306)/eztakeout?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	fmt.Println("Database connection established successfully")
	DB = db
}
