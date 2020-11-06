package chordloader

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tinosteinort/chord-renderer/types"
)

type chordFromArgsLoader struct {
	Args []string
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