package main

import (
	"fmt"
	"os"

	"github.com/tinosteinort/chord-renderer/chordinfo"
	"github.com/tinosteinort/chord-renderer/imageinfo"
)

func main() {
	args := os.Args[1:]
	err := validateArgsCount(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// https://ukuleletricks.com/wp-content/uploads/2018/10/how-to-read-ukulele-chord-diagram.png
	// https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318

	chord, imageInfo, err := loadChordAndImageInfo(args)
	if err != nil {
		fmt.Println("Could not get all needed infos: ", err)
		os.Exit(1)
	}

	fmt.Println(chord.Name)
	fmt.Printf("%s - %dx%d\n", imageInfo.TargetFile, imageInfo.Width, imageInfo.Height)
}

func validateArgsCount(args []string) error {
	if len(args) != 4 {
		return fmt.Errorf("invalid amount of arguments")
	}
	return nil
}

func loadChordAndImageInfo(args []string) (chordinfo.Chord, imageinfo.ImageInfo, error) {

	chord, err := chordinfo.LoadChord(args)
	if err != nil {
		return chordinfo.Chord{}, imageinfo.ImageInfo{}, err
	}

	imageInfo, err := imageinfo.FromArgs(args)
	if err != nil {
		return chordinfo.Chord{}, imageinfo.ImageInfo{}, err
	}

	return chord, imageInfo, nil
}
