package main

import (
	"fmt"
	"strings"

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
	mergeSort(q, 0, n-1)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(q)), " "), "[]"))
}

func mergeSort(q []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			i++
		} else {
			temp[k] = q[j]
			j++
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
}
