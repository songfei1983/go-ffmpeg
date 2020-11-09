package main

import (
	"log"
	"os/exec"
)

func main() {
	ffmpeg, err := exec.LookPath("ffmpeg")
	log.Println(ffmpeg, err)
}
