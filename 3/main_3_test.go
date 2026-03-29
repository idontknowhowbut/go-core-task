package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := make(StringIntMap)

	m.Add("one", 1)

	got := map[string]int(m)
	want := map[string]int{"one": 1}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	m := make(StringIntMap)
	m.Add("one", 1)
	m.Add("two", 2)

	m.Remove("one")

	got := map[string]int(m)
	want := map[string]int{"two": 2}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCopy(t *testing.T) {
	m := make(StringIntMap)
	m.Add("one", 1)
	m.Add("two", 2)

	copied := m.Copy()
	m.Add("three", 3)

	want := map[string]int{
		"one": 1,
		"two": 2,
	}

	if !reflect.DeepEqual(copied, want) {
		t.Fatalf("got %v, want %v", copied, want)
	}
}

func TestExists(t *testing.T) {
	m := make(StringIntMap)
	m.Add("one", 1)

	if !m.Exists("one") {
		t.Fatal(`expected key "one" to exist`)
	}

	if m.Exists("two") {
		t.Fatal(`expected key "two" not to exist`)
	}
}

func TestGet(t *testing.T) {
	m := make(StringIntMap)
	m.Add("one", 1)

	value, ok := m.Get("one")
	if !ok || value != 1 {
		t.Fatalf("got (%d, %v), want (1, true)", value, ok)
	}

	value, ok = m.Get("two")
	if ok || value != 0 {
		t.Fatalf("got (%d, %v), want (0, false)", value, ok)
	}
}
