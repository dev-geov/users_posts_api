package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Posts    []Post `json:"posts" gorm:"foreignKey:AuthorID"`
}
