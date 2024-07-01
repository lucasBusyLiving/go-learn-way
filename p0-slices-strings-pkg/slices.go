package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func intSort() {
	ints := []int{1, 1, 1, 2, 3, 4, 5, 5, 5, 7, 8, 9, 10, 10}
	maxInt := slices.Max(ints)
	fmt.Println(maxInt)

	slices.Reverse(ints)
	fmt.Println(ints)

	slices.Sort(ints)
	fmt.Println(ints)

	fmt.Println(len(ints))
	// 存在为最低下标，不存在为插入地方
	j := sort.SearchInts(ints, 10)
	fmt.Println(j)

	fmt.Println(math.MaxInt32, math.MinInt32)
}

func sortSlice() {
	intervals := [][]int{{1, 2}, {3, 4}, {2, 6}, {7, 8}}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	fmt.Println(intervals)

	type stu struct {
		Name string
		Age  int
	}

	stus := []stu{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 20}, {Name: "AA", Age: 10}}
	sort.Slice(stus, func(i, j int) bool { return stus[i].Age < stus[j].Age })
	fmt.Println(stus)
	sort.Slice(stus, func(i, j int) bool { return stus[i].Name < stus[j].Name })
	fmt.Println(stus)
}

func main() {
	sortSlice()
}
