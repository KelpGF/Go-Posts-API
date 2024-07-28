package repositories

import (
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/KelpGF/Go-Posts-API/test/database"
	internalMock "github.com/KelpGF/Go-Posts-API/test/database/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CreatePostRepositoryTestSuite struct {
	suite.Suite

	sut *CreatePostRepository
	db  *gorm.DB
}

func (suite *CreatePostRepositoryTestSuite) SetupTest() {
	suite.db = database.Setup()
	suite.sut = NewCreatePostRepository(suite.db)
}

func (suite *CreatePostRepositoryTestSuite) TestCreatePostRepositoryCreate() {
	mockPost := internalMock.NewMockPost()
	mockPost.On("GetId").Return(uuid.New().String())
	mockPost.On("GetTitle").Return("Title")
	mockPost.On("GetBody").Return("Body")
	mockPost.On("GetAuthorName").Return("Author Name")
	mockPost.On("GetPublishedAt").Return(time.Now())

	input := &repositories.CreatePostRepositoryInput{
		Data: mockPost,
	}

	err := suite.sut.Create(input)
	suite.NoError(err)

	var post entities.Post
	suite.db.First(&post, "id = ?", input.Data.GetId())

	suite.Equal(input.Data.GetId(), post.ID)
	suite.Equal(input.Data.GetTitle(), post.Title)

	mockPost.AssertExpectations(suite.T())
}

func (suite *CreatePostRepositoryTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&entities.Post{})
	database.Close(suite.db)
}

func TestSuiteCreatePost(t *testing.T) {
	suite.Run(t, new(CreatePostRepositoryTestSuite))
}
