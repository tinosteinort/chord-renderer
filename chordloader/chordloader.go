package chordloader

import (
	"fmt"

	"github.com/tinosteinort/chord-renderer/types"
)

type chordLoader interface {
	Load() (types.Chord, error)
}

func LoadChord(args []string) (types.Chord, error) {
	loader, err := selectLoader(args)
	if err != nil {
		return types.Chord{}, err
	}
	return loader.Load()
}

func selectLoader(args []string) (chordLoader, error) {
	if len(args) == 7 {
		return chordFromArgsLoader{Args: args}, nil

	} else if len(args) == 4 {
		return chordFromFileLoader{Args: args}, nil
	}
	return nil, fmt.Errorf("Could not determine loader for: %v", args)
}
