package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"unique"`
	Body  string
}

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
