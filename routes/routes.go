package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mtrentz/simple_gin_api/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/books", controllers.GetAllBooks)
	r.GET("/books/:id", controllers.GetBookById)
	r.POST("/books", controllers.CreateBook)
	r.Run()
}