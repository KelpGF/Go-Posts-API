package post

import (
	"time"

	"github.com/google/uuid"
)

type MockPost struct {
	ID uuid.UUID
}

func NewMockPost() *MockPost {
	return &MockPost{
		ID: uuid.New(),
	}
}

func (p *MockPost) GetId() string {
	return p.ID.String()
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
	return nil
}

func (p *MockPost) HasErrors() bool {
	return false
}
