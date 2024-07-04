package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/infinitywarg/go-rest-api/database"
	"net/http"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

func CreateBook(c *gin.Context) {
	var book *database.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not created"})
	}
	c.JSON(http.StatusCreated, gin.H{"book": book})
	return
}

func ReadBook(c *gin.Context) {
	var book *database.Book
	id := c.Param("id")

	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
	}
	c.JSON(http.StatusOK, gin.H{"book": book})
	return
}

func ReadBooks(c *gin.Context) {
	var books []database.Book
	res := database.DB.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
	return
}

func UpdateBook(c *gin.Context) {
	var book database.Book
	id := c.Param("id")
	err := c.ShouldBind(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateBook database.Book
	res := database.DB.Model(&updateBook).Where("id = ?", id).Updates(book)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
}

func DeleteBook(c *gin.Context) {
	var book database.Book
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
	return
}
