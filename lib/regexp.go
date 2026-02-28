package lib

import (
	"regexp"
	"strings"
)

func CreateRegexp(re string) (*regexp.Regexp, error) {
	formatted := strings.ReplaceAll(re, `(<`, `(?P<`)

	return regexp.Compile(formatted)
}
