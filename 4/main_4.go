package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	sliceRes := SliceDiff(slice1, slice2)
	fmt.Println(sliceRes)
}

func SliceDiff(sliceA []string, sliceB []string) []string {
	m := make(map[string]struct{})

	for _, s := range sliceB {
		m[s] = struct{}{}
	}

	res := make([]string, 0, len(sliceA))
	for _, s := range sliceA {
		if _, ok := m[s]; !ok {
			res = append(res, s)
		}
	}
	return res
}
