package main

import (
	"./ffmpeg_helpers"
	"fmt"
	"github.com/xfrr/goffmpeg/ffmpeg"
	"github.com/xfrr/goffmpeg/transcoder"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	resp, err := ffmpeg_helpers.FetchTrack(
		"https://www.youtube.com/watch?v=l3vbvF8bQfI",
		"YOUTUBE",
	)
	if err != nil {
		log.Fatalln(err)
	}

	getTrack(resp.StreamUrl)
}

func getTrack(url string) {
	start := time.Now()
	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)
	// Initialize an empty transcoder
	err := trans.InitializeEmptyTranscoder()
	if err != nil {
		log.Fatal(err)
	}

	conf := ffmpeg.Configuration{
		FfmpegBin:  "F:\\Guava\\bin\\ffmpeg.exe",
		FfprobeBin: "F:\\Guava\\bin\\ffprobe.exe",
	}
	trans.SetConfiguration(conf)
	new := trans.MediaFile()

	new.SetAudioFilter("volume=1")
	trans.SetMediaFile(new)

	_ = trans.SetInputPath(url)
	reader, err := trans.CreateOutputPipe("mp3")
	if err != nil {
		log.Fatalln(err)
	}

	done := trans.Run(false)
	fmt.Println(time.Now().Sub(start))

	trans.MediaFile().SetAudioFilter("volume=2")
	err = handleOutput(reader)
	if err != nil {
		log.Fatalln(err)
	}

	err = <-done
	if err != nil {
		log.Fatalln(err)
	}

}

func handleOutput(read *io.PipeReader) error {
	defer read.Close()

	file, err := os.Create("./example/example_out_2.mp3")
	if err != nil {
		return err
	}

	//buffer := make([]byte, 1024)
	// _, err := io.ReadFull(read, buffer)
	bytearr, err := ioutil.ReadAll(read)
	if err != nil {
		return err
	}

	_, _ = file.Write(bytearr)
	return nil
}
