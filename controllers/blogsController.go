package controllers

import (
	"github.com/ayowilfred95/go-gin-gorm/database"
	"github.com/ayowilfred95/go-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func BlogCreate(c *gin.Context) {
	// Get data from req body
	// we store the data first in a struct
	var blogBody struct {
		Title string
		Content string
		Author string
		Tags     []struct {
			Name        string
			Description string
		}
	}
	c.Bind(&blogBody)
	

	// Create a post
	blog := models.Blog{Title: blogBody.Title,
	 Content: blogBody.Content, Author: blogBody.Author,}

	 // Create Tag instances based on the request body

	 for _, tagData := range blogBody.Tags{
		tag:= models.Tag{
			Name: tagData.Name,
			Description: tagData.Description,
		}
		blog.Tags = append(blog.Tags, tag)
	 }

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

func GetBlogs(c *gin.Context) {
	// Get the blogs
	var blogs [] models.Blog
	database.DB.Find(&blogs)

	// Respond to the blogs
	c.JSON(200, gin.H{
		"blogs": blogs,
	})
}


