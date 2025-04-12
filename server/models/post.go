package models

import (
	"time"
)

type Post struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type CreatePost struct {
	Id        string    `json:"-"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"-"`
}

type UpdatePost struct {
	Title  *string `json:"title,omitempty"`
	Author *string `json:"author,omitempty"`
}
