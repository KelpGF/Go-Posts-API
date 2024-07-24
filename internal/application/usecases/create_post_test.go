package usecases

import (
	"errors"
	"testing"
	"time"

	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	domainErrors "github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/domain/usecases"
	"github.com/KelpGF/Go-Posts-API/test/database/post"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var emptyModelErr *domainErrors.ErrorModel

type CreatePostUseCaseTestSuite struct {
	suite.Suite

	sut                      *CreatePostUseCase
	sutInput                 *usecases.CreatePostUseCaseInput
	createPostRepositoryStub *CreatePostRepositoryMock
	postFactoryStub          *MockPostFactory
}

func (suite *CreatePostUseCaseTestSuite) SetupTest() {
	suite.createPostRepositoryStub = &CreatePostRepositoryMock{}
	suite.postFactoryStub = &MockPostFactory{}

	suite.sut = NewCreatePostUseCase(
		suite.createPostRepositoryStub,
		suite.postFactoryStub,
	)

	suite.sutInput = &usecases.CreatePostUseCaseInput{
		Title:       "Title",
		Body:        "Body",
		AuthorName:  "AuthorName",
		PublishedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func (suite *CreatePostUseCaseTestSuite) TestExecuteReturnErrorWhenPostIsInvalid() {
	input := suite.sutInput
	input.Title = ""
	input.Body = ""

	suite.postFactoryStub.On(
		"NewPost",
		input.Title,
		input.Body,
		input.AuthorName,
		input.PublishedAt,
	).Return(
		post.NewMockPost(),
		&domainErrors.ErrorModel{
			Message: "Post: Title is required, Body is required",
			Errors: []error{
				domainErrors.NewIsRequiredError("Title"),
				domainErrors.NewIsRequiredError("Body"),
			},
		},
	)

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
	mockPost := post.NewMockPost()

	suite.postFactoryStub.On(
		"NewPost",
		input.Title,
		input.Body,
		input.AuthorName,
		input.PublishedAt,
	).Return(
		mockPost,
		emptyModelErr,
	)

	suite.createPostRepositoryStub.On(
		"Create",
		&repositories.CreatePostRepositoryInput{
			Data: mockPost,
		},
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
	mockPost := post.NewMockPost()

	suite.postFactoryStub.On(
		"NewPost",
		input.Title,
		input.Body,
		input.AuthorName,
		input.PublishedAt,
	).Return(
		mockPost,
		emptyModelErr,
	)

	var emptyErr error
	suite.createPostRepositoryStub.On(
		"Create",
		&repositories.CreatePostRepositoryInput{
			Data: mockPost,
		},
	).Return(emptyErr)

	post, err := suite.sut.Execute(input)

	suite.NotNil(post)
	suite.Nil(err)
	suite.NotEmpty(post.ID)
	suite.createPostRepositoryStub.AssertExpectations(suite.T())
}

func TestSuiteCreatePost(t *testing.T) {
	suite.Run(t, new(CreatePostUseCaseTestSuite))
}

type CreatePostRepositoryMock struct {
	mock.Mock
}

func (m *CreatePostRepositoryMock) Create(input *repositories.CreatePostRepositoryInput) error {
	args := m.Called(input)

	return args.Error(0)
}

type MockPostFactory struct {
	mock.Mock
}

func (f *MockPostFactory) NewPost(title, body, authorName string, publishedAt time.Time) (entities.Post, *domainErrors.ErrorModel) {
	args := f.Called(title, body, authorName, publishedAt)

	return args.Get(0).(entities.Post), args.Get(1).(*domainErrors.ErrorModel)
}
