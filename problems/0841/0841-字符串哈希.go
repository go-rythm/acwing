package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

const (
	N = int(1e5 + 10)
	P = 131
)

var (
	// 前缀和
	h [N]int64
	// 进制
	p [N]int64
)

func main() {
	s := oj.NewScanner()
	n, m := s.ReadInt(), s.ReadInt()
	x := " " + s.ReadString()
	p[0] = 1 // 0次方
	for i := 1; i <= n; i++ {
		p[i] = p[i-1] * P
		h[i] = h[i-1]*P + int64(x[i])
	}

	for i := 0; i < m; i++ {
		l1, r1, l2, r2 := s.ReadInt(), s.ReadInt(), s.ReadInt(), s.ReadInt()
		if get(l1, r1) == get(l2, r2) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func get(l, r int) int64 {
	return h[r] - h[l-1]*p[r-l+1]
}
