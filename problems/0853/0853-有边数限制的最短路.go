package main

import "fmt"

const (
	N   = 510
	M   = 10010
	inf = 0x3f3f3f3f
)

var (
	n, m, k      int
	dist, backup [N]int
	edges        [M]edge
)

type edge struct {
	a, b, w int
}

func bellmanFord() int {
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[1] = 0

	for i := 0; i < k; i++ {
		copy(backup[:], dist[:])
		for j := 0; j < m; j++ {
			a := edges[j].a
			b := edges[j].b
			w := edges[j].w
			if dist[b] > backup[a]+w {
				dist[b] = backup[a] + w
			}
		}
	}
	return dist[n]
}

func main() {
	fmt.Scan(&n, &m, &k)
	for i := 0; i < m; i++ {
		var e edge
		fmt.Scan(&e.a, &e.b, &e.w)
		edges[i] = e
	}
	if w := bellmanFord(); w > inf/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(w)
	}
}
