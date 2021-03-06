package main

import (
	"fmt"
)

const (
	QUALITY_HIGHRES = "highres"
	QUALITY_HD1080 = "hd1080"
	QUALITY_HD720 = "hd720"
	QUALITY_LARGE = "large"
	QUALITY_MEDIUM = "medium"
	QUALITY_SMALL = "small"
	QUALITY_MIN = "min"
	QUALITY_MAX = "max"
	QUALITY_UNKNOWN = "unknown"

	FORMAT_MP4 = "mp4"
	FORMAT_WEBM = "webm"
	FORMAT_FLV = "flv"
	FORMAT_3GP = "3ggp"
	FORMAT_UNKNOWN = "unknown"

	AUDIO_BITRATE_AUTO = 0
	AUDIO_BITRATE_LOW = 64
	AUDIO_BITRATE_MEDIUM = 128
	AUDIO_BITRATE_HIGH = 192

	DEFAULT_DESTINATION = "./%title%.%format%"
	DEFAULT_DESTINATION_MP3 = "./%title%.mp3"
)

func log(format string, params ...interface{}) {
	if cfg.verbose {
		fmt.Printf(format + "\n", params...)
	}
}

func main() {
	videoId, err := cfg.findVideoId()
	if err != nil {
		fmt.Printf("ERROR: unable to detect the video id: %s\n", err)
		return
	}

	response, err := getVideoInfo(videoId)
	if err != nil {
		fmt.Printf("ERROR: unable to request the video information: %s\n", err)
		return
	}

	streams, err := decodeVideoInfo(response)
	if err != nil {
		fmt.Printf("ERROR: unable to decode the server's answer: %s\n", err)
		return
	}

	stream, err := cfg.selectStream(streams)
	if err != nil {
		fmt.Printf("ERROR: unable to select a stream: %s\n", err)
		return
	}

	out, err := getWriter(cfg, stream)
	if err != nil {
		fmt.Printf("ERROR: unable to create the output writer: %s\n", err)
		return
	}
	defer func() {
		err = out.Close()
		if err != nil {
			fmt.Printf("ERROR: unable to close destination: %s\n", err)
		}
	}()

	err = stream.download(out)
	if err != nil {
		fmt.Printf("ERROR: unable to download the stream: %s\n", err)
		return
	}

	fmt.Printf("Done\n")

	return
}

















