package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

const N = int(1e5 + 10)

var a, b [N]int

func main() {
	s := oj.NewScanner()
	n, m := s.ReadInt(), s.ReadInt()
	for i := 1; i <= n; i++ {
		a[i] = s.ReadInt()
		b[i] = a[i] - a[i-1]
	}
	for i := 0; i < m; i++ {
		l, r, c := s.ReadInt(), s.ReadInt(), s.ReadInt()
		b[l] += c
		b[r+1] -= c
	}
	for i := 1; i <= n; i++ {
		a[i] = b[i] + a[i-1]
		fmt.Printf("%d ", a[i])
	}

}
