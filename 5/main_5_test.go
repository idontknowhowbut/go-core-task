package main

import (
	"reflect"
	"testing"
)

func TestIntersectBasic(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	gotBool, gotSlice := intersect(a, b)

	wantBool := true
	wantSlice := []int{3, 64}

	if gotBool != wantBool {
		t.Fatalf("got bool %v, want %v", gotBool, wantBool)
	}

	if !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Fatalf("got slice %v, want %v", gotSlice, wantSlice)
	}
}

func TestIntersectNoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	gotBool, gotSlice := intersect(a, b)

	wantBool := false
	wantSlice := []int{}

	if gotBool != wantBool {
		t.Fatalf("got bool %v, want %v", gotBool, wantBool)
	}

	if !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Fatalf("got slice %v, want %v", gotSlice, wantSlice)
	}
}

func TestIntersectAllMatch(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	gotBool, gotSlice := intersect(a, b)

	wantBool := true
	wantSlice := []int{1, 2, 3}

	if gotBool != wantBool {
		t.Fatalf("got bool %v, want %v", gotBool, wantBool)
	}

	if !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Fatalf("got slice %v, want %v", gotSlice, wantSlice)
	}
}

func TestIntersectEmptySlices(t *testing.T) {
	a := []int{}
	b := []int{}

	gotBool, gotSlice := intersect(a, b)

	wantBool := false
	wantSlice := []int{}

	if gotBool != wantBool {
		t.Fatalf("got bool %v, want %v", gotBool, wantBool)
	}

	if !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Fatalf("got slice %v, want %v", gotSlice, wantSlice)
	}
}

func TestIntersectUniqueResult(t *testing.T) {
	a := []int{1, 2, 2, 3, 3}
	b := []int{2, 2, 3, 4}

	gotBool, gotSlice := intersect(a, b)

	wantBool := true
	wantSlice := []int{2, 3}

	if gotBool != wantBool {
		t.Fatalf("got bool %v, want %v", gotBool, wantBool)
	}

	if !reflect.DeepEqual(gotSlice, wantSlice) {
		t.Fatalf("got slice %v, want %v", gotSlice, wantSlice)
	}
}
