package repositories

import (
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/KelpGF/Go-Posts-API/test/database"
	"github.com/KelpGF/Go-Posts-API/test/database/post"
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
	input := &repositories.CreatePostRepositoryInput{
		Data: post.NewMockPost(),
	}

	err := suite.sut.Create(input)
	suite.NoError(err)

	var post entities.Post
	suite.db.First(&post, "id = ?", input.Data.GetId())

	suite.Equal(input.Data.GetId(), post.ID)
	suite.Equal(input.Data.GetTitle(), post.Title)
}

func (suite *CreatePostRepositoryTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&entities.Post{})
	database.Close(suite.db)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CreatePostRepositoryTestSuite))
}
