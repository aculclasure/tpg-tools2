package older_test

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/aculclasure/older"
	"github.com/google/go-cmp/cmp"
)

func TestFilesCorrectlyListFilesInMapFSOlderThanGivenDuration(t *testing.T) {
	t.Parallel()
	now := time.Now()
	fsys := fstest.MapFS{
		"now_file.go":                              {ModTime: now},
		"file_created_15_seconds_ago.go":           {ModTime: now.Add(-15 * time.Second)},
		"file_created_30_seconds_ago.go":           {ModTime: now.Add(-30 * time.Second)},
		"file_created_1_minute_ago.go":             {ModTime: now.Add(-time.Minute)},
		"subfolder/file_created_30_minutes_ago.go": {ModTime: now.Add(-30 * time.Minute)},
	}
	want := []string{
		"file_created_1_minute_ago.go",
		"file_created_30_seconds_ago.go",
		"subfolder/file_created_30_minutes_ago.go",
	}
	got := older.Files(fsys, 30*time.Second)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
