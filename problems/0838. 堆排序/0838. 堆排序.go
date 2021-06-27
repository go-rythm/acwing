package main

import "fmt"

const N = 100010

// h[N] 存储堆中的值, h[1] 是堆顶，
// x的左儿子是 2x, 右儿子是 2x + 1
var h [N]int

// size 存储堆大小
var size int

func down(u int) {
	t := u
	if u*2 <= size && h[u*2] < h[t] {
		t = u * 2
	}
	if u*2+1 <= size && h[u*2+1] < h[t] {
		t = u*2 + 1
	}
	if u != t {
		h[u], h[t] = h[t], h[u]
		down(t)
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	size = n
	// 存储输入，因为堆顶是 h[1]，所以从 h[1] 开始存储
	for i := 1; i <= n; i++ {
		fmt.Scan(&h[i])
	}

	// 建堆
	for i := n / 2; i > 0; i-- {
		down(i)
	}

	for i := 0; i < m; i++ {
		fmt.Print(h[1], " ")
		h[1] = h[size]
		size--
		down(1)
	}
}
