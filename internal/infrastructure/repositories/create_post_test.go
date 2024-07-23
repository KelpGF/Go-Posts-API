package repositories

import (
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/models"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/KelpGF/Go-Posts-API/test/database"
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
		Data: models.CreatePost{
			Title:       "Hello World",
			Body:        "This is a test post",
			AuthorName:  "KelpGF",
			PublishedAt: time.Now(),
			CreatedAt:   time.Now(),
		},
	}

	err := suite.sut.Create(input)
	suite.NoError(err)
}

func (suite *CreatePostRepositoryTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&entities.Post{})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CreatePostRepositoryTestSuite))
}
