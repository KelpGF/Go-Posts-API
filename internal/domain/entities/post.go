package entities

import "time"

type post struct {
	id          int
	title       string
	body        string
	authorName  string
	publishedAt time.Time
	createdAt   time.Time
}

func NewPost(title string, body string, authorName string, publishedAt time.Time) *post {
	return &post{
		title:       title,
		body:        body,
		authorName:  authorName,
		publishedAt: publishedAt,
	}
}

func (p *post) GetId() int {
	return p.id
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
