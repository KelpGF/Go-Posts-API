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
	suite.post = NewPost("Hello World", "This is a test post", "John Doe", time.Now())
}

func (suite *PostEntityTestSuite) TestPostEntity() {
	suite.Equal("Hello World", suite.post.GetTitle())
	suite.Equal("This is a test post", suite.post.GetBody())
	suite.Equal("John Doe", suite.post.GetAuthorName())
	suite.Equal(time.Now().Format(time.RFC3339), suite.post.GetPublishedAt().Format(time.RFC3339))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PostEntityTestSuite))
}
