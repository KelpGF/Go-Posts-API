package notification

import (
	goErrors "errors"
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/stretchr/testify/suite"
)

type NotificationTestSuite struct {
	suite.Suite

	notification *Notification
}

func (suite *NotificationTestSuite) SetupTest() {
	suite.notification = NewNotification("Test")
}

func (suite *NotificationTestSuite) TestNotification() {
	suite.False(suite.notification.HasErrors())

	suite.notification.AddError(errors.NewIsRequiredError("Title"))
	suite.notification.AddError(goErrors.New("Error"))

	suite.True(suite.notification.HasErrors())
	suite.Len(suite.notification.GetErrors(), 2)
	suite.Equal("Title is required", suite.notification.GetErrors()[0].Error())
	suite.Equal("Test: Title is required, Error", suite.notification.GetErrorsMessage())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(NotificationTestSuite))
}
