package main

import (
	"fmt"
)

const N = 1e5 + 10

var n, m int
var arr, sum [N]int

func main() {
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + arr[i]
	}

	res := make([]int, m)
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		res[i] = sum[r] - sum[l-1]
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
