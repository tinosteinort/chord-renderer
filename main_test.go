package main

import (
	"fmt"
	"testing"
)

func TestValidateArgsCount(t *testing.T) {

	var argTests = []struct {
		inputArgs []string
		valid     bool
	}{
		{[]string{""}, false},
		{[]string{"", ""}, false},
		{[]string{"", "", ""}, false},
		{[]string{"", "", "", ""}, true},
		{[]string{"", "", "", "", ""}, false},
		{[]string{"", "", "", "", "", ""}, false},
		{[]string{"", "", "", "", "", "", ""}, true},
		{[]string{"", "", "", "", "", "", "", ""}, false},
	}

	for _, tt := range argTests {
		t.Run(fmt.Sprintf("%d params allowed %t", len(tt.inputArgs), tt.valid), func(t *testing.T) {
			err := validateArgsCount(tt.inputArgs)
			if err == nil && !tt.valid {
				t.Errorf("%d arguments not allowed", len(tt.inputArgs))
			}
		})
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
