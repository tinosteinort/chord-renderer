package chordloader

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tinosteinort/chord-renderer/types"
)

type chordFromFileLoader struct {
	Args []string
}

func (loader chordFromFileLoader) Load() (types.Chord, error) {
	chordFile := loader.Args[0]

	chordFromFile := &types.Chord{}

	data, err := ioutil.ReadFile(chordFile)
	if err != nil {
		return types.Chord{}, err
	}

	err = json.Unmarshal(data, chordFromFile)
	if err != nil {
		return types.Chord{}, err
	}

	return *chordFromFile, nil
}
