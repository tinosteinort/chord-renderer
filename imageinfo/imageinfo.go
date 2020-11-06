package imageinfo

import (
	"fmt"
	"strconv"

	"github.com/tinosteinort/chord-renderer/types"
)

func FromArgs(args []string) (types.ImageInfo, error) {
	width, err := strconv.Atoi(args[4])
	if err != nil {
		return types.ImageInfo{}, fmt.Errorf("Could not convert %s to width", args[4])
	}
	height, err := strconv.Atoi(args[5])
	if err != nil {
		return types.ImageInfo{}, fmt.Errorf("Could not convert %s to height", args[5])
	}

	return types.ImageInfo{
		Width:      width,
		Height:     height,
		TargetFile: args[6],
	}, nil
}
