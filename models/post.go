package models

import (
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
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
