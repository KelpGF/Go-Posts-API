package repositories

import (
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/test/database"
	"github.com/KelpGF/Go-Posts-API/test/database/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type EditPostRepositoryTestSuite struct {
	suite.Suite

	db *gorm.DB
}

func (suite *EditPostRepositoryTestSuite) SetupTest() {
	suite.db = database.Setup()
}

func (suite *EditPostRepositoryTestSuite) TearDownTest() {
	database.Close(suite.db)
}

func (suite *EditPostRepositoryTestSuite) TestEditPostRepositoryEdit() {
	posts := mock.MakePosts()
	mock.InsertPosts(suite.db, posts)
	defer mock.DeletePosts(suite.db, posts)

	post := posts[1]
	mockPost := mock.NewMockPost()
	mockPost.On("GetId").Return(post.ID)
	mockPost.On("GetTitle").Return("New Title")
	mockPost.On("GetBody").Return("New Body")
	mockPost.On("GetAuthorName").Return("New Author Name")
	mockPost.On("GetPublishedAt").Return(post.PublishedAt)

	repository := NewEditPostRepository(suite.db)
	err := repository.Edit(&repositories.EditPostRepositoryInput{
		Data: mockPost,
	})

	suite.Nil(err)
	mockPost.AssertExpectations(suite.T())
}

func TestSuiteEditPost(t *testing.T) {
	suite.Run(t, new(EditPostRepositoryTestSuite))
}
