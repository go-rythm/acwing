package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

const N = 310

var (
	n, m int
	h, f [N][N]int
	dx   = [4]int{-1, 0, 1, 0}
	dy   = [4]int{0, 1, 0, -1}
)

func dp(x, y int) int {
	if f[x][y] != -1 {
		return f[x][y]
	}
	f[x][y] = 1
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 1 && a <= n && b >= 1 && b <= m && h[a][b] < h[x][y] {
			f[x][y] = oj.Max(f[x][y], dp(a, b)+1)
		}
	}
	return f[x][y]
}

func main() {
	s := oj.NewScanner()
	n, m = s.ReadInt(), s.ReadInt()
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			h[i][j] = s.ReadInt()
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = -1
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			res = oj.Max(res, dp(i, j))
		}
	}

	fmt.Println(res)
}
