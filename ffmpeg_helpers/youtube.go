package ffmpeg_helpers

import (
	"../extractors/types"
	"../extractors/youtube"
	"../utils"
)

func fetchYoutube(url string) (Track, error) {
	basicOptions := types.Options{
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
	data, err := extractor.Extract(url, basicOptions)
	check(err)

	for _, v := range data[0].Streams {
		track := Track{
			Id:            utils.GenerateId(),
			StreamUrl:     v.Parts[0].URL,
			VideoTitle:    data[0].Title,
			VideoUrl:      data[0].URL,
			VideoDuration: v.Size,
		}
		return track, nil
	}
	return Track{}, nil
}
