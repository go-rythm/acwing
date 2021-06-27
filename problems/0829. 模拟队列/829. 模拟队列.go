package main

import (
	"fmt"
	"strings"
)

const N = 100010

var q [N]int

// head 表示队头，tail表示队尾
var head = 0
var tail = -1

// 向队尾插入一个数
func push(x int) {
	tail++
	q[tail] = x
}

// 从队头弹出一个数
func pop() {
	head++
}

// 判断队列是否为空
func isEmpty() bool {
	return head > tail
}

// 队头的值
func peek() int {
	return q[head]
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
			fmt.Fprintf(&sb, "%d\n", peek())
		}
	}

	fmt.Print(sb.String())
}
