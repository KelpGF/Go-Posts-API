package usecases

import (
	"errors"
	"testing"
	"time"

	domainErrors "github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/models"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreatePostRepositoryMock struct {
	mock.Mock
}

func (m *CreatePostRepositoryMock) Create(input *repositories.CreatePostRepositoryInput) error {
	args := m.Called(input)

	return args.Error(0)
}

type CreatePostUseCaseTestSuite struct {
	suite.Suite

	sut                      *CreatePostUseCase
	sutInput                 *CreatePostUseCaseInput
	createPostRepositoryStub *CreatePostRepositoryMock
}

func (suite *CreatePostUseCaseTestSuite) SetupTest() {
	suite.createPostRepositoryStub = &CreatePostRepositoryMock{}
	suite.sut = NewCreatePostUseCase(suite.createPostRepositoryStub)
	suite.sutInput = &CreatePostUseCaseInput{
		Data: &models.CreatePost{
			Title:       "Title",
			Body:        "Body",
			AuthorName:  "AuthorName",
			PublishedAt: time.Now(),
		},
	}
}

func (suite *CreatePostUseCaseTestSuite) TestExecuteReturnErrorWhenPostIsInvalid() {
	input := suite.sutInput
	input.Data.Title = ""
	input.Data.Body = ""

	post, err := suite.sut.Execute(input)

	suite.Nil(post)
	suite.NotNil(err)
	suite.Equal(
		err.Message,
		"Post: Title is required, Body is required",
	)
	suite.Equal(
		err.Errors,
		[]error{
			domainErrors.NewIsRequiredError("Title"),
			domainErrors.NewIsRequiredError("Body"),
		},
	)
}

func (suite *CreatePostUseCaseTestSuite) TestExecuteReturnErrorWhenRepositoryFails() {
	input := suite.sutInput

	suite.createPostRepositoryStub.On(
		"Create",
		&repositories.CreatePostRepositoryInput{Data: input.Data},
	).Return(errors.New("Repository error"))

	post, err := suite.sut.Execute(input)

	suite.Nil(post)
	suite.NotNil(err)
	suite.Equal(
		err.Message,
		"Error creating post: Repository error",
	)
	suite.createPostRepositoryStub.AssertExpectations(suite.T())
}

func (suite *CreatePostUseCaseTestSuite) TestExecuteReturnSuccess() {
	input := suite.sutInput

	suite.createPostRepositoryStub.On(
		"Create",
		&repositories.CreatePostRepositoryInput{Data: input.Data},
	).Return(nil)

	post, err := suite.sut.Execute(input)

	suite.NotNil(post)
	suite.Nil(err)
	suite.NotEmpty(post.ID)
	suite.createPostRepositoryStub.AssertExpectations(suite.T())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CreatePostUseCaseTestSuite))
}
