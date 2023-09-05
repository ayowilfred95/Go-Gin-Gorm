package controllers

import (
	"github.com/ayowilfred95/go-gin-gorm/database"
	"github.com/ayowilfred95/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func BlogCreate(c *gin.Context) {
	// Get data from req body

	// Create a post
	blog := models.Blog{Title: "How to create A bitcoin address using golang",
	 Content: "Golang is one of the best programming language", Author: "Wilfred", Tags: []models.Tag{
		{Name: "Wilfred"},
		{Description: "Golang"},
	 }}
	result := database.DB.Create(&blog) 

	if result.Error != nil {
		c.Status(400)
		return
	}

	//	Return it


	c.JSON(200, gin.H{
		"blog": blog,
	})
}
