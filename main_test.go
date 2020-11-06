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
