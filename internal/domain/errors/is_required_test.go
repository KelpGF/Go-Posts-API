package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IsRequiredTestSuite struct {
	suite.Suite

	err IsRequired
}

func (suite *IsRequiredTestSuite) SetupTest() {
	suite.err = NewIsRequiredError("Title")
}

func (suite *IsRequiredTestSuite) TestIsRequiredErrorMessage() {
	errMessage := suite.err.Error()
	fmt.Println(errMessage)
	suite.Equal("Title is required", errMessage)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IsRequiredTestSuite))
}
