package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/benji93/business-card/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
