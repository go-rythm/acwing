package main

import "fmt"

const N = 10

var (
	n    int
	path [N]int
	st   [N]bool
)

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println()
		return
	}

	for i := 1; i <= n; i++ {
		if !st[i] {
			path[u] = i
			st[i] = true
			dfs(u + 1)
			st[i] = false
		}
	}
}

func main() {
	n = 3
	dfs(0)
}
