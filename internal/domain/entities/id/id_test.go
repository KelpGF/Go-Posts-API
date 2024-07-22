package entities

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IDEntityTestSuite struct {
	suite.Suite

	id ID
}

func (suite *IDEntityTestSuite) SetupTest() {
	id := NewID()

	suite.id = id
}

func (suite *IDEntityTestSuite) TestIDEntity() {
	suite.NotEmpty(suite.id)
}

func (suite *IDEntityTestSuite) TestParseID() {
	id, err := ParseID(suite.id.String())

	suite.NoError(err)
	suite.Equal(suite.id, id)
}

func TestIDEntitySuite(t *testing.T) {
	suite.Run(t, new(IDEntityTestSuite))
}
