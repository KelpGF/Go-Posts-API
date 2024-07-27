package usecases

import (
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ListPostsUseCaseTestSuite struct {
	suite.Suite

	sut                     *ListPostsUseCase
	sutInput                *dto.ListPostsInput
	listPostsRepositoryStub *ListPostsRepositoryMock
}

func (suite *ListPostsUseCaseTestSuite) SetupTest() {
	suite.listPostsRepositoryStub = &ListPostsRepositoryMock{}

	suite.sut = NewListPostsUseCase(
		suite.listPostsRepositoryStub,
	)

	suite.sutInput = &dto.ListPostsInput{
		AuthorName:    "AuthorName",
		PublishedSort: "asc",
		Paginate: dto.Paginate{
			Page:  2,
			Limit: 20,
		},
	}
}

func (suite *ListPostsUseCaseTestSuite) TestExecuteReturnListPostsOutput() {
	input := suite.sutInput

	output := makeListPostsOutput()

	suite.listPostsRepositoryStub.On(
		"List",
		input,
	).Return(output)

	result := suite.sut.Execute(input)

	suite.Equal(output, result)
}

func (suite *ListPostsUseCaseTestSuite) TestExecuteReturnListPostsOutputWithDefaultPublishedSort() {
	input := suite.sutInput
	input.PublishedSort = "invalid"

	output := makeListPostsOutput()

	suite.listPostsRepositoryStub.On(
		"List",
		&dto.ListPostsInput{
			AuthorName:    input.AuthorName,
			PublishedSort: "desc",
			Paginate: dto.Paginate{
				Page:  2,
				Limit: 20,
			},
		},
	).Return(output)

	result := suite.sut.Execute(input)

	suite.Equal(output, result)
}

func makeListPostsOutput() *dto.ListPostsOutput {
	return &dto.ListPostsOutput{
		Posts: []dto.Post{
			{
				ID:          "1",
				Title:       "Title",
				Body:        "Body",
				AuthorName:  "AuthorName",
				PublishedAt: "2021-01-01T00:00:00Z",
			},
		},
	}
}

func TestSuiteListPostsUseCase(t *testing.T) {
	suite.Run(t, new(ListPostsUseCaseTestSuite))
}

type ListPostsRepositoryMock struct {
	mock.Mock
}

func (m *ListPostsRepositoryMock) List(input *dto.ListPostsInput) *dto.ListPostsOutput {
	args := m.Called(input)
	return args.Get(0).(*dto.ListPostsOutput)
}
