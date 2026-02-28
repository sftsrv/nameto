package lib

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func FindPaths(re *regexp.Regexp) []string {
	var paths []string

	filepath.WalkDir(".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			isMatch := re.MatchString(path)
			if d.IsDir() || !isMatch {
				return nil
			}

			paths = append(paths, path)

			return nil
		},
	)

	return paths
}

func exists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}

func renameFile(old, new string) error {
	if exists(new) {
		return fmt.Errorf("Path already exists, will not overwrite %s", new)
	}

	err := copyFile(old, new)
	if err != nil {
		return err
	}

	err = os.Remove(old)
	if err != nil {
		return err
	}

	return nil
}

func RenameFiles(changes Changes) error {
	for _, change := range changes {
		err := renameFile(change.Old, change.New)
		if err != nil {
			return err
		}
	}

	return nil
}

// implementation from https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file
func copyFile(old, new string) error {
	if exists(new) {
		return fmt.Errorf("Path already exists, will not overwrite %s", new)
	}

	dir := filepath.Dir(new)
	if !exists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}

	o, err := os.Open(old)
	if err != nil {
		panic(err)
	}

	defer o.Close()
	n, err := os.Create(new)
	if err != nil {
		panic(err)
	}

	defer n.Close()

	n.ReadFrom(o)

	return nil
}

func CopyFiles(changes Changes) error {
	for _, change := range changes {
		err := copyFile(change.Old, change.New)
		if err != nil {
			return err
		}
	}

	return nil
}
