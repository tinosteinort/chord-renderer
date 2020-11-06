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
	offset := offset(args)

	width, err := strconv.Atoi(args[4+offset])
	if err != nil {
		return ImageInfo{}, fmt.Errorf("Could not convert %s to width", args[4+offset])
	}
	height, err := strconv.Atoi(args[5+offset])
	if err != nil {
		return ImageInfo{}, fmt.Errorf("Could not convert %s to height", args[5+offset])
	}

	return ImageInfo{
		Width:      width,
		Height:     height,
		TargetFile: args[6+offset],
	}, nil
}

func offset(args []string) int {
	if len(args) == 7 {
		return 0
	}
	return -3
}
