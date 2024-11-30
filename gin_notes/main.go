package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := "notes:tmp_pwd@tcp(localhost:3306)/notes?charset=utf8mb4"
	// dsn := "root:example@tcp(localhost:3306)/notes?charset=utf8mb4&parseTime=True"
	databse, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error:", err)
		panic("Failed to connect to database!")
	}

	DB = databse
}

func main(){
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/*")

	connectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes Application",
		})
	})

	log.Println("Server Started!")
	r.Run() //Defailt Port 8080
}