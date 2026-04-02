package main

import "fmt"

func main() {
	ch1 := generate(100)
	ch2 := transform(ch1)
	for v := range ch2 {
		fmt.Println(v)
	}

}

func generate(limit uint8) chan uint8 {
	ch := make(chan uint8)
	go func() {
		for i := range limit {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func transform(in chan uint8) chan float64 {
	ch := make(chan float64)
	go func() {
		for v := range in {
			t := float64(v)
			ch <- t * t
		}
		close(ch)
	}()
	return ch
}
