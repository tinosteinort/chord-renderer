package main

import (
	"fmt"
	"os"

	"github.com/tinosteinort/chord-renderer/chordloader"
	"github.com/tinosteinort/chord-renderer/imageinfo"
	"github.com/tinosteinort/chord-renderer/types"
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
	fmt.Printf("%dx%d\n", imageInfo.Width, imageInfo.Height)
}

func validateArgsCount(args []string) error {
	if len(args) != 7 {
		return fmt.Errorf("invalid amount of arguments")
	}
	return nil
}

func loadChordAndImageInfo(args []string) (types.Chord, types.ImageInfo, error) {

	chord, err := chordloader.LoadChord(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imageInfo, err := imageinfo.FromArgs(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return chord, imageInfo, nil
}
