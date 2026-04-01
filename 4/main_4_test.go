package main

import (
	"reflect"
	"testing"
)

func TestSliceDiffBasic(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date"}
	slice2 := []string{"banana", "date"}

	got := sliceDiff(slice1, slice2)
	want := []string{"apple", "cherry"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSliceDiffNoMatches(t *testing.T) {
	slice1 := []string{"apple", "banana"}
	slice2 := []string{"fig", "grape"}

	got := sliceDiff(slice1, slice2)
	want := []string{"apple", "banana"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSliceDiffAllRemoved(t *testing.T) {
	slice1 := []string{"apple", "banana"}
	slice2 := []string{"apple", "banana"}

	got := sliceDiff(slice1, slice2)
	want := []string{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSliceDiffEmptyFirstSlice(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"apple"}

	got := sliceDiff(slice1, slice2)
	want := []string{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSliceDiffDuplicatesInFirstSlice(t *testing.T) {
	slice1 := []string{"apple", "apple", "banana", "cherry"}
	slice2 := []string{"banana"}

	got := sliceDiff(slice1, slice2)
	want := []string{"apple", "apple", "cherry"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
