package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

const N = int(1e5 + 3)

var (
	h, e, ne [N]int
	idx      int
)

func main() {
	s := oj.NewScanner()
	n := s.ReadInt()

	for i := 0; i < len(h); i++ {
		h[i] = -1
	}

	for i := 0; i < n; i++ {
		op := s.ReadString()
		x := s.ReadInt()
		switch op {
		case "I":
			insert(x)
		case "Q":
			if find(x) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

func insert(x int) {
	k := (x%N + N) % N
	e[idx] = x
	ne[idx] = h[k]
	h[k] = idx
	idx++
}

func find(x int) bool {
	k := (x%N + N) % N
	for i := h[k]; i != -1; i = ne[i] {
		if e[i] == x {
			return true
		}
	}
	return false
}
