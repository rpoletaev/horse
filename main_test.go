package main

import (
	"fmt"
	"testing"
)

type testEntry struct {
	Input    string
	Expected string
}

func TestInputs(t *testing.T) {
	entries := []testEntry{
		{"b1", `[a3 c3 d2]`},
		{"11", "Wrong position letter"},
		{"1b", "Wrong position letter"},
		{"k9", "Wrong position letter"},
		{"a0", "Wrong position number"},
		{"a9", "Wrong position number"},
		{"b4", `[a6 c6 d5 d3 a2 c2]`},
	}

	for _, entry := range entries {
		res, err := getAvailablePositions(entry.Input)
		if err != nil {
			if err.Error() != entry.Expected {
				t.Errorf(`Error on value %s: expected %s, got %s`, entry.Input, entry.Expected, err.Error())
			}
			continue
		}

		if fmt.Sprintf("%v", res) != entry.Expected {
			t.Errorf(`Error on value %s: expected %s, got %s`, entry.Input, entry.Expected, res)
		}
	}
}
