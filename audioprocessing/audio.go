package audioprocessing

import (
	"github.com/xfrr/goffmpeg/ffmpeg"
	"github.com/xfrr/goffmpeg/transcoder"

	"../config"
	"../utils"
)

type FFmpegPlayer struct {
	trans *transcoder.Transcoder
	tracks
}
