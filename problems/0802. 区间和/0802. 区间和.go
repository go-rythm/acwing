package main

import (
	"fmt"
	"sort"
)

const N = 3e5 + 10

var n, m int
var a, s [N]int

type pair struct { // 定义操作
	p1 int
	p2 int
}

func main() {
	fmt.Scan(&n, &m)
	var alls []int
	add := make([]pair, 0)
	for i := 0; i < n; i++ {
		var x, c int
		fmt.Scan(&x, &c)
		add = append(add, pair{x, c})
		alls = append(alls, x)
	}

	query := make([]pair, 0)
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		query = append(query, pair{l, r})
		alls = append(alls, l)
		alls = append(alls, r)
	}

	sort.Ints(alls)
	alls = unique(alls)

	for _, v := range add {
		a[findIndex(alls, v.p1)] += v.p2
	}

	for i := 0; i < len(alls); i++ {
		s[i+1] = s[i] + a[i]
	}

	for _, v := range query {
		l := findIndex(alls, v.p1)
		r := findIndex(alls, v.p2)
		fmt.Println(s[r+1] - s[l])
	}
}

func unique(a []int) []int {
	j := 0
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] != a[i-1] {
			a[j] = a[i]
			j++
		}
	}
	return a[:j]
}

func findIndex(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l+1 < r {
		mid := (l + r) >> 1
		if arr[mid] < target {
			l = mid
		} else if arr[mid] > target {
			r = mid
		} else {
			return mid
		}
	}
	if arr[l] == target {
		return l
	}
	if arr[r] == target {
		return r
	}
	return -1
}
