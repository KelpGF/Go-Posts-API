package integration_test

import (
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/application/usecases"
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/repositories"
	"github.com/KelpGF/Go-Posts-API/test/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CreatePostIntegrationTestSuite struct {
	suite.Suite

	db  *gorm.DB
	sut *usecases.CreatePostUseCase
}

func (suite *CreatePostIntegrationTestSuite) SetupTest() {
	db := database.Setup()
	suite.db = db

	suite.sut = usecases.NewCreatePostUseCase(
		repositories.NewCreatePostRepository(db),
		entities.NewPostFactory(),
	)
}

func (suite *CreatePostIntegrationTestSuite) TestCreatePostUseCaseExecute() {
	input := &dto.CreatePostInput{
		Title:       "title",
		Body:        "body",
		AuthorName:  "author",
		PublishedAt: time.Now(),
	}

	output, err := suite.sut.Execute(input)

	suite.Nil(err)
	suite.NotEmpty(output.ID)
}

func (suite *CreatePostIntegrationTestSuite) TestCreatePostUseCaseExecuteInvalidInput() {
	input := &dto.CreatePostInput{
		Title:       "",
		Body:        "",
		AuthorName:  "",
		PublishedAt: time.Now(),
	}

	output, err := suite.sut.Execute(input)

	suite.Empty(output)
	suite.Equal(err.Message, "Post: Title is required, Body is required, Author's name is required")
	suite.Equal(err.Errors[0].Error(), "Title is required")
	suite.Equal(err.Errors[1].Error(), "Body is required")
	suite.Equal(err.Errors[2].Error(), "Author's name is required")
}

func (suite *CreatePostIntegrationTestSuite) TearDownTest() {
	database.Close(suite.db)
}

func TestSuiteCreatePostIntegration(t *testing.T) {
	suite.Run(t, new(CreatePostIntegrationTestSuite))
}
