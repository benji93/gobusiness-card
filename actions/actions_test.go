package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/benji93/business-card/actions"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}
