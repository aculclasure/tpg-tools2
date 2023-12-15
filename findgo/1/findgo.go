package findgo

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Files(rootPath string) []string {
	var files []string
	fsys := os.DirFS(rootPath)
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})
	return files
}
