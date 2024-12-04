package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "notes:tmp_pwd@tcp(localhost:3306)/notes?charset=utf8mb4&parseTime=true"
	// dsn := "root:example@tcp(localhost:3306)/notes?charset=utf8mb4&parseTime=True"
	databse, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error:", err)
		panic("Failed to connect to database!")
	}

	DB = databse
}

func DBMigrate() {
	DB.AutoMigrate(&Note{})
}