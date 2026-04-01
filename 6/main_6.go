package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {
	done := make(chan struct{})
	ch, err := randomIntGenerator(-1000, 101, done)
	if err != nil {
		fmt.Println(err)
		return
	}

	for range 1000 {
		fmt.Println(<-ch)
	}
	close(done)

}

func randomIntGenerator(min int, max int, done <-chan struct{}) (<-chan int, error) {
	if min >= max {
		return nil, errors.New("incorrect input")
	}

	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- (rand.IntN(max-min) + min):
			}
		}
	}()

	return ch, nil
}
