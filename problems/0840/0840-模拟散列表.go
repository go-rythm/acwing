package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

// // 方法一：拉链法
// // 大于数据范围的第一个质数
// const N = int(1e5 + 3)

// var (
// 	h, e, ne [N]int
// 	idx      int
// )

// func main() {
// 	s := oj.NewScanner()
// 	n := s.ReadInt()

// 	for i := 0; i < len(h); i++ {
// 		h[i] = -1
// 	}

// 	for i := 0; i < n; i++ {
// 		op := s.ReadString()
// 		x := s.ReadInt()
// 		switch op {
// 		case "I":
// 			insert(x)
// 		case "Q":
// 			if find(x) {
// 				fmt.Println("Yes")
// 			} else {
// 				fmt.Println("No")
// 			}
// 		}
// 	}
// }

// func insert(x int) {
// 	k := (x%N + N) % N
// 	e[idx] = x
// 	ne[idx] = h[k]
// 	h[k] = idx
// 	idx++
// }

// func find(x int) bool {
// 	k := (x%N + N) % N
// 	for i := h[k]; i != -1; i = ne[i] {
// 		if e[i] == x {
// 			return true
// 		}
// 	}
// 	return false
// }

// 方法二：开放寻址法
const N = int(2e5 + 3)

var (
	h    [N]int
	null = 0x3f3f3f
)

func main() {
	s := oj.NewScanner()
	n := s.ReadInt()

	for i := 0; i < len(h); i++ {
		h[i] = null
	}

	for i := 0; i < n; i++ {
		op := s.ReadString()
		x := s.ReadInt()
		k := find(x)
		switch op {
		case "I":
			h[k] = x
		case "Q":
			if h[k] != null {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

func find(x int) int {
	k := (x%N + N) % N
	for h[k] != null && h[k] != x {
		k++
		if k == N {
			k = 0
		}
	}
	return k
}
