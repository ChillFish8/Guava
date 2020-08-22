package audioprocessing

import (
	"github.com/xfrr/goffmpeg/transcoder"
	"log"
)

type FFmpegPlayer struct {
	PlayerNo int
	Trans    *transcoder.Transcoder
	Tracks   TrackManager
}

func (fp *FFmpegPlayer) Queue(t Track) int {
	index := fp.Tracks.AddTrack(t)
	log.Printf("[ Player %v ] Queued track at position %v", fp.PlayerNo, index)
	return index
}

func (fp *FFmpegPlayer) RemoveTrackByIndex(n int) bool {
	success, track := fp.Tracks.RemoveTrackByIndex(n)
	if success {
		log.Printf("[ Player %v ] Removed track %v", fp.PlayerNo, track.Id)
	} else {
		log.Printf("[ Player %v ] Failed to removed track as index %v", fp.PlayerNo, n)
	}
	return success
}

func (fp *FFmpegPlayer) RemoveTrackById(id string) bool {
	success, track := fp.Tracks.RemoveTrackById(id)
	if success {
		log.Printf("[ Player %v ] Removed track %v", fp.PlayerNo, track.Id)
	} else {
		log.Printf("[ Player %v ] Failed to removed track with id %v", fp.PlayerNo, id)
	}
	return success
}
