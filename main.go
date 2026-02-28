package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sftsrv/nameto/lib"
)

const usage = `nameto

A utility for bulk file management

## Usage

You can use the '--help' flag to view usage information:

'''sh
nameto --help
'''

'nameto' can be used in a few different ways:


### Interactively Editing Changeset Using Your Configured $EDITOR

'''sh
nameto -f '.*\.go' -t 'new/path/$'

# commit changes
nameto -f '.*\.go' -t 'new/path/$' --commit --edit
'''

### Using a Dry-Run and Accepting Changes

'''sh
nameto -f '.*\.go' -t 'new/path/$'

# commit changes
nameto -f '.*\.go' -t 'new/path/$' --commit
'''

### Using Existing Changeset File

'''sh
# escaping of regex special chars will depend on your shell
nameto -f '.*\.go' -t 'new/path/$' --from-file path/to/changeset

# commit changes
nameto -f '.*\.go' -t 'new/path/$' --from-file path/to/changset --commit
'''

## Additional Details

### Changeset Format

A changeset looks like so:


'''
# Commented-out Lines Start with a Hash
R old/path/rename -> new/path/for/rename
C old/path/copy -> new/path/for/copy
'''

### Using Regexes

The structure of the regexp provided should be compatible with Go's implementation,
with the following affordances made:

1. Named capture groups can be specified as '<name>' instead of '?P<name>'
2. Regexes are always matched against a single line - so multiline captures are not meaningful

### Using Patterns

References in pattern are indicated with a '$' in the pattern

	- '$' refers to the entire match
	- '$0' (entire match), '$1' (first capture group), etc. refer to capture groups in order of capture
	- '$<name>' or '$name' refer to named capture groups
	
`

func main() {
	defaultEditor, _ := os.LookupEnv("EDITOR")

	helpFlag := flag.Bool("help", false, "show usage info")
	fileFlag := flag.String("from-file", "", "use an existing changeset instead of creating one")
	commitFlag := flag.Bool("commit", false, "execute changes")
	editorFlag := flag.String("editor", defaultEditor, "editor to edit file paths with")

	renameFlag := flag.Bool("r", false, "rename files by default instead of copy")
	fromFlag := flag.String("f", ".*", "regex for matching files")
	toFlag := flag.String("t", "$", "pattern to use when renaming files")
	editFlag := flag.Bool("edit", false, "accept changes without editing")

	flag.Parse()

	if *helpFlag {
		fmt.Print(usage)
		flag.Usage()
		return
	}

	var changeFile string
	rename := *renameFlag

	if *fileFlag != "" {
		file, err := os.ReadFile(*fileFlag)
		if err != nil {
			panic(err)
		}

		changeFile = string(file)
	} else {
		re, err := lib.CreateRegexp(*fromFlag)
		if err != nil {
			panic(fmt.Errorf("Failed to parse given regexp '%s' with error: %v", *fromFlag, err))
		}

		paths := lib.FindPaths(re)

		var mode lib.ChangeMode = lib.ChangeModeCopy
		if rename {
			mode = lib.ChangeModeRename
		}

		changes := lib.GenerateChanges(mode, paths, re, *toFlag)
		changeFile = changes.String()
	}

	commit := *commitFlag
	edit := *editFlag

	if edit {
		editor := *editorFlag
		fmt.Println("Opening changes with", editor)
		result, err := lib.EditFile(editor, changeFile)
		if err != nil {
			panic(fmt.Errorf("Error editing file with %s with error: %v", *editorFlag, err))
		}

		fmt.Println("Editing complete")

		changeFile = result
	}

	changes, err := lib.ParseFile(changeFile)
	if err != nil {
		panic(fmt.Errorf("Error parsing change file: %v", err))
	}

	if !commit {
		fmt.Println(changeFile)
		return
	}

	fmt.Println("Executing changes")
	lib.PersistChanges(changes)
}
