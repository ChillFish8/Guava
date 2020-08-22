package audioprocessing

import "../filters"

type TrackManager struct {
	applyFilters filters.FFMpegFilters

	tracks  []Track
	current Track
	next    Track

	loop        bool
	loopCurrent bool
}

func (t *TrackManager) AddTrack(track Track) int {
	t.tracks = append(t.tracks, track)
	return len(t.tracks) - 1
}

func (t *TrackManager) RemoveTrackByIndex(n int) (bool, Track) {
	if n < len(t.tracks) {
		return false, Track{}
	}
	track := t.tracks[n]
	t.tracks = append(t.tracks[:n], t.tracks[n+1:]...)
	return true, track
}

func (t *TrackManager) RemoveTrackById(id string) (bool, Track) {
	removed := Track{}
	var tempTracks []Track
	for _, v := range t.tracks {
		if v.Id == id {
			removed = v
		} else {
			tempTracks = append(tempTracks, v)
		}
	}
	return removed.Id != "", removed
}

func (t *TrackManager) NextTrack() (bool, Track) {
	if len(t.tracks) == 0 {
		return false, Track{}
	}
	track := t.tracks[0]
	t.tracks = t.tracks[1:]
	return true, track
}

func (t *TrackManager) Clear() {
	t.tracks = *new([]Track)
}

func (t *TrackManager) ListTracks() []Track {
	return t.tracks
}

type Track struct {
	Id        string
	StreamUrl string

	VideoTitle    string
	VideoUrl      string
	VideoDuration float64
}
