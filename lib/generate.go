package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func generateChange(mode ChangeMode, path string, re *regexp.Regexp, pattern string) Change {
	result := pattern
	names := re.SubexpNames()

	matches := re.FindAllStringSubmatch(path, -1)
	for _, match := range matches {
		for groupIndex, submatch := range match {
			groupName := names[groupIndex]

			if groupName == "" {
				groupName = strconv.Itoa(groupIndex)
			}

			replacers := []string{
				fmt.Sprintf("$%s", groupName),
				fmt.Sprintf("$<%s>", groupName),
			}

			for _, replacer := range replacers {
				result = strings.ReplaceAll(result, replacer, submatch)
			}
		}
	}

	// Replace lone $ only after all others have been done.
	// This is a convenience syntax for $0
	result = strings.ReplaceAll(result, "$", path)

	return Change{mode, path, result}
}

func GenerateChanges(mode ChangeMode, paths []string, re *regexp.Regexp, pattern string) Changes {
	result := Changes{}

	for _, p := range paths {
		change := generateChange(mode, p, re, pattern)
		result = append(result, change)
	}

	return result
}
