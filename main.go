package main

import (
	"log"

	"github.com/giorgisio/goav/avformat"
)

func main() {

	filename := "sample.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	// Dump information about file onto standard error.
	ctx.AvDumpFormat(0, filename, 0)

}
