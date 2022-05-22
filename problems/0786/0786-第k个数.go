package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

var q []int

func main() {
	s := oj.NewScanner()
	n, k := s.ReadInt(), s.ReadInt()
	q = s.ReadInts(n)
	res := quickSort(0, n-1, k)
	fmt.Println(res)
}

func quickSort(l, r, k int) int {
	if l >= r {
		return q[l]
	}
	x := q[l]
	i := l - 1
	j := r + 1
	for i < j {
		i++
		for q[i] < x {
			i++
		}
		j--
		for q[j] > x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	sl := j - l + 1
	if k <= sl {
		return quickSort(l, j, k)
	}
	return quickSort(j+1, r, k-sl)
}
