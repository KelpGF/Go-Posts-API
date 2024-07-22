package models

import "time"

type CreatePost struct {
	Title       string
	Body        string
	AuthorName  string
	PublishedAt time.Time
	CreatedAt   time.Time
}
