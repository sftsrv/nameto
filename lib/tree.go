package lib

import (
	"io/fs"
	"path/filepath"
	"regexp"
)

func FindPaths(root string, re regexp.Regexp) []string {
	var paths []string

	filepath.WalkDir(root,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			isMatch := !re.MatchString(path)
			if d.IsDir() || !isMatch {
				return nil
			}

			paths = append(paths, path)

			return nil
		},
	)

	return paths
}
