package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Node struct{ Name string }
type nodes []Node

func (n nodes) Len() int            { return len(n) }
func (n nodes) Less(i, j int) bool  { return n[i].Name < n[j].Name }
func (n nodes) Swap(i, j int)       { n[i], n[j] = n[j], n[i] }
func (n *nodes) Push(x interface{}) { *n = append(*n, x.(Node)) }
func (n *nodes) Pop() interface{} {
	old := *n
	l := len(old)
	x := old[l-1]
	*n = (*n)[:l-1]
	return x
}

// 打印堆的树形结构
func printHeapTree(heap nodes, index, level int) {
	if index < len(heap) {
		printHeapTree(heap, 2*index+2, level+1)
		fmt.Println(strings.Repeat("  ", level), heap[index].Name)
		printHeapTree(heap, 2*index+1, level+1)
	}
}

func nodeHeap() {
	tNodes := make(nodes, 0)
	heap.Init(&tNodes)
	heap.Push(&tNodes, Node{"7"})
	heap.Push(&tNodes, Node{"6"})
	heap.Push(&tNodes, Node{"5"})
	heap.Push(&tNodes, Node{"4"})
	heap.Push(&tNodes, Node{"3"})
	heap.Push(&tNodes, Node{"2"})
	heap.Push(&tNodes, Node{"1"})
	fmt.Println(tNodes)

	// minHeap := &nodes{}
	// fmt.Println(len(*minHeap), cap(*minHeap))

	// printHeapTree(tNodes, 0, 0)
}

type h []int

func (a h) Len() int           { return len(a) }
func (a h) Less(i, j int) bool { return a[i] < a[j] }
func (a h) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *h) Push(v any)        { *a = append(*a, v.(int)) }
func (a *h) Pop() any {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[:n-1]
	return x
}

func main() {
	a := make(h, 0)
	heap.Push(&a, 1)
	heap.Push(&a, 9)
	heap.Push(&a, 3)
	heap.Push(&a, 7)
	fmt.Println(a)

	// nodeHeap()
}
