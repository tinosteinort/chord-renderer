package renderer

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/tinosteinort/chord-renderer/chordinfo"
	"github.com/tinosteinort/chord-renderer/imageinfo"
)

func Render(chord chordinfo.Chord, imageInfo imageinfo.ImageInfo) error {

	image := &chordImage{image.NewRGBA(image.Rect(0, 0, imageInfo.Width, imageInfo.Height))}

	drawImage(image, chord, imageInfo)

	imageFile, err := os.Create(imageInfo.TargetFile)
	if err != nil {
		return err
	}

	png.Encode(imageFile, image.image)

	return nil
}

func drawImage(image *chordImage, chord chordinfo.Chord, imageInfo imageinfo.ImageInfo) {
	bgColor := color.RGBA{100, 200, 200, 255}

	image.Box(0, 0, imageInfo.Width, imageInfo.Height, bgColor)
}

type chordImage struct {
	image *image.RGBA
}

func (chordImage *chordImage) Box(x int, y int, width int, height int, color color.RGBA) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			chordImage.image.Set(x, y, color)
		}
	}
}
