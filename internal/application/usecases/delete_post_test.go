package usecases

import (
	"errors"
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	domainErrors "github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DeletePostUseCaseTestSuite struct {
	suite.Suite

	sut                      *DeletePostUseCase
	sutInput                 *dto.DeletePostInput
	deletePostRepositoryStub *DeletePostRepositoryMock
}

func (suite *DeletePostUseCaseTestSuite) SetupTest() {
	suite.sutInput = &dto.DeletePostInput{
		ID: "ID",
	}
	suite.deletePostRepositoryStub = &DeletePostRepositoryMock{}

	suite.sut = NewDeletePostUseCase(
		suite.deletePostRepositoryStub,
	)
}

func (suite *DeletePostUseCaseTestSuite) TestExecuteReturnErrorWhenDeletePostRepositoryReturnError() {
	input := suite.sutInput

	suite.deletePostRepositoryStub.On(
		"Delete",
		&repositories.DeletePostRepositoryInput{
			ID: input.ID,
		},
	).Return(
		errors.New("Error"),
	)

	err := suite.sut.Execute(input)

	suite.Equal(
		domainErrors.NewErrorModel(nil, "We couldn't delete the post! Check if the ID is correct."),
		err,
	)
}

func (suite *DeletePostUseCaseTestSuite) TestExecuteReturnNilWhenDeletePostRepositoryReturnNil() {
	input := suite.sutInput

	suite.deletePostRepositoryStub.On(
		"Delete",
		&repositories.DeletePostRepositoryInput{
			ID: input.ID,
		},
	).Return(
		nil,
	)

	err := suite.sut.Execute(input)

	suite.Nil(err)
}

func TestSuiteDeletePost(t *testing.T) {
	suite.Run(t, new(DeletePostUseCaseTestSuite))
}

type DeletePostRepositoryMock struct {
	mock.Mock
}

func (m *DeletePostRepositoryMock) Delete(input *repositories.DeletePostRepositoryInput) error {
	args := m.Called(input)
	return args.Error(0)
}
