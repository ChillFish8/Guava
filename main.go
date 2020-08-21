package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"./config"
	"./ffmpeg_lib"
	"./ffmpeg_lib/filters"
	"./utils"
)

func main() {
	resp, err := ffmpeg_lib.FetchTrack(
		"https://www.youtube.com/watch?v=l3vbvF8bQfI",
		"YOUTUBE",
	)
	utils.Check(err)
	testFfmpeg(resp.StreamUrl)
}

func testFfmpeg(url string) {
	// Create a New FFMpeg player
	trans := ffmpeg_lib.NewFFmpeg()

	// Set stream file
	_ = trans.SetInputPath(url)

	// Set the output pipe
	reader, err := trans.CreateOutputPipe("mp3")
	utils.Check(err)

	// Create a buffer
	ffmpegBuff := bufio.NewReaderSize(reader, 16384)

	// Start ffmpeg
	done := trans.Run(false)

	file, _ := os.Create("./example/example_out.mp3")

	for {
		audioBuff := make([]int16, config.FrameSize*config.Channels)
		err = binary.Read(ffmpegBuff, binary.LittleEndian, &audioBuff)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			log.Fatalln("error reading from ffmpeg stdout", err)
			return
		}

		newAudioBuff := filters.Volume(audioBuff, 0)

		buf := new(bytes.Buffer)
		err = binary.Write(buf, binary.LittleEndian, newAudioBuff)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			file.Write(buf.Bytes())
			break
		}
		if err != nil {
			log.Fatalln("error reading from ffmpeg stdout", err)
			return
		}

		file.Write(buf.Bytes())
	}

	err = <-done
	utils.Check(err)

	fmt.Println("Encoded File!")
}
