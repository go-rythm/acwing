package main

import "fmt"

const N = 1010

var (
	n, m int
	v, w [N]int

	f [N][N]int

	g [N]int
)

func bag1() {
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k*v[i] <= j; k++ {
				if f[i-1][j-v[i]*k]+w[i]*k > f[i][j] {
					f[i][j] = f[i-1][j-v[i]*k] + w[i]*k
				}
			}
		}
	}
	fmt.Println(f[n][m])
}

func bag2() {
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j-v[i] >= 0 && f[i][j-v[i]]+w[i] > f[i][j] {
				f[i][j] = f[i][j-v[i]] + w[i]
			}
		}
	}
	fmt.Println(f[n][m])
}

func bag3() {
	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
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
