package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/floostack/transcoder"
	ffmpeg "github.com/floostack/transcoder/ffmpeg"
)

func main() {

	var pathFFM, pathFFP string
	var err error
	if pathFFM, err = exec.LookPath("ffmpeg"); err != nil {
		log.Panic(err)
	}
	if pathFFP, err = exec.LookPath("ffprobe"); err != nil {
		log.Panic(err)
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   pathFFM,
		FfprobeBinPath:  pathFFP,
		ProgressEnabled: true,
	}
	m, err := ffmpeg.New(ffmpegConf).Input("sample.mp4").GetMetadata()
	if err != nil {
		log.Fatal(err)
	}
	print(m)
}

func print(m transcoder.Metadata) {
	s, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(s))
}

func trancode(ffmpegConf *ffmpeg.Config) {
	format := "mp4"
	overwrite := true

	opts := ffmpeg.Options{
		OutputFormat: &format,
		Overwrite:    &overwrite,
	}
	progress, err := ffmpeg.
		New(ffmpegConf).
		Input("sample.mp4").
		Output("out.mp4").
		WithOptions(opts).
		Start(opts)

	if err != nil {
		log.Fatal(err)
	}

	for msg := range progress {
		log.Printf("%+v", msg)
	}
}
