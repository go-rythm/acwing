package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

const N = 100010

var stack [N]int
var tt int

func main() {
	s := oj.NewScanner()
	n := s.ReadInt()
	res := make([]int, n)
	for i := 0; i < n; i++ {
		x := s.ReadInt()
		for tt > 0 && stack[tt] >= x {
			tt--
		}
		if tt == 0 {
			res[i] = -1
		} else {
			res[i] = stack[tt]
		}
		tt++
		stack[tt] = x
	}

	for _, v := range res {
		fmt.Printf("%d ", v)
	}
}
