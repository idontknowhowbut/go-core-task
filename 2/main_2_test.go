package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	got := sliceExample(input)
	want := []int{2, 4, 6}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	got := addElements(input, 4)
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	got := copySlice(input)

	input[0] = 999

	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{10, 20, 30, 40}
	got := removeElement(input, 1)
	want := []int{10, 30, 40}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRemoveElementInvalidIndex(t *testing.T) {
	input := []int{10, 20, 30}
	got := removeElement(input, 50)
	want := []int{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
