package lib_test

import (
	"reflect"
	"testing"

	"github.com/sftsrv/nameto/lib"
)

func TestParseValid(t *testing.T) {
	content := `# some example paths to parse
from/my/path.txt to/my/path.txt
     some/more_complicated.path/to          another/.poorly.formatted_type./path
`

	expected := []lib.Change{
		{"from/my/path.txt", "to/my/path.txt"},
		{"some/more_complicated.path/to", "another/.poorly.formatted_type./path"},
	}

	result, _ := lib.ParseFile(content)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestParseInvalid(t *testing.T) {
	content := `# some example paths to parse
	this/file/has/some/weird/structure.x
`

	_, err := lib.ParseFile(content)

	if err == nil {
		t.Errorf("Parsing did not return an error")
	}
}
