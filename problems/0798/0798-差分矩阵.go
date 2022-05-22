package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

const N = 1010

var (
	a, b [N][N]int
)

func main() {
	s := oj.NewScanner()
	n, m, q := s.ReadInt(), s.ReadInt(), s.ReadInt()
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] = s.ReadInt()
			insert(i, j, i, j, a[i][j])
		}
	}
	for i := 0; i < q; i++ {
		op := s.ReadInts(5)
		insert(op[0], op[1], op[2], op[3], op[4])
	}
	// 对生成的差分矩阵求二维前缀和
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] = a[i-1][j] + a[i][j-1] - a[i-1][j-1] + b[i][j]
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}

func insert(x1, y1, x2, y2, c int) {
	b[x1][y1] += c
	b[x1][y2+1] -= c
	b[x2+1][y1] -= c
	b[x2+1][y2+1] += c
}
