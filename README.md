# nameto

A utility for bulk file management

## Installation

The project is still in the early stages of development, at the moment it's only possible to install using `go install`

### `go install`

```sh
go install github.com/sftsrv/nameto@latest
```

## Usage

You can use the `--help` flag to view usage information:

```sh
nameto --help
```

`nameto` can be used in a few different ways:


### Interactively Editing Changeset Using Your Configured $EDITOR

```sh
nameto -f `.*\.go` -t 'new/path/$'

# commit changes
nameto -f `.*\.go` -t 'new/path/$' --commit --edit
```

### Using a Dry-Run and Accepting Changes

```sh
nameto -f `.*\.go` -t 'new/path/$'

# commit changes
nameto -f `.*\.go` -t 'new/path/$' --commit
```

### Using Existing Changeset File

```sh
# escaping of regex special chars will depend on your shell
nameto -f `.*\.go` -t 'new/path/$' --from-file path/to/changeset

# commit changes
nameto -f `.*\.go` -t 'new/path/$' --from-file path/to/changset --commit
```

## Additional Details

### Changeset Format

A changeset looks like so:

```
# Commented-out Lines Start with a Hash
R old/path/rename -> new/path/for/rename
C old/path/copy -> new/path/for/copy
```

### Using Regexes

The structure of the regexp provided should be compatible with Go's implementation,
with the following affordances made:

1. Named capture groups can be specified as '<name>' instead of '?P<name>'
2. Regexes are always matched against a single line - so multiline captures are not meaningful

### Using Patterns

References in pattern are indicated with a `$` in the pattern

	- `$` refers to the entire match
	- `$0` (entire match), `$1` (first capture group), etc. refer to capture groups in order of capture
	- `$<name>` or `$name` refer to named capture groups

## Features

- [x] Regex based file matching
- [x] Copy a file
- [x] Rename a file
- [x] Automatically create directories
- [x] Preview files to be moved/copied
- [x] Get list of files to work with from existing file
- [ ] Support files with spaces (Will implement this if it ever becomes something I care about)
