package mock

import (
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockPost struct {
	mock.Mock
	ID uuid.UUID
}

func NewMockPost() *MockPost {
	return &MockPost{
		ID: uuid.New(),
	}
}

func (p *MockPost) GetId() string {
	args := p.Called()
	return args.String(0)
}

func (p *MockPost) GetTitle() string {
	return "Title"
}

func (p *MockPost) GetBody() string {
	return "Body"
}

func (p *MockPost) GetAuthorName() string {
	return "AuthorName"
}

func (p *MockPost) GetPublishedAt() time.Time {
	return time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
}

func (p *MockPost) GetCreatedAt() time.Time {
	return time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
}

func (p *MockPost) GetNotificationErrors() []error {
	args := p.Called()
	return args.Get(0).([]error)
}

func (p *MockPost) HasErrors() bool {
	args := p.Called()
	return args.Bool(0)
}

func (p *MockPost) GetNotificationErrorMessage() string {
	args := p.Called()
	return args.String(0)
}

func (p *MockPost) SetTitle(title string) {
	p.Called(title)
}

func (p *MockPost) SetBody(body string) {
	p.Called(body)
}

func (p *MockPost) SetAuthorName(authorName string) {
	p.Called(authorName)
}

func (p *MockPost) SetPublishedAt(publishedAt time.Time) {
	p.Called(publishedAt)
}

func InsertPosts(db *gorm.DB, posts []entities.Post) {
	for _, post := range posts {
		db.Create(&post)
	}
}

func DeletePosts(db *gorm.DB, posts []entities.Post) {
	for _, post := range posts {
		db.Delete(&post)
	}
}

func MakePosts() []entities.Post {
	return []entities.Post{
		{
			ID:          uuid.New().String(),
			Title:       "title1-a1",
			AuthorName:  "author1",
			Body:        "body",
			PublishedAt: time.Now().Add(-time.Hour * 2),
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Title:       "title2-a1",
			AuthorName:  "author1",
			Body:        "body",
			PublishedAt: time.Now().Add(-time.Hour),
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Title:       "title",
			AuthorName:  "author",
			Body:        "body",
			PublishedAt: time.Now(),
			CreatedAt:   time.Now(),
		},
	}
}
