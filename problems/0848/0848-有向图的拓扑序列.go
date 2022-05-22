package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

const N = int(1e5 + 10)

var (
	n, m     int
	h, e, ne [N]int
	idx      int
	q, d     [N]int
)

func main() {
	s := oj.NewScanner()
	n, m = s.ReadInt(), s.ReadInt()
	for i := 0; i < len(h); i++ {
		h[i] = -1
	}
	for i := 0; i < m; i++ {
		a, b := s.ReadInt(), s.ReadInt()
		add(a, b)
		d[b]++
	}
	if topSort() {
		for i := 0; i < n; i++ {
			fmt.Print(q[i], " ")
		}
	} else {
		fmt.Println(-1)
	}
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func topSort() bool {
	hh := 0
	tt := -1
	for i := 1; i <= n; i++ {
		if d[i] == 0 {
			tt++
			q[tt] = i
		}
	}
	for hh <= tt {
		t := q[hh]
		hh++
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			d[j]--
			if d[j] == 0 {
				tt++
				q[tt] = j
			}
		}
	}
	return tt == n-1
}
