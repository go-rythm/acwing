package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"

	"github.com/go-rythm/acwing/internal/oj"
)

type pair struct {
	l int
	r int
}

func main() {

	// in := oj.In()
	// os.Stdin = in
	// defer in.Close()

	in := bufio.NewReader(os.Stdin)
	var cnt int
	fmt.Scan(&cnt)
	pairs := make([]*pair, cnt)
	for i := 0; i < cnt; i++ {
		p := &pair{}
		fmt.Fscan(in, &p.l, &p.r)
		pairs[i] = p
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].l < pairs[j].l
	})
	minHeap := &oj.MinHeap{}
	heap.Init(minHeap)
	for _, p := range pairs {
		if minHeap.Len() == 0 {
			heap.Push(minHeap, p.r)
			continue
		}
		top := heap.Pop(minHeap).(int)
		if minHeap.Len() == 0 || p.l <= top {
			heap.Push(minHeap, p.r)
			heap.Push(minHeap, top)
		} else {
			heap.Push(minHeap, p.r)
		}
	}
	fmt.Println(minHeap.Len())
}
