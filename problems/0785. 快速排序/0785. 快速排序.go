package main

import (
	"fmt"
	"strings"
)

func quickSort(q []int, l int, r int) {
	if l >= r {
		return
	}

	i, j := l-1, r+1
	x := q[(l+r)>>1]
	for i < j {
		i++
		for q[i] < x {
			i++
		}
		j--
		for q[j] > x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	// [l,j] 和 [j+1,r] 的 pivot(line14) 不能选择 q[r]，会有边界问题，导致死循环，选择数组中任意其他位置都可以。
	// 循环不变量:
	// [l,i] 区间若合法则其中所有值 <= pivot
	// [j,r] 区间若合法则其中所有值 >= pivot
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func main() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	quickSort(arr, 0, n-1)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), " "), "[]"))
}
