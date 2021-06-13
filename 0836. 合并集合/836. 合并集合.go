package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	new(os.Stdin)
	var n, m int
	fmt.Scan(&n, &m)
	// 初始化，根据节点编号是 1~n
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	var sb strings.Builder
	for i := 0; i < m; i++ {
		v := readLine()
		op := v[0]
		a, _ := strconv.Atoi(v[1])
		b, _ := strconv.Atoi(v[2])
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

var scanner *bufio.Scanner

func new(reader io.Reader) {
	scanner = bufio.NewScanner(reader)
	bs := make([]byte, 20000*1024)
	scanner.Buffer(bs, len(bs))
}
func readLine() []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}
