package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infinitywarg/go-rest-api/controllers"
	"github.com/infinitywarg/go-rest-api/database"
)

func main() {
	fmt.Println("starting http service")
	database.InitPostgres()
	r := gin.Default()
	r.GET("/", controllers.GetHealth)
	r.GET("/books/:id", controllers.ReadBook)
	r.GET("/books", controllers.ReadBooks)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	err := r.Run(":5000")
	if err != nil {
		return
	}
}
