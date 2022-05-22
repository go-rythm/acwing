package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

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
			dist[j] = oj.Min(dist[j], dist[t]+g[t][j])
		}
	}

	if dist[n] == inf {
		return -1
	}
	return dist[n]
}

func main() {
	s := oj.NewScanner()
	n, m = s.ReadInt(), s.ReadInt()
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			g[i][j] = inf
		}
	}

	for i := 0; i < m; i++ {
		x, y, z := s.ReadInt(), s.ReadInt(), s.ReadInt()
		g[x][y] = oj.Min(g[x][y], z)
	}

	fmt.Println(dijkstra())
}
