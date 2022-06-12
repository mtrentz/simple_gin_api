package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	AuthorID int `json:"author_id"`
	Author Author `json:"author"`
	PublisherID int `json:"publisher_id"`
	Publisher Publisher `json:"publisher"`
	Genres []Genre `gorm:"many2many:book_genres" json:"genres"` 
}

type Author struct {
	gorm.Model
	Name string `json:"name"`	
}

type Publisher struct {
	gorm.Model
	Name string `json:"name"`	
}

type Genre struct {
	gorm.Model
	Name string `json:"name"`
}