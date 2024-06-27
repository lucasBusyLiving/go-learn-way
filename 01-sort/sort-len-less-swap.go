package main

import (
	"fmt"
	"sort"
)

type foo struct {
	Name string
}

type fooSlice []foo

func (f fooSlice) Len() int           { return len(f) }
func (f fooSlice) Less(i, j int) bool { return f[i].Name < f[j].Name }
func (f fooSlice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	myFoos := []foo{{"Bob"}, {"Alice"}, {"CiCi"}, {"AA"}}
	sort.Sort(fooSlice(myFoos))
	fmt.Println(myFoos)

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sort.Ints(ints)
	fmt.Println(ints)

	reverseInts := sort.Reverse(sort.IntSlice(ints))
	sort.Sort(reverseInts)
	fmt.Println(reverseInts)
}
