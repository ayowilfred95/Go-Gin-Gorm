package main

import (
	"github.com/ayowilfred95/go-gin-gorm/api"
	"github.com/ayowilfred95/go-gin-gorm/database"
	"github.com/ayowilfred95/go-gin-gorm/models"
)


func init(){
	api.LoadEnv()
	database.ConnectDB()
}
func main(){
	database.DB.AutoMigrate(&models.Blog{})
	database.DB.AutoMigrate(&models.Tag{})
}
