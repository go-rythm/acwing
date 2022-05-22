package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

func main() {
	s := oj.NewScanner()
	n, m := s.ReadInt(), s.ReadInt()
	q := s.ReadInts(n)
	for i := 0; i < m; i++ {
		x := s.ReadInt()

		var leftBound int
		var rightBound int
		l := 0
		r := n - 1
		for l+1 < r {
			mid := (l + r) >> 1
			if q[mid] >= x {
				r = mid
			} else {
				l = mid
			}
		}
		if q[l] == x {
			leftBound = l
		} else if q[r] == x {
			leftBound = r
		} else {
			fmt.Println("-1 -1")
			continue
		}

		l = 0
		r = n - 1
		for l+1 < r {
			mid := (l + r) >> 1
			if q[mid] > x {
				r = mid
			} else {
				l = mid
			}
		}
		if q[r] == x {
			rightBound = r
		} else if q[l] == x {
			rightBound = l
		} else {
			fmt.Println("-1 -1")
			continue
		}

		fmt.Printf("%d %d\n", leftBound, rightBound)
	}
}
