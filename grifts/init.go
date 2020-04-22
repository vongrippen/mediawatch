package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/vongrippen/mediawatch/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
