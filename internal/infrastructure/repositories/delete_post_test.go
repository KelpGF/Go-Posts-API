package repositories

import (
	"testing"
	"time"

	"github.com/KelpGF/Go-Posts-API/internal/domain/repositories"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/entities"
	"github.com/KelpGF/Go-Posts-API/test/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type DeletePostRepositoryTestSuite struct {
	suite.Suite

	db  *gorm.DB
	sut *DeletePostRepository
}

func (suite *DeletePostRepositoryTestSuite) SetupTest() {
	suite.db = database.Setup()
	suite.sut = NewDeletePostRepository(suite.db)
}

func (suite *DeletePostRepositoryTestSuite) TestDeletePostRepositoryDelete() {
	input := &repositories.DeletePostRepositoryInput{
		ID: uuid.New().String(),
	}

	suite.db.Create(&entities.Post{
		ID:          input.ID,
		Title:       "title",
		AuthorName:  "author",
		Body:        "body",
		PublishedAt: time.Now(),
		CreatedAt:   time.Now(),
	})

	err := suite.sut.Delete(input)
	suite.NoError(err)

	var post entities.Post
	suite.db.First(&post, "id = ?", input.ID)

	suite.Empty(post)
}

func TestSuiteDeletePost(t *testing.T) {
	suite.Run(t, new(DeletePostRepositoryTestSuite))
}
