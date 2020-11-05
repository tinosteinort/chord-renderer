package chordloader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tinosteinort/chord-renderer/types"
)

type ChordLoader interface {
	Load() (types.Chord, error)
}

type chordFromArgsLoader struct {
	Args []string
}

type chordFromFileLoader struct {
	Args []string
}

func LoadChord(args []string) (types.Chord, error) {
	loader, err := selectLoader(args)
	if err != nil {
		return types.Chord{}, err
	}

	chord, err := loader.Load()
	if err != nil {
		return types.Chord{}, err
	}
	return chord, nil
}

func selectLoader(args []string) (ChordLoader, error) {
	if len(args) == 7 {
		return chordFromArgsLoader{Args: args}, nil

	} else if len(args) == 4 {
		return chordFromFileLoader{Args: args}, nil
	}
	return nil, fmt.Errorf("Could not determine loader for: %v", args)
}

func (loader chordFromArgsLoader) Load() (types.Chord, error) {
	fretCount, err := strconv.Atoi(loader.Args[1])
	if err != nil {
		return types.Chord{}, fmt.Errorf("Could not convert %s to fretCount", loader.Args[1])
	}
	stringCount, err := strconv.Atoi(loader.Args[2])
	if err != nil {
		return types.Chord{}, fmt.Errorf("Could not convert %s to stringCount", loader.Args[2])
	}

	frettedNotes, err := toFrettedNotes(loader.Args[3])
	if err != nil {
		return types.Chord{}, err
	}

	return types.Chord{
		Name:         loader.Args[0],
		FretCount:    fretCount,
		StringCount:  stringCount,
		FrettedNotes: frettedNotes,
	}, nil
}

func toFrettedNotes(notesAsString string) ([]types.FrettedNote, error) {
	notes := strings.Split(notesAsString, " ")
	result := make([]types.FrettedNote, 0, len(notes))

	for _, v := range notes {
		pos := strings.Split(v, ":")
		fretValue, err := strconv.Atoi(pos[0][1:])
		if err != nil {
			return []types.FrettedNote{}, fmt.Errorf("Invalid fret value: " + pos[0])
		}
		stringValue, err := strconv.Atoi(pos[1][1:])
		if err != nil {
			return []types.FrettedNote{}, fmt.Errorf("Invalid string value: " + pos[1])
		}
		result = append(result, types.FrettedNote{
			FretNumber:   fretValue,
			StringNumber: stringValue,
		})
	}
	return result, nil
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
