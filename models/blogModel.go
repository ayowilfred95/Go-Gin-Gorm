package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string
	Content string
	Author string
	Tags	[]Tag `gorm:"many2many:blog_tags;"`
}

type Tag struct {
	gorm.Model
	Name string
	Description string
}
