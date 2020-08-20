package ffmpeg_helpers


import (
	"github.com/xfrr/goffmpeg/ffmpeg"
	"github.com/xfrr/goffmpeg/transcoder"
	"log"
)

//
//	General structs for configuration
//
type PlayerConfig struct {
	FFmpegBin string
	FFprobeBin string
}


func NewPlayer(config PlayerConfig) Player {
	player := Player{
		trans: newAudio(config),

	}

	return player
}

func newAudio(config PlayerConfig) *transcoder.Transcoder {
	trans := new(transcoder.Transcoder)
	err := trans.InitializeEmptyTranscoder()
	check(err)

	conf :=  ffmpeg.Configuration{
		FfmpegBin:   config.FFmpegBin,
		FfprobeBin:  config.FFprobeBin,
	}
	trans.SetConfiguration(conf)

	return trans
}

func check(err error) {
	if err != nil { log.Fatal(err) }
}

type Player struct {
	trans *transcoder.Transcoder
	tracks map[int][]Track
}



