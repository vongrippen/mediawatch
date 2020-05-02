package grifts

import (
	"math"
	"os"

	"github.com/markbates/grift/grift"
	"github.com/vongrippen/mediawatch/models"
)

var _ = grift.Namespace("plex", func() {

	grift.Desc("cleanup", "Task Description")
	grift.Add("cleanup", func(c *grift.Context) error {
		mediaFilesCount, _ := models.DB.Count(&models.MediaFile{})
		limit := 500
		pages := int(math.Ceil(float64(mediaFilesCount) / float64(limit)))
		currentFiles := []models.MediaFile{}
		toDestroy := []models.MediaFile{}

		for i := 0; i < pages; i++ {
			models.DB.Paginate(i, limit).Select("id, pathname").All(&currentFiles)

			for _, current := range currentFiles {
				_, err := os.Stat(current.Pathname)
				if os.IsNotExist(err) {
					toDestroy = append(toDestroy, current)
				}
			}
		}

		models.DB.Destroy(toDestroy)

		return nil
	})

})
