package chordloader

import (
	"fmt"
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
	loader, err := selectLoader(args)
	if err != nil {
		return Chord{}, err
	}
	return loader.Load()
}

type chordLoader interface {
	Load() (Chord, error)
}

func selectLoader(args []string) (chordLoader, error) {
	if len(args) == 7 {
		return chordFromArgsLoader{Args: args}, nil

	} else if len(args) == 4 {
		return chordFromFileLoader{Args: args}, nil
	}
	return nil, fmt.Errorf("Could not determine loader for: %v", args)
}
