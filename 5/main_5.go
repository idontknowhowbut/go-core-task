package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	resBool, resSlice := intersect(a, b)

	fmt.Println(resBool)
	fmt.Println(resSlice)

}

func intersect(sliceA []int, sliceB []int) (bool, []int) {
	m := make(map[int]struct{})

	for _, v := range sliceB {
		m[v] = struct{}{}
	}

	doIntersect := false
	res := make([]int, 0, len(sliceA))

	for _, v := range sliceA {
		if _, ok := m[v]; ok {
			doIntersect = true
			res = append(res, v)
			delete(m, v)
		}
	}

	return doIntersect, res
}
