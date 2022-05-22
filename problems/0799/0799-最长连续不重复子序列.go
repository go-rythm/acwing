package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

const N = 100010

var (
	a []int
	s [N]int
)

func main() {
	in := oj.NewScanner()
	n := in.ReadInt()
	a = in.ReadInts(n)
	var max int
	for i, j := 0, 0; i < n; i++ {
		s[a[i]]++
		for s[a[i]] > 1 {
			s[a[j]]--
			j++
		}
		max = oj.Max(max, i-j+1)
	}
	fmt.Println(max)
}
