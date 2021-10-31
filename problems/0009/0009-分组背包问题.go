package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

const N = 110

var (
	n, m int
	v, w [N][N]int
	s, g [N]int
)

func main() {
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
		for j := 0; j < s[i]; j++ {
			fmt.Scan(&v[i][j], &w[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := m; j >= 0; j-- {
			for k := 0; k < s[i]; k++ {
				if v[i][k] <= j {
					g[j] = oj.Max(g[j], g[j-v[i][k]]+w[i][k])
				}
			}
		}
	}

	fmt.Println(g[m])
}
