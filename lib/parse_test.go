package lib_test

import (
	"reflect"
	"testing"

	"github.com/sftsrv/nameto/lib"
)

func TestParseValid(t *testing.T) {
	content := `# some example paths to parse
C from/my/path.txt -> to/my/path.txt
  R   some/more_complicated.path/to  ->        another/.poorly.formatted_type./path
`

	expected := lib.Changes{
		{lib.ChangeModeCopy, "from/my/path.txt", "to/my/path.txt"},
		{lib.ChangeModeRename, "some/more_complicated.path/to", "another/.poorly.formatted_type./path"},
	}

	result, _ := lib.ParseFile(content)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n%v \n\nbut got:\n%v", expected, result)
	}
}

func TestParseInvalid(t *testing.T) {
	content := `# some example paths to parse
	this/file/has/some/weird/structure.x without/correct/separator
`

	_, err := lib.ParseFile(content)

	if err == nil {
		t.Errorf("Parsing did not return an error")
	}
}
