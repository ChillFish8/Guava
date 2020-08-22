package extractors

import (
	"net/url"
	"strings"

	"../../utils"
	"./pornhub"
	"./tiktok"
	"./types"
	"./udn"
	"./universal"
	"./vimeo"
	"./weibo"
	"./xvideos"
	"./yinyuetai"
	"./youku"
	"./youtube"
)

var extractorMap map[string]types.Extractor

func init() {
	youtubeExtractor := youtube.New()

	extractorMap = map[string]types.Extractor{
		"": universal.New(), // universal extractor

		"youku":     youku.New(),
		"youtube":   youtubeExtractor,
		"youtu":     youtubeExtractor, // youtu.be
		"vimeo":     vimeo.New(),
		"weibo":     weibo.New(),
		"yinyuetai": yinyuetai.New(),
		"pornhub":   pornhub.New(),
		"xvideos":   xvideos.New(),
		"udn":       udn.New(),
		"tiktok":    tiktok.New(),
	}
}

// Extract is the main function to extract the data.
func Extract(u string, option types.Options) ([]*types.Data, error) {
	u = strings.TrimSpace(u)
	var domain string

	bilibiliShortLink := utils.MatchOneOf(u, `^(av|BV|ep)\w+`)
	if len(bilibiliShortLink) > 1 {
		bilibiliURL := map[string]string{
			"av": "https://www.bilibili.com/video/",
			"BV": "https://www.bilibili.com/video/",
			"ep": "https://www.bilibili.com/bangumi/play/",
		}
		domain = "bilibili"
		u = bilibiliURL[bilibiliShortLink[1]] + u
	} else {
		u, err := url.ParseRequestURI(u)
		if err != nil {
			return nil, err
		}
		if u.Host == "haokan.baidu.com" {
			domain = "haokan"
		} else {
			domain = utils.Domain(u.Host)
		}
	}
	extractor := extractorMap[domain]
	videos, err := extractor.Extract(u, option)
	if err != nil {
		return nil, err
	}
	for _, v := range videos {
		v.FillUpStreamsData()
	}
	return videos, nil
}
