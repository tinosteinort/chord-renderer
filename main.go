package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	validateArgsCount(args)

	// https://ukuleletricks.com/wp-content/uploads/2018/10/how-to-read-ukulele-chord-diagram.png

	chord, imageInfo, err := loadChordAndImageInfo(args)
	if err != nil {
		fmt.Println("Could not get infos from parameters: ", err)
		os.Exit(1)
	}

	fmt.Println(chord.name)
	fmt.Printf("%dx%d\n", imageInfo.width, imageInfo.height)
}

func validateArgsCount(args []string) error {
	if len(args) != 7 {
		return fmt.Errorf("invalid amount of arguments")
	}
	return nil
}

func loadChordAndImageInfo(args []string) (*chord, *imageInfo, error) {
	var chord *chord
	var imageInfo *imageInfo

	if len(args) == 7 {

		chordFromArgs, err := chordFromArgs(args)
		if err != nil {
			return nil, nil, err
		}
		chord = &chordFromArgs

		imageInfoFromArgs, err := argsToImageInfo(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		imageInfo = &imageInfoFromArgs

	} else if len(args) == 3 {

		chord = nil
	}

	return chord, imageInfo, nil
}

func chordFromArgs(args []string) (chord, error) {

	fretCount, err := strconv.Atoi(args[1])
	if err != nil {
		return chord{}, fmt.Errorf("Could not convert %s to fretCount", args[1])
	}
	stringCount, err := strconv.Atoi(args[2])
	if err != nil {
		return chord{}, fmt.Errorf("Could not convert %s to stringCount", args[2])
	}

	frettedNotes, err := toFrettedNotes(args[3])
	if err != nil {
		return chord{}, err
	}

	return chord{args[0], fretCount, stringCount, frettedNotes}, nil
}

func argsToImageInfo(args []string) (imageInfo, error) {
	width, err := strconv.Atoi(args[4])
	if err != nil {
		return imageInfo{}, fmt.Errorf("Could not convert %s to width", args[4])
	}
	height, err := strconv.Atoi(args[5])
	if err != nil {
		return imageInfo{}, fmt.Errorf("Could not convert %s to height", args[5])
	}

	return imageInfo{width, height, args[6]}, nil
}

func toFrettedNotes(notesAsString string) ([]frettedNote, error) {
	notes := strings.Split(notesAsString, " ")
	result := make([]frettedNote, 0, len(notes))

	for _, v := range notes {
		pos := strings.Split(v, ":")
		fretValue, err := strconv.Atoi(pos[0][1:])
		if err != nil {
			return []frettedNote{}, fmt.Errorf("Invalid fret value: " + pos[0])
		}
		stringValue, err := strconv.Atoi(pos[1][1:])
		if err != nil {
			return []frettedNote{}, fmt.Errorf("Invalid string value: " + pos[1])
		}
		result = append(result, frettedNote{fretValue, stringValue})
	}
	return result, nil
}

type frettedNote struct {
	fretNumber   int
	stringNumber int
}

type chord struct {
	name        string
	fretCount   int
	stringCount int

	frettedNotes []frettedNote
}

type imageInfo struct {
	width      int
	height     int
	targetFile string
}
