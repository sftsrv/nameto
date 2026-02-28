package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sftsrv/nameto/lib"
)

const usage = `nameto

> using regexes

The structure of the regexp provided should be compatible with Go's implementation,
with the following affordances made:

1. named capture groups can be specified as '<name>' instead of '?P<name>'
2. regexes are always matched against a single line - so multiline captures are not meaningful

> using patterns

references in pattern are indicated with a $ in the pattern
	- '$' refers to the entire match
	- '$0' (entire match), '$1' (first capture group), etc. refer to capture groups in order of capture
	- '$<name>' or '$name' refer to named capture groups

`

func main() {
	defaultEditor, _ := os.LookupEnv("EDITOR")

	helpFlag := flag.Bool("help", false, "show usage info")
	fileFlag := flag.String("from-file", "", "use an existing changeset instead of creating one")
	dryRunFlag := flag.Bool("dry-run", false, "print out results, do not execute changes")
	editorFlag := flag.String("editor", defaultEditor, "editor to edit file paths with")

	copyFlag := flag.Bool("c", false, "copy files instead of rename")
	fromFlag := flag.String("f", ".*", "regex for matching files")
	toFlag := flag.String("t", "$", "pattern to use when renaming files")
	noEditFlag := flag.Bool("y", false, "accept changes without previewing or editing")

	flag.Parse()

	if *helpFlag {
		fmt.Print(usage)
		flag.Usage()
		return
	}

	var changeFile string

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

		changes := lib.GenerateChanges(paths, re, *toFlag)
		changeFile = changes.String()
	}

	if *dryRunFlag {
		fmt.Println(changeFile)
		return
	}

	edit := !*noEditFlag
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

	if *copyFlag {
		lib.CopyFiles(changes)
	} else {
		lib.RenameFiles(changes)
	}
}
