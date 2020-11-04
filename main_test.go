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

func TestArgsToImageInfo(t *testing.T) {
	args := []string{"", "", "", "", "100", "200", "./file"}
	info, err := argsToImageInfo(args)

	if err != nil {
		t.Error(err)
	}
	if info.Width != 100 {
		t.Errorf("got %d; want 100", info.Width)
	}
	if info.Height != 200 {
		t.Errorf("got %d; want 200", info.Height)
	}
	if info.TargetFile != "./file" {
		t.Errorf("go %s; want ./file", info.TargetFile)
	}
}
