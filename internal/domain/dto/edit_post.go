package dto

import "time"

type EditPostInput struct {
	ID          string    `json:"-"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	AuthorName  string    `json:"author_name"`
	PublishedAt time.Time `json:"published_at"`
}
