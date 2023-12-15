package older

import (
	"io/fs"
	"time"
)

func Files(fsys fs.FS, age time.Duration) []string {
	threshold := time.Now().Add(-age)
	var files []string
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil || info.IsDir() {
			return nil
		}
		if info.ModTime().Before(threshold) {
			files = append(files, path)
		}
		return nil
	})
	return files
}
