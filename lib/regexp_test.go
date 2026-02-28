package lib_test

import (
	"testing"

	"github.com/sftsrv/nameto/lib"
)

func TestCreateValidRegexp(t *testing.T) {
	input := "my-(<escaped>)-regex"
	expected := "my-(?P<escaped>)-regex"

	result, err := lib.CreateRegexp(input)

	if err != nil {
		t.Errorf("Error creating regexp %v", err)
	}

	if result.String() != expected {
		t.Errorf("Invalid regexp, expected: %s got: %s", expected, result)
	}
}

func TestCreateInvalidRegexp(t *testing.T) {
	input := "my-($invalid)-regex"

	result, err := lib.CreateRegexp(input)

	if err != nil {
		t.Errorf("Expected error but got %s", result.String())
	}
}
