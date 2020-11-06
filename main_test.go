package main

import (
	"fmt"
	"testing"
)

func TestValidateArgsCount(t *testing.T) {
	var argTests = []struct {
		length  int
		allowed bool
	}{
		{1, false},
		{2, false},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, true},
		{8, false},
	}

	for _, tt := range argTests {
		testname := fmt.Sprintf("%d params allowed %t", tt.length, tt.allowed)
		args := make([]string, tt.length)
		t.Run(testname, func(t *testing.T) {
			err := validateArgsCount(args)
			if err != nil && tt.allowed {
				t.Errorf("%s arguments not allowed", err)
			}
		})
	}
}
