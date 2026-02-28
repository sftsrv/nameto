package lib_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/sftsrv/nameto/lib"
)

func TestSimplePattern(t *testing.T) {
	expected := []lib.Change{
		{"my/path/1.txt", "new/my/path/1.txt"},
		{"my/path/2.txt", "new/my/path/2.txt"},
	}

	paths := []string{}
	for _, c := range expected {
		paths = append(paths, c.Old)
	}

	re := regexp.MustCompile("my/path/.*")
	pattern := "new/$"

	result := lib.GenerateChanges(paths, re, pattern)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestNamedGroups(t *testing.T) {
	expected := []lib.Change{
		{"my/first/1.txt", "my/new/first_path_1.txt"},
		{"my/second/2.txt", "my/new/second_path_2.txt"},
	}

	paths := []string{}
	for _, c := range expected {
		paths = append(paths, c.Old)
	}

	re := regexp.MustCompile(`my/(?P<prefix>\w*)/(?P<name>.*)`)
	pattern := "my/new/$prefix_path_$<name>"

	result := lib.GenerateChanges(paths, re, pattern)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestUnnamedGroups(t *testing.T) {
	expected := []lib.Change{
		{"my/first/1.txt", "my/new/first_path_1.txt"},
		{"my/second/2.txt", "my/new/second_path_2.txt"},
	}

	paths := []string{}
	for _, c := range expected {
		paths = append(paths, c.Old)
	}

	re := regexp.MustCompile(`my/(\w*)/(.*)`)
	pattern := "my/new/$1_path_$2"

	result := lib.GenerateChanges(paths, re, pattern)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
