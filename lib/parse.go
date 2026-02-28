package lib

import (
	"fmt"
	"regexp"
	"strings"
)

var separator = regexp.MustCompile(`\s+->\s+`)

func parseLine(line string) (Change, error) {
	l := strings.TrimSpace(line)
	if len(line) < 3 {
		return Change{}, fmt.Errorf("expected at least 3 characters in line: %s", line)
	}

	mode := NewMode(l[0:1])
	parts := separator.Split(strings.TrimSpace(l[2:]), -1)

	if len(parts) != 2 {
		return Change{}, fmt.Errorf("line did not match the expected structure:\n%s", line)
	}

	old := parts[0]
	new := parts[1]

	return Change{mode, old, new}, nil
}

func ParseFile(content string) (Changes, error) {
	result := Changes{}

	lines := strings.Lines(content)
	for l := range lines {
		l = strings.TrimSpace(l)
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
