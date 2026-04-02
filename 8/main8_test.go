package main

import (
	"testing"
	"time"

	"go-core-task/8/wg"
)

func TestWaitImmediate(t *testing.T) {
	myWG := wg.New()

	done := make(chan struct{})

	go func() {
		myWG.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait should not block when counter is 0")
	}
}

func TestWaitBlocks(t *testing.T) {
	myWG := wg.New()
	myWG.Add(1)

	done := make(chan struct{})

	go func() {
		myWG.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("Wait returned too early")
	case <-time.After(200 * time.Millisecond):
	}

	myWG.Done()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait did not unblock")
	}
}

func TestNegativeCounter(t *testing.T) {
	myWG := wg.New()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()

	myWG.Add(-1)
}
