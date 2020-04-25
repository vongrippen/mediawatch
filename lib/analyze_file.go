package lib

import (
	"log"
	"os"
	"time"

	"github.com/gobuffalo/nulls"
	ff "github.com/vansante/go-ffprobe"
	"github.com/vongrippen/mediawatch/models"
)

// AnalyzeFile - Analyze a file with ffprobe and return a MediaFile
func AnalyzeFile(path string, file os.FileInfo) models.MediaFile {
	timeout := 60 * time.Second

	data, err := ff.GetProbeData(path, timeout)
	media := models.MediaFile{}
	if err != nil {
		log.Printf("Error getting data (%v): %v", path, err)
	} else {
		log.Printf("Analyzing %v", path)
		videoStream := data.GetFirstVideoStream()
		audioStream := data.GetFirstAudioStream()
		var videoCodec string
		var audioCodec string
		var width int
		var height int
		if videoStream != nil {
			videoCodec = videoStream.CodecName
			width = videoStream.Width
			height = videoStream.Height
		}
		if audioStream != nil {
			audioCodec = audioStream.CodecName
		}
		media = models.MediaFile{
			Pathname:   path,
			Filename:   file.Name(),
			Size:       nulls.NewInt64(file.Size()),
			Format:     nulls.NewString(data.Format.FormatName),
			VideoCodec: nulls.NewString(videoCodec),
			AudioCodec: nulls.NewString(audioCodec),
			Duration:   nulls.NewFloat64(data.Format.DurationSeconds),
			Width:      nulls.NewInt(width),
			Height:     nulls.NewInt(height),
		}
	}

	return media
}
