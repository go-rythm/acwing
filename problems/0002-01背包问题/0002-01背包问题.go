package main

import (
	"fmt"
)

const N = 1010

var (
	n, m int
	v, w [N]int

	// bag1 二维数组的方法
	f [N][N]int

	// bag2 一维数组的方法
	g [N]int
)

func bag1() {
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if v[i] <= j {
				if f[i-1][j-v[i]]+w[i] > f[i][j] {
					f[i][j] = f[i-1][j-v[i]] + w[i]
				}
			}
		}
	}
	fmt.Println(f[n][m])
}

func bag2() {
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			if g[j-v[i]]+w[i] > g[j] {
				g[j] = g[j-v[i]] + w[i]
			}
		}
	}
	fmt.Println(g[m])
}

func main() {
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i])
	}
	bag2()
}
