package main

import (
	"fmt"
)

func main() {
	var (
		n, m int
		v, w []int
		g    []int
	)

	// in, err := os.Open("in.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer in.Close()
	// os.Stdin = in

	fmt.Scan(&n, &m)
	v = make([]int, 11*n+1)
	w = make([]int, 11*n+1)
	g = make([]int, m+1)

	index := 0
	for i := 1; i <= n; i++ {
		var vi, wi, si int
		fmt.Scan(&vi, &wi, &si)
		newItem := 1
		var total int
		for total = 1; total <= si; {
			index++
			v[index] = vi * newItem
			w[index] = wi * newItem
			newItem *= 2
			total += newItem
		}
		total -= newItem
		if total < si {
			index++
			v[index] = vi * (si - total)
			w[index] = wi * (si - total)
		}
	}
	for i := 1; i <= index; i++ {
		for j := m; j >= v[i]; j-- {
			if g[j-v[i]]+w[i] > g[j] {
				g[j] = g[j-v[i]] + w[i]
			}
		}
	}
	fmt.Println(g[m])
}
