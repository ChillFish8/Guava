package config

// FakeHeaders fake http headers
var FakeHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Accept-Charset":  "UTF-8,*;q=0.5",
	"Accept-Encoding": "gzip,deflate,sdch",
	"Accept-Language": "en-US,en;q=0.8",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
}

const (
	Channels  int = 2                   // 1 for mono, 2 for stereo
	AudioRate int = 48000               // audio sampling rate
	FrameSize int = 960                 // uint16 size of each audio frame
	MaxBytes  int = (FrameSize * 2) * 2 // max size of opus data
)

var FFMpegConfig = map[string]string{
	"FFmpegBin":  "./bin/ffmpeg.exe",
	"FFprobeBin": "./bin/ffplay.exe",
}
