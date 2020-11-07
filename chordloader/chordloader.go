package chordloader

import (
	"encoding/json"
	"io/ioutil"
)

type Chord struct {
	Name         string
	FretCount    int
	StringCount  int
	FrettedNotes []FrettedNote
}

type FrettedNote struct {
	FretNumber   int
	StringNumber int
}

func LoadChord(args []string) (Chord, error) {
	chordFile := args[0]

	chordFromFile := &Chord{}

	data, err := ioutil.ReadFile(chordFile)
	if err != nil {
		return Chord{}, err
	}

	err = json.Unmarshal(data, chordFromFile)
	if err != nil {
		return Chord{}, err
	}

	return *chordFromFile, nil
}
