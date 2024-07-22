package entities

import (
	"time"

	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/notification"
)

type post struct {
	id          entities.ID
	title       string
	body        string
	authorName  string
	publishedAt time.Time
	createdAt   time.Time

	notification *notification.Notification
}

func NewPost(title string, body string, authorName string, publishedAt time.Time) (*post, error) {
	post := &post{
		id:           entities.NewID(),
		title:        title,
		body:         body,
		authorName:   authorName,
		publishedAt:  publishedAt,
		createdAt:    time.Now(),
		notification: notification.NewNotification(),
	}

	err := post.validate()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *post) GetId() string {
	return p.id.String()
}

func (p *post) GetTitle() string {
	return p.title
}

func (p *post) GetBody() string {
	return p.body
}

func (p *post) GetAuthorName() string {
	return p.authorName
}

func (p *post) GetPublishedAt() time.Time {
	return p.publishedAt
}

func (p *post) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *post) GetNotificationErrors() []error {
	return p.notification.GetErrors()
}

func (p *post) HasErrors() bool {
	return p.notification.HasErrors()
}

func (p *post) validate() error {
	if p.title == "" {
		err := errors.NewIsRequiredError("Title")
		p.notification.AddError(err)
	}

	if p.body == "" {
		err := errors.NewIsRequiredError("Body")
		p.notification.AddError(err)
	}

	if p.authorName == "" {
		err := errors.NewIsRequiredError("AuthorName")
		p.notification.AddError(err)
	}

	return nil
}
