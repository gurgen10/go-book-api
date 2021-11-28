package main

import (
	"go/controllers"
	"go/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Main function
func main() {
	models.Connection()

	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.GET("/books", controllers.FindBooks)
	r.GET("/book/:id", controllers.FindBook)
	r.POST("/book", controllers.CreateBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
	r.GET("/people", controllers.FindPeople)

	r.Run(":3000")
}
