package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"testing"
)

func TestInsertSaltEven(t *testing.T) {
	got := string(insertSalt([]rune("abcd"), "go-2024"))
	want := "abgo-2024cd"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestInsertSaltOdd(t *testing.T) {
	got := string(insertSalt([]rune("abcde"), "go-2024"))
	want := "abgo-2024cde"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestInsertSaltEmpty(t *testing.T) {
	got := string(insertSalt([]rune(""), "go-2024"))
	want := "go-2024"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestTypes(t *testing.T) {
	v := reflect.ValueOf(vars{
		NumDecimal:     42,
		NumOctal:       052,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	})

	got := make([]string, 0, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		got = append(got, fmt.Sprintf("%T", v.Field(i).Interface()))
	}

	want := []string{
		"int",
		"int",
		"int",
		"float64",
		"string",
		"bool",
		"complex64",
	}

	if len(got) != len(want) {
		t.Fatalf("len mismatch: got %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("at index %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestJoinedString(t *testing.T) {
	v := reflect.ValueOf(vars{
		NumDecimal:     42,
		NumOctal:       052,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	})

	got := ""
	for i := 0; i < v.NumField(); i++ {
		got += fmt.Sprint(v.Field(i).Interface())
	}

	want := "4242423.14Golangtrue(1+2i)"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestFinalHash(t *testing.T) {
	joined := "4242423.14Golangtrue(1+2i)"
	runes := []rune(joined)
	salted := insertSalt(runes, "go-2024")

	sum := sha256.Sum256([]byte(string(salted)))
	got := fmt.Sprintf("%x", sum)

	want := "53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
