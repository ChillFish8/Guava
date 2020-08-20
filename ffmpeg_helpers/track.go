package ffmpeg_helpers


type Track struct {
	Id string
	StreamUrl string

	VideoTitle string
	VideoUrl string
	VideoDuration float64
}


func FetchTrack(url string, videoType string) (Track, error) {
	switch videoType {
	case "YOUTUBE":
		return fetchYoutube(url)
	case "SOUNDCLOUD":
		return fetchSoundCloud(url)
	default:
		return Track{}, nil
	}
}

func fetchSoundCloud(url string) (Track, error) {
	return Track{}, nil
}



