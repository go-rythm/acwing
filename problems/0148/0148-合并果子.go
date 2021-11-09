package main

import (
	"container/heap"
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	minHeap := &oj.MinHeap{}
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		heap.Push(minHeap, x)
	}
	res := 0
	for minHeap.Len() > 0 {
		if minHeap.Len() == 1 {
			break
		}
		a := heap.Pop(minHeap).(int)
		b := heap.Pop(minHeap).(int)
		sum := a + b
		res += sum
		heap.Push(minHeap, sum)
	}
	fmt.Println(res)
}
