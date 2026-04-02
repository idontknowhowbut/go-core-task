package main

import (
	"testing"
	"time"
)

func TestRandomIntGeneratorInvalidInput(t *testing.T) {
	ch, err := randomIntGenerator(10, 10, make(chan struct{}))
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if ch != nil {
		t.Fatal("expected nil channel on invalid input")
	}
}

func TestRandomIntGeneratorRange(t *testing.T) {
	done := make(chan struct{})
	ch, err := randomIntGenerator(-1000, 101, done)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer close(done)

	for range 100 {
		n := <-ch
		if n < -1000 || n >= 101 {
			t.Fatalf("number %d is out of range", n)
		}
	}
}

func TestRandomIntGeneratorStopsOnDone(t *testing.T) {
	done := make(chan struct{})
	ch, err := randomIntGenerator(0, 10, done)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for range 5 {
		<-ch
	}

	close(done)

	select {
	case _, ok := <-ch:
		if ok {
			t.Fatal("expected channel to be closed")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timeout waiting for channel to close")
	}
}
