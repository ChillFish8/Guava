package main

import (
	"fmt"
	"log"

	"github.com/xfrr/goffmpeg/ffmpeg"
	"github.com/xfrr/goffmpeg/transcoder"

	"./ffmpeg_helpers"
)


func main() {
	res, err := ffmpeg_helpers.FetchTrack(
		"https://www.youtube.com/watch?v=kydqgoVUiVs",
		"YOUTUBE",
		)
	if err != nil { log.Fatalln(err) }

	fmt.Println(res)
}


func main2() {

	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)
	// Initialize an empty transcoder
	err := trans.InitializeEmptyTranscoder()
	if err != nil {
		log.Fatal(err)
	}

	conf :=  ffmpeg.Configuration{
		FfmpegBin:   "F:\\Guava\\bin\\ffmpeg.exe",
		FfprobeBin:  "F:\\Guava\\bin\\ffprobe.exe",
	}
	trans.SetConfiguration(conf)
	new := trans.MediaFile()
	//new.SetSeekTime("1:20")
	new.SetAudioFilter("volume=2,atempo=2.0")
	trans.SetMediaFile(new)

	// Create an input pipe to write to, which will return *io.PipeWriter
	_ = trans.SetInputPath("./example/example.mp3")
	_ = trans.SetOutputPath("./example/example_out.mp3")

	//f, _ := os.Create("./example/example_out.mp3")
	//go func() {
	//	defer r.Close()
	//	defer wg.Done()
	//	defer f.Close()
	//
	//	// Read data from output pipe
	//	buff := bytes.Buffer{}
	//	new := buff.Bytes()
	//	n, err := io.ReadFull(r, new)
	//	if err != nil {
	//		fmt.Println("well oof,", n)
	//	}
	//	_, err = f.Write(new)
	//	if err != nil { println("fuck") }
		// Handle error and data...
	//}()

	// Start transcoder process without checking progress
	done := trans.Run(true)

	progress := trans.Output()

	thing :=<- progress
	fmt.Println(thing)
	// This channel is used to wait for the transcoding process to end
	err = <-done
	// Handle error...

	//time.Sleep(4 * time.Second)

}