package findgo

import (
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) []string {
	var files []string
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})
	return files
}
