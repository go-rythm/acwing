package main

import (
	"fmt"
	"strings"
)

const N = 100010

var stk [N]int

// top表示栈顶
var top int

// 向栈顶插入一个数
func push(x int) {
	top++
	stk[top] = x
}

// 从栈顶弹出一个数
func pop() {
	top--
}

// 判断栈是否为空
func isEmpty() bool {
	return top == 0
}

// 栈顶的值
func getTop() int {
	return stk[top]
}

func main() {
	var sb strings.Builder
	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var op string
		var x int
		fmt.Scan(&op)
		switch op {
		case "push":
			fmt.Scan(&x)
			push(x)
		case "pop":
			pop()
		case "empty":
			if isEmpty() {
				sb.WriteString("YES\n")
			} else {
				sb.WriteString("NO\n")
			}
		case "query":
			fmt.Fprintf(&sb, "%d\n", getTop())
		}
	}

	fmt.Print(sb.String())
}
