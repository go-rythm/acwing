package main

import "fmt"

const (
	N   = 510
	inf = 0x3f3f3f3f
)

var (
	g    [N][N]int
	dist [N]int
	st   [N]bool

	n, m int
)

func dijkstra() int {
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[1] = 0

	for i := 0; i < n; i++ {
		t := -1

		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		st[t] = true

		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], dist[t]+g[t][j])
		}
	}

	if dist[n] == inf {
		return -1
	}
	return dist[n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			g[i][j] = inf
		}
	}

	var x, y, z int
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y, &z)
		g[x][y] = min(g[x][y], z)
	}

	fmt.Println(dijkstra())
}
