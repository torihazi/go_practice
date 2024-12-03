package main

import (
	"gin_notes/controllers"
	"gin_notes/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/*")

	models.ConnectDatabase()
	models.DBMigrate()

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes" ,controllers.NotesCreate)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes Application",
		})
	})

	log.Println("Server Started!")
	r.Run() //Defailt Port 8080
}