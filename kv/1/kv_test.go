package kv_test

import (
	"testing"

	"github.com/aculclasure/kv"
)

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
