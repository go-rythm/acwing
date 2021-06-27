package main

import (
	"fmt"
	"strings"
)

const N = 100010

var p, cnt [N]int

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		p[i] = i
		cnt[i] = 1
	}
	var op string
	var a, b int
	var sb strings.Builder
	for i := 0; i < m; i++ {
		fmt.Scan(&op)
		switch op {
		case "C":
			fmt.Scan(&a, &b)
			if find(a) != find(b) {
				sum := cnt[find(a)] + cnt[find(b)]
				cnt[find(a)] = sum
				cnt[find(b)] = sum
				p[find(a)] = find(b)
			}
		case "Q1":
			fmt.Scan(&a, &b)
			if sb.String() != "" {
				sb.WriteString("\n")
			}
			if find(a) == find(b) {
				sb.WriteString("Yes")
			} else {
				sb.WriteString("No")
			}
		case "Q2":
			fmt.Scan(&a)
			c := cnt[find(a)]
			if sb.String() != "" {
				sb.WriteString("\n")
			}
			sb.WriteString(fmt.Sprintf("%d", c))
		}
	}
	fmt.Println(sb.String())
}
