package chordloader

import (
	"encoding/json"
	"io/ioutil"
)

type chordFromFileLoader struct {
	Args []string
}

func (loader chordFromFileLoader) Load() (Chord, error) {
	chordFile := loader.Args[0]

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
