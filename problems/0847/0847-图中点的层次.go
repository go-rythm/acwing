package main

import "fmt"

const N = 100010

var (
	h, e, ne, d, q [N]int
	idx            int
)

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func bfs(n int) int {
	hh, tt := 0, 0
	// 0号节点是编号为1的节点
	q[0] = 1
	// 初始化存储每个节点离起点的距离
	for i := 0; i < len(d); i++ {
		d[i] = -1
	}
	d[1] = 0

	for hh <= tt {
		t := q[hh]
		hh++
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if d[j] == -1 {
				d[j] = d[t] + 1
				tt++
				q[tt] = j
			}
		}
	}
	return d[n]
}

func main() {
	// 初始化模拟队列
	for i := 0; i < len(q); i++ {
		h[i] = -1
	}
	var n, m int
	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		add(a, b)
	}
	res := bfs(n)
	fmt.Println(res)
}
