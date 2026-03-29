package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	s := make([]int, 10)

	for i := range s {
		s[i] = rand.IntN(100)
	}

	fmt.Println(s)

	s = sliceExample(s)

	fmt.Println(s)

	m := copySlice(s)
	m = addElements(m, 999)
	fmt.Println(s)
	fmt.Println(m)

	m = removeElement(m, 50)
	fmt.Println(m)
}

func sliceExample(s []int) []int {
	var res []int

	for _, n := range s {
		if n%2 == 0 {
			res = append(res, n)
		}
	}

	return res
}

func addElements(s []int, n int) []int {
	return append(s, n)
}

func copySlice(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}

func removeElement(s []int, i int) []int {
	if i > len(s)-1 || i < 0 {
		return []int{}
	}
	var res []int
	res = append(res, s[:i]...)
	res = append(res, s[i+1:]...)
	return res
}
