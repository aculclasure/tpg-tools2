package kv_test

import (
	"os"
	"testing"

	"github.com/aculclasure/kv"
	"github.com/google/go-cmp/cmp"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"kv": kv.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestAllReturnsAllKeysAndValues(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore(t.TempDir() + "/temp.store")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("k1", "v1")
	s.Set("k2", "v2")
	s.Set("k3", "v3")
	want := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}
	got := s.All()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	store, err := kv.OpenStore("")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := store.Get("key")
	if ok {
		t.Error("wanted ok to be false when getting non-existent key")
	}
}

func TestGetReturnsValueAndOKIfKeyDoesExist(t *testing.T) {
	t.Parallel()
	store, err := kv.OpenStore("")
	if err != nil {
		t.Fatal(err)
	}
	store.Set("key", "value")
	v, ok := store.Get("key")
	if !ok {
		t.Fatal("wanted ok to be true when getting key that exists")
	}
	if v != "value" {
		t.Errorf("want 'value', got %s", v)
	}
}

func TestSetUpdatesExistingKeyToNewValue(t *testing.T) {
	t.Parallel()
	store, err := kv.OpenStore("")
	if err != nil {
		t.Fatal(err)
	}
	store.Set("key", "initialValue")
	store.Set("key", "updatedValue")
	v, ok := store.Get("key")
	if !ok {
		t.Fatal("wanted ok to be true when getting key that exists")
	}
	if v != "updatedValue" {
		t.Errorf("want 'updatedValue', got %s", v)
	}
}

func TestSaveSavesDataPersistently(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/test.store"
	store1, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	store1.Set("key1", "value1")
	store1.Set("key2", "value2")
	store1.Set("key3", "value3")
	err = store1.Save()
	if err != nil {
		t.Fatal(err)
	}
	store2, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := store2.Get("key1"); v != "value1" {
		t.Errorf("want key1 to have value 'value1', got %s", v)
	}
	if v, _ := store2.Get("key2"); v != "value2" {
		t.Errorf("want key2 to have value 'value2', got %s", v)
	}
	if v, _ := store2.Get("key3"); v != "value3" {
		t.Errorf("want key3 to have value 'value3', got %s", v)
	}
}

func TestSaveErrorsOnUnwritablePath(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("fakedir/unwritable.store")
	if err != nil {
		t.Fatal(err)
	}
	err = s.Save()
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ReturnsErrorForUnreadablePath(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/unreadable.store"
	if _, err := os.Create(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0000); err != nil {
		t.Fatal(err)
	}
	_, err := kv.OpenStore(path)
	if err == nil {
		t.Fatal("no error")
	}
}

func TestOpenStore_ReturnsErrorOnInvalidData(t *testing.T) {
	t.Parallel()
	_, err := kv.OpenStore("testdata/empty.store")
	if err == nil {
		t.Fatal("no error")
	}
}
