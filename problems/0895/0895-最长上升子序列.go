package main

import (
	"fmt"
)

const N = 1010

func main() {
	var (
		n    int
		a, f [N]int
	)

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 1; i <= n; i++ {
		f[i] = 1 // 只有a[i]一个数
		for j := 1; j < i; j++ {
			if a[j] < a[i] {
				f[i] = Max(f[i], f[j]+1)
			}
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		res = Max(f[i], res)
	}
	fmt.Println(res)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
