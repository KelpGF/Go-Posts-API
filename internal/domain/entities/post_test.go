package entities

import (
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
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

	suite.Equal("Hello World", suite.post.GetTitle())
	suite.Equal("This is a test post", suite.post.GetBody())
	suite.Equal("John Doe", suite.post.GetAuthorName())
	suite.Equal(time.Now().Format(time.RFC3339), suite.post.GetPublishedAt().Format(time.RFC3339))
}

func (suite *PostEntityTestSuite) TestPostEntityValidation() {
	post, _ := NewPost("", "", "", time.Now())

	suite.True(post.HasErrors())
	suite.Equal([]error{
		errors.NewIsRequiredError("Title"),
		errors.NewIsRequiredError("Body"),
		errors.NewIsRequiredError("AuthorName"),
	}, post.notification.GetErrors())

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PostEntityTestSuite))
}
