package renderer

import (
	"github.com/fogleman/gg"
	"github.com/tinosteinort/chord-renderer/chordinfo"
	"github.com/tinosteinort/chord-renderer/imageinfo"
)

func Render(chord chordinfo.Chord, imageInfo imageinfo.ImageInfo) error {

	dc := gg.NewContext(imageInfo.Width, imageInfo.Height)

	drawImage(dc, chord, imageInfo)

	if err := dc.SavePNG(imageInfo.TargetFile); err != nil {
		return err
	}

	return nil
}

func drawImage(dc *gg.Context, chord chordinfo.Chord, imageInfo imageinfo.ImageInfo) {

	dc.SetRGB255(255, 255, 255)
	dc.Clear()

	dc.SetRGB255(0, 0, 0)
	if err := dc.LoadFontFace("font/Roboto/Roboto-Medium.ttf", 30); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(chord.Name, float64(imageInfo.Width/2), 20, 0.5, 0.5)

	var sideSpace float64 = 20
	var bottomSpace float64 = 20

	var nutLineY float64 = 50
	var nutLineWidth float64 = 10
	dc.SetRGB255(0, 0, 0)
	dc.SetLineWidth(nutLineWidth)
	dc.DrawLine(sideSpace, nutLineY, float64(imageInfo.Width)-sideSpace, nutLineY)
	dc.Stroke()

	var spaceBetweenStrings float64 = (float64(imageInfo.Width) - (2 * sideSpace)) / float64(chord.StringCount-1)
	var stringLineWidth float64 = 3
	for i := 0; i < chord.StringCount; i++ {
		dc.SetRGB255(0, 0, 0)
		dc.SetLineWidth(stringLineWidth)

		var x float64 = float64(i)*spaceBetweenStrings + sideSpace
		dc.DrawLine(x, nutLineY, x, float64(imageInfo.Height)-bottomSpace)
		dc.Stroke()
	}

	var spaceBetweenFrets float64 = (float64(imageInfo.Height) - nutLineY - bottomSpace) / float64(chord.FretCount)
	var fretLineWidth float64 = 3
	for i := 0; i < chord.FretCount; i++ {
		dc.SetRGB255(0, 0, 0)
		dc.SetLineWidth(fretLineWidth)

		var y float64 = float64(i+1) * spaceBetweenFrets
		dc.DrawLine(sideSpace, nutLineY+y, float64(imageInfo.Width)-sideSpace, nutLineY+y)
		dc.Stroke()
	}

	var frettedNoteRadius float64 = 10
	var spaceFromTop = nutLineY
	var heightWithoutSpace float64 = float64(imageInfo.Height) - spaceFromTop - bottomSpace
	var fretHeight float64 = heightWithoutSpace / float64(chord.FretCount)

	for i := 0; i < len(chord.FrettedNotes); i++ {
		dc.SetRGB255(0, 0, 0)

		frettedNode := chord.FrettedNotes[i]
		var x float64 = float64(imageInfo.Width) - sideSpace - (spaceBetweenStrings * float64(frettedNode.StringNumber-1))
		var y float64 = spaceFromTop + (fretHeight / 2) + (fretHeight * float64(frettedNode.FretNumber-1))
		dc.DrawCircle(x, y, frettedNoteRadius)
		dc.Fill()
	}
}
