package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

var (
	q, temp []int
)

func main() {
	s := oj.NewScanner()
	n := s.ReadInt()
	q = s.ReadInts(n)
	temp = make([]int, n)
	fmt.Println(mergeSort(q, 0, n-1))
}

func mergeSort(q []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) >> 1
	cnt := mergeSort(q, l, mid) + mergeSort(q, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			i++
		} else {
			temp[k] = q[j]
			j++
			cnt += mid - i + 1
		}
		k++
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		k++
		j++
	}
	copy(q[l:r+1], temp)
	return cnt
}
