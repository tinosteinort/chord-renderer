package chordloader_test

import (
	"reflect"
	"testing"

	"github.com/tinosteinort/chord-renderer/chordloader"
)

func TestLoadChordFromFile(t *testing.T) {
	args := []string{"testdata/g-chord.json", "", "", ""}
	chord, err := chordloader.LoadChord(args)
	if err != nil {
		t.Errorf("Could not load chord %s", err)
	}

	expectedChord := chordloader.Chord{
		Name:        "G",
		FretCount:   5,
		StringCount: 4,
		FrettedNotes: []chordloader.FrettedNote{
			chordloader.FrettedNote{FretNumber: 2, StringNumber: 3},
			chordloader.FrettedNote{FretNumber: 3, StringNumber: 2},
			chordloader.FrettedNote{FretNumber: 2, StringNumber: 1},
		},
	}

	if reflect.DeepEqual(expectedChord, chord) == false {
		t.Errorf("got  %v; want %v", chord, expectedChord)
	}
}
