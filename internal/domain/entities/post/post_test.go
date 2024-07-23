package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type PostEntityTestSuite struct {
	suite.Suite

	factory PostFactoryImpl
}

func (suite *PostEntityTestSuite) SetupTest() {
	suite.factory = NewPostFactory()
}

func (suite *PostEntityTestSuite) TestPostEntity() {
	post, _ := suite.factory.NewPost("Hello World", "This is a test post", "John Doe", time.Now())

	suite.False(post.HasErrors())

	suite.NotEmpty(post.GetId())
	suite.Equal("Hello World", post.GetTitle())
	suite.Equal("This is a test post", post.GetBody())
	suite.Equal("John Doe", post.GetAuthorName())
	suite.Equal(time.Now().Format(time.RFC3339), post.GetPublishedAt().Format(time.RFC3339))
	suite.NotEmpty(post.GetCreatedAt())
}

func (suite *PostEntityTestSuite) TestPostEntityValidation() {
	post, entityError := suite.factory.NewPost("", "", "", time.Now())

	suite.Nil(post)
	suite.Equal(
		entityError.Message,
		"Post: Title is required, Body is required, Author's name is required",
	)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(PostEntityTestSuite))
}
