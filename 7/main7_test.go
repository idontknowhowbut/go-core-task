package main

import (
	"slices"
	"testing"
)

func makeChan(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func collect[T any](ch <-chan T) []T {
	var res []T
	for v := range ch {
		res = append(res, v)
	}
	return res
}

func TestMergeBasic(t *testing.T) {
	ch1 := makeChan([]int{1, 2, 3})
	ch2 := makeChan([]int{10, 20})
	ch3 := makeChan([]int{100})

	got := collect(merge(ch1, ch2, ch3))
	want := []int{1, 2, 3, 10, 20, 100}

	slices.Sort(got)
	slices.Sort(want)

	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMergeEmptyChannels(t *testing.T) {
	ch1 := makeChan([]int{})
	ch2 := makeChan([]int{})

	got := collect(merge(ch1, ch2))
	want := []int{}

	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMergeSingleChannel(t *testing.T) {
	ch1 := makeChan([]int{5, 6, 7})

	got := collect(merge(ch1))
	want := []int{5, 6, 7}

	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMergeNoChannels(t *testing.T) {
	got := collect(merge[int]())
	want := []int{}

	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
