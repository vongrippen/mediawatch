package grifts

import (
	"fmt"

	"github.com/markbates/grift/grift"
	"github.com/vongrippen/mediawatch/lib"
)

var _ = grift.Namespace("plex", func() {

	grift.Desc("scan", "Scan Plex directories")
	grift.Add("scan", func(c *grift.Context) error {
		foo := lib.AnalyzeFile()
		fmt.Println(foo)
		return nil
	})

})
