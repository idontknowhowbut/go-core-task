package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	limit := uint8(5)
	ch := generate(limit)

	var result []uint8
	for v := range ch {
		result = append(result, v)
	}

	expected := []uint8{0, 1, 2, 3, 4} // генерируем от 0 до limit-1
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("generate(%d) = %v; want %v", limit, result, expected)
	}
}

func TestTransform(t *testing.T) {
	input := make(chan uint8, 3)
	input <- 1
	input <- 2
	input <- 3
	close(input)

	ch := transform(input)

	var result []float64
	for v := range ch {
		result = append(result, v)
	}

	expected := []float64{1, 4, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("transform(input) = %v; want %v", result, expected)
	}
}

func TestPipeline(t *testing.T) {
	ch1 := generate(4) // 0,1,2,3
	ch2 := transform(ch1)

	var result []float64
	for v := range ch2 {
		result = append(result, v)
	}

	expected := []float64{0, 1, 4, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("pipeline result = %v; want %v", result, expected)
	}
}
