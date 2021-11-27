package main

import (
	"fmt"
	"strings"
)

const N = 100010

var h, hp, ph [N]int
var size int

func swap(a, b *int) {
	*a, *b = *b, *a
}

func heapSwap(u, v int) {
	swap(&hp[u], &hp[v])
	swap(&ph[hp[u]], &ph[hp[v]])
	swap(&h[u], &h[v])
}

func up(u int) {
	if u/2 > 0 && h[u/2] > h[u] {
		heapSwap(u/2, u)
		up(u / 2)
	}
}

func down(u int) {
	t := u
	if u*2 <= size && h[u*2] < h[t] {
		t = u * 2
	}
	if u*2+1 <= size && h[u*2+1] < h[t] {
		t = u*2 + 1
	}
	if u != t {
		heapSwap(u, t)
		down(t)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var op string
	var k, x int
	var cnt int
	var sb strings.Builder
	var isFirst = true

	for i := 0; i < n; i++ {
		fmt.Scan(&op)
		switch op {
		case "I":
			fmt.Scan(&x)
			size++
			h[size] = x
			cnt++
			// 第 cnt 个插入的数是堆中下标为 size 的值
			ph[cnt] = size
			// 堆中下标为 size 的值是第 cnt 个被插入的
			hp[size] = cnt
			up(size)
		case "PM":
			if isFirst {
				isFirst = false
			} else {
				sb.WriteString("\n")
			}
			sb.WriteString(fmt.Sprint(h[1]))
		case "DM":
			heapSwap(1, size)
			size--
			down(1)
		case "D":
			fmt.Scan(&k)
			u := ph[k]
			heapSwap(u, size)
			size--
			up(u)
			down(u)
		case "C":
			fmt.Scan(&k, &x)
			h[ph[k]] = x
			up(ph[k])
			down(ph[k])
		}
	}
	fmt.Print(sb.String())
}
