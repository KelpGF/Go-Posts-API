package repositories

import (
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/dto"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/KelpGF/Go-Posts-API/test/database"
	"github.com/google/uuid"
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

func (suite *ListPostsRepositoryTestSuite) TestListPostsRepositoryListByAuthorAsc() {
	posts := makePosts()
	insertPosts(suite.db, posts)
	defer deletePosts(suite.db, posts)

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
	posts := makePosts()
	insertPosts(suite.db, posts)
	defer deletePosts(suite.db, posts)

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
	posts := makePosts()
	insertPosts(suite.db, posts)
	defer deletePosts(suite.db, posts)

	input := &dto.ListPostsInput{
		AuthorName:    "author",
		PublishedSort: "desc",
		Paginate:      dto.Paginate{Page: 1, Limit: 10},
	}

	output := suite.repository.List(input)

	suite.Len(output.Posts, 3)
}

func (suite *ListPostsRepositoryTestSuite) TestListPostsRepositoryListByPaginate() {
	posts := makePosts()
	insertPosts(suite.db, posts)
	defer deletePosts(suite.db, posts)

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

func insertPosts(db *gorm.DB, posts []entities.Post) {
	for _, post := range posts {
		db.Create(&post)
	}
}

func deletePosts(db *gorm.DB, posts []entities.Post) {
	for _, post := range posts {
		db.Delete(&post)
	}
}

func makePosts() []entities.Post {
	return []entities.Post{
		{
			ID:          uuid.New().String(),
			Title:       "title1-a1",
			AuthorName:  "author1",
			Body:        "body",
			PublishedAt: time.Now().Add(-time.Hour * 2),
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Title:       "title2-a1",
			AuthorName:  "author1",
			Body:        "body",
			PublishedAt: time.Now().Add(-time.Hour),
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Title:       "title",
			AuthorName:  "author",
			Body:        "body",
			PublishedAt: time.Now(),
			CreatedAt:   time.Now(),
		},
	}
}
