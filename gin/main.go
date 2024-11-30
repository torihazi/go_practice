package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Hello Gin",
		})
	})

	log.Println("Server Started!")
	r.Run() //Defailt Port 8080
}