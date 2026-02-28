package lib

import (
	"fmt"
	"strings"
)

type ChangeMode string

const (
	ChangeModeRename ChangeMode = "R"
	ChangeModeCopy   ChangeMode = "C"
)

type Change struct {
	Mode ChangeMode
	Old  string
	New  string
}

func NewMode(str string) ChangeMode {
	if str == string(ChangeModeRename) {
		return ChangeModeRename
	}

	return ChangeModeCopy
}

func (change Change) String() string {
	return fmt.Sprintf("%s %s -> %s", change.Mode, change.Old, change.New)
}

type Changes []Change

func (changes Changes) String() string {
	lines := []string{}
	for _, change := range changes {
		lines = append(lines, change.String())
	}

	return strings.Join(lines, "\r\n")
}
