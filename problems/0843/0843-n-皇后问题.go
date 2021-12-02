package main

import (
	"fmt"
)

const N = 20

var (
	n                 int
	g                 [N][N]string
	col, row, dg, udg [N]bool
)

func dfs1(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(g[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		if !col[i] && !dg[i+u] && !udg[i-u+n] {
			g[u][i] = "Q"
			col[i], dg[i+u], udg[i-u+n] = true, true, true
			dfs1(u + 1)
			col[i], dg[i+u], udg[i-u+n] = false, false, false
			g[u][i] = "."
		}
	}
}

func dfs2(x, y, s int) {
	if x == n {
		y++
		x = 0
	}
	if y == n {
		if s == n {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					fmt.Print(g[i][j])
				}
				fmt.Println()
			}
			fmt.Println()
		}
		return
	}

	dfs2(x+1, y, s)

	if !col[x] && !row[y] && !dg[x+y] && !udg[x-y+n] {
		g[x][y] = "Q"
		col[x], row[y], dg[x+y], udg[x-y+n] = true, true, true, true
		dfs2(x+1, y, s+1)
		col[x], row[y], dg[x+y], udg[x-y+n] = false, false, false, false
		g[x][y] = "."
	}
}

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = "."
		}
	}
	// dfs1(0)
	dfs2(0, 0, 0)
}
