package integration_test

import (
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/application/usecases"
	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/repositories"
	"github.com/KelpGF/Go-Posts-API/test/database"
	"github.com/KelpGF/Go-Posts-API/test/database/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ListPostsIntegrationTestSuite struct {
	suite.Suite

	db  *gorm.DB
	sut *usecases.ListPostsUseCase
}

func (suite *ListPostsIntegrationTestSuite) SetupTest() {
	db := database.Setup()
	suite.db = db

	suite.sut = usecases.NewListPostsUseCase(
		repositories.NewListPostsRepository(db),
	)
}

func (suite *ListPostsIntegrationTestSuite) TestListPostsUseCaseExecuteWithoutReturnDeletedPosts() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	suite.db.Delete(&posts[0])

	input := &dto.ListPostsInput{}
	output := suite.sut.Execute(input)

	suite.Len(output.Posts, 2)
}

func (suite *ListPostsIntegrationTestSuite) TestListPostsUseCaseExecuteFilterByAuthor() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	input := &dto.ListPostsInput{
		AuthorName:    "author1",
		PublishedSort: "asc",
	}
	output := suite.sut.Execute(input)

	suite.Len(output.Posts, 2)
	suite.Equal(posts[0].ID, output.Posts[0].ID)
	suite.Equal(posts[1].ID, output.Posts[1].ID)
	suite.Equal("author1", output.Posts[0].AuthorName)
	suite.Equal("author1", output.Posts[1].AuthorName)
}

func (suite *ListPostsIntegrationTestSuite) TestListPostsUseCaseExecute() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	input := &dto.ListPostsInput{}
	output := suite.sut.Execute(input)

	suite.Len(output.Posts, 3)
}

func (suite *ListPostsIntegrationTestSuite) TearDownTest() {
	database.Close(suite.db)
}

func TestSuiteListPostsIntegration(t *testing.T) {
	suite.Run(t, new(ListPostsIntegrationTestSuite))
}
