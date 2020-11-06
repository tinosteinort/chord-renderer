package imageinfo

import (
	"fmt"
	"strconv"
)

type ImageInfo struct {
	Width      int
	Height     int
	TargetFile string
}

func FromArgs(args []string) (ImageInfo, error) {
	width, err := strconv.Atoi(args[4])
	if err != nil {
		return ImageInfo{}, fmt.Errorf("Could not convert %s to width", args[4])
	}
	height, err := strconv.Atoi(args[5])
	if err != nil {
		return ImageInfo{}, fmt.Errorf("Could not convert %s to height", args[5])
	}

	return ImageInfo{
		Width:      width,
		Height:     height,
		TargetFile: args[6],
	}, nil
}
