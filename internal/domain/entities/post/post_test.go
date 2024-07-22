package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type PostEntityTestSuite struct {
	suite.Suite

	post *post
}

func (suite *PostEntityTestSuite) SetupTest() {
	post, _ := NewPost("Hello World", "This is a test post", "John Doe", time.Now())
	suite.post = post
}

func (suite *PostEntityTestSuite) TestPostEntity() {
	suite.False(suite.post.HasErrors())

	suite.NotEmpty(suite.post.GetId())
	suite.Equal("Hello World", suite.post.GetTitle())
	suite.Equal("This is a test post", suite.post.GetBody())
	suite.Equal("John Doe", suite.post.GetAuthorName())
	suite.Equal(time.Now().Format(time.RFC3339), suite.post.GetPublishedAt().Format(time.RFC3339))
	suite.NotEmpty(suite.post.createdAt)
}

func (suite *PostEntityTestSuite) TestPostEntityValidation() {
	post, entityError := NewPost("", "", "", time.Now())

	suite.Nil(post)
	suite.Equal(
		entityError.Message,
		"Post: Title is required, Body is required, Author's name is required",
	)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PostEntityTestSuite))
}
