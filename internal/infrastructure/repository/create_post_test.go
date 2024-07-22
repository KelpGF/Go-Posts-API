package repository

import (
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/models"
	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CreatePostRepositoryTestSuite struct {
	suite.Suite

	sut *CreatePostRepository
	db  *gorm.DB
}

func (suite *CreatePostRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	}

	db.AutoMigrate(&entities.Post{})

	suite.db = db
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
