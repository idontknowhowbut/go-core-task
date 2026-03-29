package main

import "fmt"

type StringIntMap map[string]int

func main() {
	m := make(StringIntMap)

	fmt.Println("Empty map:")
	fmt.Println(m)

	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	fmt.Println("Elements added:")
	fmt.Println(m)

	fmt.Println(`Exists("two"):`, m.Exists("two"))
	fmt.Println(`Exists("ten"):`, m.Exists("ten"))

	value, ok := m.Get("three")
	fmt.Println(`Get("three"):`, value, ok)

	value, ok = m.Get("ten")
	fmt.Println(`Get("ten"):`, value, ok)

	copied := m.Copy()
	fmt.Println("Copied map:")
	fmt.Println(copied)

	m.Remove("two")
	fmt.Println(`After Remove("two"):`)
	fmt.Println("Original map:", m)
	fmt.Println("Copied map:", copied)
}

func (m StringIntMap) Add(key string, value int) {
	m[key] = value
}

func (m StringIntMap) Remove(key string) {
	delete(m, key)

}

func (m StringIntMap) Copy() map[string]int {
	res := make(map[string]int)
	for k, v := range m {
		res[k] = v
	}
	return res
}

func (m StringIntMap) Exists(key string) bool {
	_, ok := m[key]
	return ok
}

func (m StringIntMap) Get(key string) (int, bool) {
	v, ok := m[key]
	return v, ok
}
