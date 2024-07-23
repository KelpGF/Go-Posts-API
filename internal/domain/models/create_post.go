package models

import "time"

type CreatePost struct {
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	AuthorName  string    `json:"author_name"`
	PublishedAt time.Time `json:"published_at"`
}
