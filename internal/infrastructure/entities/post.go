package entities

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	ID          string    `gorm:"uuid;primaryKey"`
	Title       string    `gorm:"text;not null"`
	Body        string    `gorm:"text;not null"`
	AuthorName  string    `gorm:"text;not null"`
	PublishedAt time.Time `gorm:"datetime;not null"`
	CreatedAt   time.Time `gorm:"datetime;not null"`
}
