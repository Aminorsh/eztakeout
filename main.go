package main

import (
	"eztakeout/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "aminorsh:gray@2minor4.me@tcp(127.0.0.1:3306)/eztakeout?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := router.InitRouter(db)
	r.Run(":8080") // Run on port 8080
}
