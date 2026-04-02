package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for _, v := range []int{1, 2, 3} {
			ch1 <- v
		}
	}()

	go func() {
		defer close(ch2)
		for _, v := range []int{10, 20, 30} {
			ch2 <- v
		}
	}()

	go func() {
		defer close(ch3)
		for _, v := range []int{100, 200, 300} {
			ch3 <- v
		}
	}()

	for v := range merge(ch1, ch2, ch3) {
		fmt.Println(v)
	}
}

func merge[T any](chans ...<-chan T) <-chan T {
	out := make(chan T)
	wg := new(sync.WaitGroup)
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(ch <-chan T) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)

	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
