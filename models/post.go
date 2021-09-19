package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
	Author   *User  `json:"author" gorm:"ForeignKey:AuthorID;AssociationForeignKey:ID"`
}
