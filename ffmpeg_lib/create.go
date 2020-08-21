package ffmpeg_lib

import (
	"github.com/xfrr/goffmpeg/ffmpeg"
	"github.com/xfrr/goffmpeg/transcoder"

	"../config"
	"../utils"
)

//
//	General structs for configuration
//
type PlayerConfig struct {
	FFmpegBin  string
	FFprobeBin string
}

func NewFFmpeg() *transcoder.Transcoder {
	trans := new(transcoder.Transcoder)
	err := trans.InitializeEmptyTranscoder()
	utils.Check(err)

	conf := ffmpeg.Configuration{
		FfmpegBin:  config.FFMpegConfig["FFmpegBin"],
		FfprobeBin: config.FFMpegConfig["FFprobeBin"],
	}

	trans.SetConfiguration(conf)
	newMedia := trans.MediaFile()
	newMedia.SetAudioRate(config.AudioRate)
	newMedia.SetAudioChannels(config.Channels)
	newMedia.SetOutputFormat("s16le")
	trans.SetMediaFile(newMedia)
	return trans
}
