package findgo_test

import (
	"archive/zip"
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

func TestFilesCorrectlyListsFilesInZIPArchive(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader("testdata/files.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"tree/file.go",
		"tree/subfolder/subfolder.go",
		"tree/subfolder2/another.go",
		"tree/subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}

func BenchmarkFilesOnMapFS(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}
