package chordloader

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/tinosteinort/chord-renderer/types"
)

func TestSelectLoaderFromArgs(t *testing.T) {

	loader, err := selectLoader([]string{"", "", "", "", "", "", ""})
	if err != nil {
		t.Errorf("no error expected; got %s", err)
	}

	switch loader.(type) {
	case chordFromArgsLoader:
		break
	default:
		t.Errorf("want chordFromArgsLoader")
		break
	}
}

func TestSelectLoaderFromFile(t *testing.T) {

	loader, err := selectLoader([]string{"", "", "", ""})
	if err != nil {
		t.Errorf("no error expected; got %s", err)
	}

	switch loader.(type) {
	case chordFromFileLoader:
		break
	default:
		t.Errorf("want chordFromFileLoader")
		break
	}
}

func TestChordFromArgsLoader(t *testing.T) {
	args := []string{"G", "5", "4", "f2:s3 f3:s2 f2:s1", "", "", ""}
	loader := chordFromArgsLoader{Args: args}
	chord, err := loader.Load()

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

func TestChordFromFileLoader(t *testing.T) {

	expectedChord := &types.Chord{
		Name:        "G",
		FretCount:   5,
		StringCount: 4,
		FrettedNotes: []types.FrettedNote{
			types.FrettedNote{FretNumber: 2, StringNumber: 3},
			types.FrettedNote{FretNumber: 3, StringNumber: 2},
			types.FrettedNote{FretNumber: 2, StringNumber: 1},
		},
	}

	chordFromFile := &types.Chord{}

	data, err := ioutil.ReadFile("testdata/test.json")
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(data, chordFromFile)
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(expectedChord, chordFromFile) == false {
		t.Errorf("got  %v; want %v", chordFromFile, expectedChord)
	}
}
