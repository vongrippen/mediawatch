package grifts

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/nulls"
	"github.com/markbates/grift/grift"
	"github.com/vongrippen/mediawatch/lib"
	"github.com/vongrippen/mediawatch/models"
)

var _ = grift.Namespace("plex", func() {

	grift.Desc("scan", "Scan Plex directories")
	grift.Add("scan", func(c *grift.Context) error {
		watchDirs := lib.GetWatchDirs()

		for _, watchDir := range watchDirs {
			err := filepath.Walk(watchDir.Path, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() != true {
					media := lib.AnalyzeFile(path, info)
					media.ContentType = nulls.NewString(watchDir.ContentType)
					if media.ContentType.String == "tv" {
						pathList := strings.Split(path, string(filepath.Separator))
						pathLength := len(pathList)
						media.Season = nulls.NewString(pathList[pathLength-2])
						media.Show = nulls.NewString(pathList[pathLength-3])
					}

					// Search for existing records on this path and filename
					existing := []models.MediaFile{}
					models.DB.Select("id").Where("pathname = ?", media.Pathname).All(&existing)

					if len(existing) > 0 {
						media.ID = existing[0].ID
					}

					// Save record
					models.DB.Save(&media)
				}
				return nil
			})
			if err != nil {
				log.Println(err)
			}
		}

		return nil
	})
})
