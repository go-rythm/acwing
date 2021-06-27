package main

import "fmt"

const N = 100010

// head为头指针
// idx指针指向当前可以使用的节点
var idx, head int

// e数组用于存储当前下标指针指向节点的值
// ne数组用于存储当前下标指针指向节点的next指针
var e, ne [N]int

// 初始化链表
func init() {
	head = -1
	idx = 0
}

// 向头指针指向位置插入节点
func addToHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

// 向第k+1个插入的节点后插入节点
func add(k, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

// 删除第k+1个插入的节点后面的节点
func remove(k int) {
	ne[k] = ne[ne[k]]
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)
		switch op {
		case "H":
			var x int
			fmt.Scan(&x)
			addToHead(x)
		case "I":
			var k, x int
			fmt.Scan(&k, &x)
			add(k-1, x)
		case "D":
			var k int
			fmt.Scan(&k)
			if k == 0 {
				head = ne[head]
			} else {
				remove(k - 1)
			}
		}
	}

	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}
