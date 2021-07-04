package main

import "fmt"

const (
	// 输入规模
	N = 100010
	// 邻接表存储无向无无环连通图（树），每个点都会存储两次
	M = 2 * N
)

var (
	// 存储n个邻接表的头指针
	h [N]int
	// 单链表的值
	e [M]int
	// 单链表的next指针
	ne [M]int
	// 单链表的可用指针
	idx int
	// 输入的树的结点数
	n int
	// 表示以每一个点为重心的剩余各个连通块中点数的最大值
	// 中的最小值
	ans = N
	// 记录节点是否被访问过，访问过则标记为true
	st [N]bool
)

// 以a为头结点的单链表中插入b
func add(a int, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func dfs(u int) int {
	// 表示以u为重心的剩余各个连通块中点数的最大值
	res := 0
	st[u] = true
	// 存储以u为根的树的节点数, 包括u
	sum := 1

	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if !st[j] {
			s := dfs(j)
			res = max(res, s)
			sum += s
		}
	}

	res = max(res, n-sum)
	ans = min(res, ans)
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	for i := 0; i < len(h); i++ {
		h[i] = -1
	}

	fmt.Scan(&n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		add(a, b)
		add(b, a)
	}

	// 可以任意选定一个节点开始 0< u <=n
	dfs(1)
	fmt.Println(ans)
}
