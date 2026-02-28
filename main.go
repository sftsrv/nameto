package main

import (
	"flag"
	"fmt"
	"os"
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
	copyFlag := flag.Bool("c", false, "copy files instead of rename")
	fromFlag := flag.String("f", ".*", "regex for matching files")
	toFlag := flag.String("t", "$", "pattern to use when renaming files")
	fileFlag := flag.String("from-file", "", "")
	dryRunFlag := flag.Bool("dry-run", false, "print out results, do not execute changes")
	editorFlag := flag.String("editor", defaultEditor, "editor to edit file paths with")
	noEditFlag := flag.Bool("y", false, "accept changes without previewing or editing")

	flag.Parse()

	if *helpFlag {
		fmt.Print(usage)
		flag.Usage()
		return
	}

	fmt.Println(fromFlag, toFlag, editorFlag, dryRunFlag, copyFlag, fileFlag, noEditFlag)
}
