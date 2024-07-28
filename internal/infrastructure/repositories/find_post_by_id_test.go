package repositories

import (
	"testing"
	"time"

	entities "github.com/KelpGF/Go-Posts-API/internal/domain/entities/id"
	entityPost "github.com/KelpGF/Go-Posts-API/internal/domain/entities/post"
	"github.com/KelpGF/Go-Posts-API/internal/infrastructure/errors"
	"github.com/KelpGF/Go-Posts-API/test/database"
	internalMock "github.com/KelpGF/Go-Posts-API/test/database/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type FindPostByIdRepositoryTestSuite struct {
	suite.Suite

	db              *gorm.DB
	postFactoryStub *RestorePostFactoryStub
	sut             *FindPostByIdRepository
}

func (suite *FindPostByIdRepositoryTestSuite) SetupTest() {
	suite.db = database.Setup()
	suite.postFactoryStub = NewRestorePostFactoryStub()
	suite.sut = NewFindPostByIdRepository(suite.db, suite.postFactoryStub)
}

func (suite *FindPostByIdRepositoryTestSuite) TearDownTest() {
	database.Close(suite.db)
}

func (suite *FindPostByIdRepositoryTestSuite) TestFindPostByIdRepositoryFindByIdReturnError() {
	posts := internalMock.MakePosts()
	internalMock.InsertPosts(suite.db, posts)
	defer internalMock.DeletePosts(suite.db, posts)

	randomId := entities.NewID()
	post, err := suite.sut.FindById(&randomId)

	suite.Equal(err, errors.NewEntityNotFound("Post"))
	suite.Nil(post)
}

func (suite *FindPostByIdRepositoryTestSuite) TestFindPostByIdRepositoryFindById() {
	posts := internalMock.MakePosts()
	internalMock.InsertPosts(suite.db, posts)
	defer internalMock.DeletePosts(suite.db, posts)

	post := posts[1]

	mockPost := internalMock.NewMockPost()
	mockPost.On("GetId").Return(post.ID)
	suite.postFactoryStub.On("Restore", post.ID, post.Title, post.Body, post.AuthorName, post.PublishedAt.Local(), post.CreatedAt.Local()).Return(mockPost)

	postId, _ := entities.ParseID(post.ID)
	foundPost, err := suite.sut.FindById(&postId)

	suite.NoError(err)
	suite.NotNil(foundPost)
	suite.Equal(post.ID, foundPost.GetId())

	suite.postFactoryStub.AssertExpectations(suite.T())
}

func TestSuiteFindPostById(t *testing.T) {
	suite.Run(t, new(FindPostByIdRepositoryTestSuite))
}

type RestorePostFactoryStub struct {
	mock.Mock
}

func NewRestorePostFactoryStub() *RestorePostFactoryStub {
	return &RestorePostFactoryStub{}
}

func (f *RestorePostFactoryStub) Restore(
	id string,
	title string,
	body string,
	authorName string,
	publishedAt time.Time,
	createdAt time.Time,
) entityPost.Post {
	args := f.Called(id, title, body, authorName, publishedAt, createdAt)
	return args.Get(0).(entityPost.Post)
}
