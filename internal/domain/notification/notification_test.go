package notification

import (
	"testing"

	"github.com/KelpGF/Go-Posts-API/internal/domain/errors"
	"github.com/stretchr/testify/suite"
)

type NotificationTestSuite struct {
	suite.Suite

	notification *Notification
}

func (suite *NotificationTestSuite) SetupTest() {
	suite.notification = NewNotification()
}

func (suite *NotificationTestSuite) TestNotification() {
	suite.False(suite.notification.HasErrors())

	suite.notification.AddError(errors.NewIsRequiredError("Title"))
	suite.True(suite.notification.HasErrors())
	suite.Len(suite.notification.GetErrors(), 1)
	suite.Equal("Title is required", suite.notification.GetErrors()[0].Error())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(NotificationTestSuite))
}
