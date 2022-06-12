package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mtrentz/simple_gin_api/database"
	"github.com/mtrentz/simple_gin_api/models"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book
	database.DB.First(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}
