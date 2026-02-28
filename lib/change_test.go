package lib_test

import (
	"reflect"
	"testing"

	"github.com/sftsrv/nameto/lib"
)

func TestGenratedChangeParses(t *testing.T) {

	input := lib.Changes{
		{lib.ChangeModeRename, "from/my/path.txt", "to/my/path.txt"},
		{lib.ChangeModeCopy, "some/more_complicated.path/to", "another/.poorly.formatted_type./path"},
	}

	output, _ := lib.ParseFile(input.String())

	if !reflect.DeepEqual(output, input) {
		t.Errorf("Expected \n%v \n\ngot: \n%v", output, input)
	}
}
