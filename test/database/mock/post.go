package mock

import (
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
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
