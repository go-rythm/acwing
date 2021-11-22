package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

func main() {
	s := oj.NewScanner()
	n, m, x := s.ReadInt(), s.ReadInt(), s.ReadInt()
	a := s.ReadInts(n)
	b := s.ReadInts(m)
	for i, j := 0, m-1; i < n; i++ {
		for j > 0 && a[i]+b[j] > x {
			j--
		}
		if a[i]+b[j] == x {
			fmt.Printf("%d %d\n", i, j)
			break
		}
	}
}
