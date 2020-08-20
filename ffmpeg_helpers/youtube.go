package ffmpeg_helpers

import (
	"../extractors/types"
	"../extractors/youtube"
	"fmt"
)

func fetchYoutube(url string) (Track, error) {
	x := types.Options{
		Playlist:         false,
		Items:            "",
		ItemStart:        0,
		ItemEnd:          0,
		ThreadNumber:     0,
		Cookie:           "",
		EpisodeTitleOnly: false,
		YoukuCcode:       "",
		YoukuCkey:        "",
		YoukuPassword:    "",
	}
	extractor := youtube.New()
	data, err := extractor.Extract(url, x)
	check(err)

	for k, v := range data[0].Streams {
		fmt.Println(k, v.Quality)
	}
	return Track{}, nil
}
