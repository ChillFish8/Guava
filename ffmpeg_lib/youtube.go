package ffmpeg_lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../extractors/types"
	"../extractors/youtube"
	"../utils"
)

var reg, _ = regexp.Compile("&dur=([0-9]*\\.[0-9]*)&")

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
	utils.Check(err)

	for _, v := range data[0].Streams {
		match := reg.FindAllString(v.Parts[0].URL, 1)

		dur := 1.23
		if len(match) >= 1 {
			sub := strings.ReplaceAll(match[0], "&dur=", "")
			sub = strings.ReplaceAll(sub, "&", "")
			fmt.Println(sub)
			dur, _ = strconv.ParseFloat(sub, 64)
		} else {
			dur = 0.0
		}

		utils.Check(err)

		track := Track{
			Id:            utils.GenerateId(),
			StreamUrl:     v.Parts[0].URL,
			VideoTitle:    data[0].Title,
			VideoUrl:      data[0].URL,
			VideoDuration: dur,
		}
		return track, nil
	}
	return Track{}, nil
}
