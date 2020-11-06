package chordloader_test

import (
	"reflect"
	"testing"

	"github.com/tinosteinort/chord-renderer/chordloader"
)

func TestLoadChordFromArgs(t *testing.T) {
	args := []string{"G", "5", "4", "f2:s3 f3:s2 f2:s1", "", "", ""}
	chord, err := chordloader.LoadChord(args)

	if err != nil {
		t.Error(err)
	}
	if chord.Name != "G" {
		t.Errorf("got %s; want G", chord.Name)
	}
	if chord.FretCount != 5 {
		t.Errorf("got %d; want 5", chord.FretCount)
	}
	if chord.StringCount != 4 {
		t.Errorf("got %d; want 4", chord.StringCount)
	}
	if len(chord.FrettedNotes) != 3 {
		t.Errorf("got %d; want 3", len(chord.FrettedNotes))
	}
	if chord.FrettedNotes[0].FretNumber != 2 {
		t.Errorf("got %d; want 2", chord.FrettedNotes[0].FretNumber)
	}
	if chord.FrettedNotes[0].StringNumber != 3 {
		t.Errorf("got %d; want 3", chord.FrettedNotes[0].StringNumber)
	}
	if chord.FrettedNotes[1].FretNumber != 3 {
		t.Errorf("got %d; want 3", chord.FrettedNotes[1].FretNumber)
	}
	if chord.FrettedNotes[1].StringNumber != 2 {
		t.Errorf("got %d; want 2", chord.FrettedNotes[1].StringNumber)
	}
	if chord.FrettedNotes[2].FretNumber != 2 {
		t.Errorf("got %d; want 2", chord.FrettedNotes[2].FretNumber)
	}
	if chord.FrettedNotes[2].StringNumber != 1 {
		t.Errorf("got %d; want 1", chord.FrettedNotes[2].StringNumber)
	}
}

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
