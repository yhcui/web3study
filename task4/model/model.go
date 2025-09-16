package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
}

type Posts struct {
	gorm.Model
	Title   string
	Content string
	UserId  uint
	User    User
}

type Comments struct {
	gorm.Model
	Content string
	UserId  uint
	User    User
	PostsId uint
	Posts   Posts
}
