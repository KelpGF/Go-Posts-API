package repositories

import (
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/test/database"
	"github.com/KelpGF/Go-Posts-API/test/database/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ListPostsRepositoryTestSuite struct {
	suite.Suite

	db         *gorm.DB
	repository *ListPostsRepository
}

func (suite *ListPostsRepositoryTestSuite) SetupTest() {
	suite.db = database.Setup()
	suite.repository = NewListPostsRepository(suite.db)
}

func (suite *ListPostsRepositoryTestSuite) TearDownTest() {
	database.Close(suite.db)
}

func (suite *ListPostsRepositoryTestSuite) TestListPostsRepositoryListByAuthorAsc() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	input := &dto.ListPostsInput{
		AuthorName:    "author1",
		PublishedSort: "asc",
		Paginate:      dto.Paginate{Page: 1, Limit: 10},
	}

	output := suite.repository.List(input)

	suite.Len(output.Posts, 2)
	suite.Equal(posts[0].ID, output.Posts[0].ID)
	suite.Equal(posts[1].ID, output.Posts[1].ID)
}

func (suite *ListPostsRepositoryTestSuite) TestListPostsRepositoryListByAuthorDesc() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	input := &dto.ListPostsInput{
		AuthorName:    "author1",
		PublishedSort: "desc",
		Paginate:      dto.Paginate{Page: 1, Limit: 10},
	}

	output := suite.repository.List(input)

	suite.Len(output.Posts, 2)
	suite.Equal(posts[0].ID, output.Posts[1].ID)
	suite.Equal(posts[1].ID, output.Posts[0].ID)
}

func (suite *ListPostsRepositoryTestSuite) TestListPostsRepositoryListByAuthorName() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	input := &dto.ListPostsInput{
		AuthorName:    "author",
		PublishedSort: "desc",
		Paginate:      dto.Paginate{Page: 1, Limit: 10},
	}

	output := suite.repository.List(input)

	suite.Len(output.Posts, 3)
}

func (suite *ListPostsRepositoryTestSuite) TestListPostsRepositoryListByPaginate() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	input := &dto.ListPostsInput{
		AuthorName:    "author",
		PublishedSort: "desc",
		Paginate:      dto.Paginate{Page: 2, Limit: 1},
	}

	output := suite.repository.List(input)

	suite.Len(output.Posts, 1)
	suite.Equal(posts[1].ID, output.Posts[0].ID)
}

func TestSuiteListPosts(t *testing.T) {
	suite.Run(t, new(ListPostsRepositoryTestSuite))
}
