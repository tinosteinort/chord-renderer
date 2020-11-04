package main

import (
	"testing"
)

func TestValidateArgsCount(t *testing.T) {
	tooFewArgs := []string{"", "", "", "", "", ""}
	errTooFew := validateArgsCount(tooFewArgs)

	if errTooFew == nil {
		t.Errorf("Expected error, but got none")
	}
	if errTooFew != nil && errTooFew.Error() != "invalid amount of arguments" {
		t.Errorf("got %s; want 'invalid amount of arguments'", errTooFew.Error())
	}

	tooManyArgs := []string{"", "", "", "", "", "", "", ""}
	errTooMany := validateArgsCount(tooManyArgs)

	if errTooMany == nil {
		t.Errorf("Expected error, but got none")
	}
	if errTooMany != nil && errTooMany.Error() != "invalid amount of arguments" {
		t.Errorf("got %s; want 'invalid amount of arguments'", errTooMany.Error())
	}
}

func TestChordFromArgs(t *testing.T) {
	args := []string{"G", "5", "4", "f2:s3 f3:s2 f2:s1", "", "", ""}
	chord, err := chordFromArgs(args)

	if err != nil {
		t.Error(err)
	}
	if chord.name != "G" {
		t.Errorf("got %s; want G", chord.name)
	}
	if chord.fretCount != 5 {
		t.Errorf("got %d; want 5", chord.fretCount)
	}
	if chord.stringCount != 4 {
		t.Errorf("got %d; want 4", chord.stringCount)
	}
	if len(chord.frettedNotes) != 3 {
		t.Errorf("got %d; want 3", len(chord.frettedNotes))
	}
	if chord.frettedNotes[0].fretNumber != 2 {
		t.Errorf("got %d; want 2", chord.frettedNotes[0].fretNumber)
	}
	if chord.frettedNotes[0].stringNumber != 3 {
		t.Errorf("got %d; want 3", chord.frettedNotes[0].stringNumber)
	}
	if chord.frettedNotes[1].fretNumber != 3 {
		t.Errorf("got %d; want 3", chord.frettedNotes[1].fretNumber)
	}
	if chord.frettedNotes[1].stringNumber != 2 {
		t.Errorf("got %d; want 2", chord.frettedNotes[1].stringNumber)
	}
	if chord.frettedNotes[2].fretNumber != 2 {
		t.Errorf("got %d; want 2", chord.frettedNotes[2].fretNumber)
	}
	if chord.frettedNotes[2].stringNumber != 1 {
		t.Errorf("got %d; want 1", chord.frettedNotes[2].stringNumber)
	}
}

func TestArgsToImageInfo(t *testing.T) {
	args := []string{"", "", "", "", "100", "200", "./file"}
	info, err := argsToImageInfo(args)

	if err != nil {
		t.Error(err)
	}
	if info.width != 100 {
		t.Errorf("got %d; want 100", info.width)
	}
	if info.height != 200 {
		t.Errorf("got %d; want 200", info.height)
	}
	if info.targetFile != "./file" {
		t.Errorf("go %s; want ./file", info.targetFile)
	}
}
