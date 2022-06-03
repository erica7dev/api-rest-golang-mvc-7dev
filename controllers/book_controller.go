package controllers

import (
	"strconv"
	"github.com/erica7dev/api-rest-golang/database"
	"github.com/erica7dev/api-rest-golang/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error //return object if has error or not
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book" + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func CreateBook(c *gin.Context){
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error":"cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error":"cannot create book: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func ShowAllBooks(c *gin.Context){
	db := database.GetDatabase()

	var p []models.Book
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error":"cannot list books: " + err.Error(),
		})
		return
	}
	c.JSON(200, p)
}

func UpdateBook(c *gin.Context){
	db := database.GetDatabase()

	var p models.Book

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error":"cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error":"cannot create book: " + err.Error(),
		})
		return
	}
	c.JSON(201, p)
}

func DeleteBook(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book",
		})
		return
	}
	c.Status(204)
}