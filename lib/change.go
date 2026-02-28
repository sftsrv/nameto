package lib

import (
	"fmt"
	"strings"
)

type Change struct {
	Old string
	New string
}

func (change Change) String() string {
	return fmt.Sprintf("%s -> %s", change.Old, change.New)
}

type Changes []Change

func (changes Changes) String() string {
	lines := []string{}
	for _, change := range changes {
		lines = append(lines, change.String())
	}

	return strings.Join(lines, "\r\n")
}
