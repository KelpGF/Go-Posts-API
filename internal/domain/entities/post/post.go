package entities

import (
	"time"

	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/notification"
)

type Post interface {
	GetId() string
	GetTitle() string
	GetBody() string
	GetAuthorName() string
	GetPublishedAt() time.Time
	GetCreatedAt() time.Time

	GetNotificationErrorMessage() string
	GetNotificationErrors() []error
	HasErrors() bool

	SetTitle(title string)
	SetBody(body string)
	SetAuthorName(authorName string)
	SetPublishedAt(publishedAt time.Time)
}

type post struct {
	id          entities.ID
	title       string
	body        string
	authorName  string
	publishedAt time.Time
	createdAt   time.Time

	notification *notification.Notification
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

func (p *post) GetNotificationErrorMessage() string {
	return p.notification.GetErrorsMessage()
}

func (p *post) HasErrors() bool {
	return p.notification.HasErrors()
}

func (p *post) validate() *errors.ErrorModel {
	p.notification.ClearErrors()

	if p.title == "" {
		err := errors.NewIsRequiredError("Title")
		p.notification.AddError(err)
	}

	if p.body == "" {
		err := errors.NewIsRequiredError("Body")
		p.notification.AddError(err)
	}

	if p.authorName == "" {
		err := errors.NewIsRequiredError("Author's name")
		p.notification.AddError(err)
	}

	if p.HasErrors() {
		return errors.NewErrorModel(p.GetNotificationErrors(), p.GetNotificationErrorMessage())
	}

	return nil
}

func (p *post) SetTitle(title string) {
	p.title = title
	p.validate()
}

func (p *post) SetBody(body string) {
	p.body = body
	p.validate()
}

func (p *post) SetAuthorName(authorName string) {
	p.authorName = authorName
	p.validate()
}

func (p *post) SetPublishedAt(publishedAt time.Time) {
	p.publishedAt = publishedAt
	p.validate()
}
