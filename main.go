package main

import (
	"github.com/ayowilfred95/go-gin-gorm/api"
	"github.com/ayowilfred95/go-gin-gorm/controllers"
	"github.com/ayowilfred95/go-gin-gorm/database"
	"github.com/gin-gonic/gin"
)

func init() {
	api.LoadEnv()
	database.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/blogs", controllers.BlogCreate)
	r.GET("/blogs", controllers.GetBlogs)
	r.Run()
}
