package main

import "fmt"

const N = 100010

// e 存储节点的值
// l 存储节点的左指针
// r 存储节点的右指针
var e, l, r [N]int
var idx = 0

func init() {
	r[0] = 1
	l[1] = 0
	idx = 2
}

// 在第 k 个点的右边插入 x
func insert(k, x int) {
	e[idx] = x
	l[idx] = k
	r[idx] = r[k]
	// 先修改第 k 个点右指针指向的节点的左指针
	// 再修改第 k 个点的右指针
	l[r[k]] = idx
	r[k] = idx
	idx++
}

// 删除第 k 个点
func remove(k int) {
	r[l[k]] = r[k]
	l[r[k]] = l[k]
}

func main() {
	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var op string
		var k int
		var x int
		fmt.Scan(&op)
		switch op {
		case "L":
			fmt.Scan(&x)
			insert(0, x)
		case "R":
			fmt.Scan(&x)
			insert(l[1], x)
		case "D":
			fmt.Scan(&k)
			remove(k + 1)
		case "IL":
			fmt.Scan(&k, &x)
			insert(l[k+1], x)
		case "IR":
			fmt.Scan(&k, &x)
			insert(k+1, x)
		}
	}
	for i := r[0]; i != 1; i = r[i] {
		fmt.Printf("%d ", e[i])
	}
}
