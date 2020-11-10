package chordinfo_test

import (
	"reflect"
	"testing"

	"github.com/tinosteinort/chord-renderer/chordinfo"
)

func TestLoadChordFromFile(t *testing.T) {
	args := []string{"testdata/g-chord.json", "", "", ""}
	chord, err := chordinfo.LoadChord(args)
	if err != nil {
		t.Errorf("Could not load chord %s", err)
	}

	expectedChord := chordinfo.Chord{
		Name:        "G",
		FretCount:   5,
		StringCount: 4,
		FrettedNotes: []chordinfo.FrettedNote{
			chordinfo.FrettedNote{FretNumber: 2, StringNumber: 3},
			chordinfo.FrettedNote{FretNumber: 3, StringNumber: 2},
			chordinfo.FrettedNote{FretNumber: 2, StringNumber: 1},
		},
	}

	if reflect.DeepEqual(expectedChord, chord) == false {
		t.Errorf("got  %v; want %v", chord, expectedChord)
	}
}
