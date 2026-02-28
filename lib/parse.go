package lib

import (
	"fmt"
	"regexp"
	"strings"
)

type Change struct {
	Old string
	New string
}

var spaces = regexp.MustCompile(`\s+`)

func parseLine(line string) (Change, error) {
	l := strings.TrimSpace(line)
	parts := spaces.Split(l, -1)

	if len(parts) != 2 {
		return Change{}, fmt.Errorf("line did not match the expected structure:\n%s", line)
	}

	old := parts[0]
	new := parts[1]

	return Change{old, new}, nil
}

func ParseFile(content string) ([]Change, error) {
	result := []Change{}

	lines := strings.Lines(content)
	for l := range lines {
		if l == "" || l[0] == '#' {
			continue
		}

		line, err := parseLine(l)

		if err != nil {
			return result, err
		}

		result = append(result, line)
	}

	return result, nil
}
