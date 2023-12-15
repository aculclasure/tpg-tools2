package findgo_test

import (
	"os"
	"testing"
	"testing/fstest"

	"github.com/aculclasure/findgo"
	"github.com/google/go-cmp/cmp"
)

func TestFilesCorrectlyListsFilesInTree(t *testing.T) {
	t.Parallel()
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	fsys := os.DirFS("testdata/tree")
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFilesCorrectlyListsFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
