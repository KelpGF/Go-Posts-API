package usecases

import (
	"errors"
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	entityId "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	internalMock "github.com/KelpGF/Go-Posts-API/test/database/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EditPostUseCaseTestSuite struct {
	suite.Suite

	sut                        *EditPostUseCase
	sutInput                   *dto.EditPostInput
	findPostByIdRepositoryStub *FindPostByIdRepositoryMock
	editPostRepositoryStub     *EditPostRepositoryMock
}

func (suite *EditPostUseCaseTestSuite) SetupTest() {
	suite.findPostByIdRepositoryStub = &FindPostByIdRepositoryMock{}
	suite.editPostRepositoryStub = &EditPostRepositoryMock{}

	suite.sut = NewEditPostUseCase(suite.findPostByIdRepositoryStub, suite.editPostRepositoryStub)
	suite.sutInput = &dto.EditPostInput{
		ID:          uuid.New().String(),
		Title:       "Title",
		Body:        "Body",
		AuthorName:  "Author",
		PublishedAt: time.Now(),
	}
}

func (suite *EditPostUseCaseTestSuite) TestExecuteWhenInvalidIDReturnsError() {
	input := suite.sutInput
	input.ID = "invalid"

	output := suite.sut.Execute(input)

	suite.NotNil(output)
	suite.Equal("Invalid ID", output.Message)
}

func (suite *EditPostUseCaseTestSuite) TestExecuteWhenNoPostFoundForGivenIDReturnsError() {
	input := suite.sutInput

	suite.findPostByIdRepositoryStub.On("FindById", GetUuid(input.ID)).Return(&internalMock.MockPost{}, errors.New("error"))
	output := suite.sut.Execute(input)

	suite.NotNil(output)
	suite.Equal("Error finding post: error", output.Message)
}

func (suite *EditPostUseCaseTestSuite) TestExecuteWhenPostHasErrorsReturnsError() {
	input := suite.sutInput

	postMock := internalMock.NewMockPost()

	postMock.On("HasErrors").Return(true)
	postMock.On("GetNotificationErrors").Return([]error{errors.New("error")})
	postMock.On("GetNotificationErrorMessage").Return("error")
	postMock.On("SetTitle", input.Title).Return()
	postMock.On("SetBody", input.Body).Return()
	postMock.On("SetAuthorName", input.AuthorName).Return()
	postMock.On("SetPublishedAt", input.PublishedAt).Return()

	suite.findPostByIdRepositoryStub.On("FindById", GetUuid(input.ID)).Return(postMock, nil)

	output := suite.sut.Execute(input)

	suite.NotNil(output)
	suite.Equal("error", output.Message)
	suite.Equal([]error{errors.New("error")}, output.Errors)

	postMock.AssertExpectations(suite.T())
}

func (suite *EditPostUseCaseTestSuite) TestExecuteWhenEditPostRepositoryReturnsError() {
	input := suite.sutInput

	postMock := makePostMock(input)

	suite.findPostByIdRepositoryStub.On("FindById", GetUuid(input.ID)).Return(postMock, nil)

	suite.editPostRepositoryStub.On("Edit", &repositories.EditPostRepositoryInput{
		Data: postMock,
	}).Return(errors.New("error"))

	output := suite.sut.Execute(input)

	suite.NotNil(output)
	suite.Equal("Error editing post: error", output.Message)
}

func (suite *EditPostUseCaseTestSuite) TestExecuteReturnSuccess() {
	input := suite.sutInput

	postMock := makePostMock(input)

	suite.findPostByIdRepositoryStub.On("FindById", GetUuid(input.ID)).Return(postMock, nil)

	suite.editPostRepositoryStub.On("Edit", &repositories.EditPostRepositoryInput{
		Data: postMock,
	}).Return(nil)

	output := suite.sut.Execute(input)

	suite.Nil(output)
}

func TestSuiteEditPost(t *testing.T) {
	suite.Run(t, new(EditPostUseCaseTestSuite))
}

func GetUuid(id string) *entityId.ID {
	ID, _ := entityId.ParseID(id)

	return &ID
}

func makePostMock(input *dto.EditPostInput) *internalMock.MockPost {
	postMock := internalMock.NewMockPost()

	postMock.On("HasErrors").Return(false)
	postMock.On("GetNotificationErrors").Return([]error{})
	postMock.On("GetNotificationErrorMessage").Return("")
	postMock.On("SetTitle", input.Title).Return()
	postMock.On("SetBody", input.Body).Return()
	postMock.On("SetAuthorName", input.AuthorName).Return()
	postMock.On("SetPublishedAt", input.PublishedAt).Return()

	return postMock
}

type FindPostByIdRepositoryMock struct {
	mock.Mock
}

func (m *FindPostByIdRepositoryMock) FindById(id *entityId.ID) (entities.Post, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Post), args.Error(1)
}

type EditPostRepositoryMock struct {
	mock.Mock
}

func (m *EditPostRepositoryMock) Edit(input *repositories.EditPostRepositoryInput) error {
	args := m.Called(input)
	return args.Error(0)
}
