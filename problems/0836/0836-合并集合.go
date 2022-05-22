package main

import (
	"fmt"
	"strings"

	"github.com/go-rythm/acwing/internal/oj"
)

const N = 100010

// 存储每个点的祖宗节点
var p [N]int

// 返回x的祖宗节点
func find(x int) int {
	if p[x] != x {
		// 路径压缩优化
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	s := oj.NewScanner()
	n, m := s.ReadInt(), s.ReadInt()
	// 初始化，根据节点编号是 1~n
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	var sb strings.Builder
	for i := 0; i < m; i++ {
		op := s.ReadString()
		a, b := s.ReadInt(), s.ReadInt()
		switch op {
		case "M":
			// 合并 a 和 b 所在的两个集合：
			p[find(a)] = find(b)
		case "Q":
			if sb.String() != "" {
				fmt.Fprintln(&sb)
			}
			if find(a) == find(b) {
				fmt.Fprint(&sb, "Yes")
			} else {
				fmt.Fprint(&sb, "No")
			}
		}
	}
	fmt.Println(sb.String())
}
