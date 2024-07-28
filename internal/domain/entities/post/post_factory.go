package entities

import (
	"time"

	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/notification"
)

type PostFactoryImpl struct{}

func NewPostFactory() PostFactoryImpl {
	return PostFactoryImpl{}
}

func (f PostFactoryImpl) NewPost(title string, body string, authorName string, publishedAt time.Time) (Post, *errors.ErrorModel) {
	post := &post{
		id:           entities.NewID(),
		title:        title,
		body:         body,
		authorName:   authorName,
		publishedAt:  publishedAt,
		createdAt:    time.Now(),
		notification: notification.NewNotification("Post"),
	}

	entityError := post.validate()
	if entityError != nil {
		return nil, entityError
	}

	return post, nil
}

func (f PostFactoryImpl) Restore(
	id string,
	title string,
	body string,
	authorName string,
	publishedAt time.Time,
	createdAt time.Time,
) Post {
	uuId, _ := entities.ParseID(id)

	return &post{
		id:           uuId,
		title:        title,
		body:         body,
		authorName:   authorName,
		publishedAt:  publishedAt,
		createdAt:    createdAt,
		notification: notification.NewNotification("Post"),
	}
}
