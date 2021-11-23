package main

import (
	"fmt"
	"sort"

	"github.com/luxcgo/go-acwing/internal/oj"
)

const N = 3e5 + 10

var a, s [N]int

type pair struct { // 定义操作
	p1 int
	p2 int
}

func main() {
	in := oj.NewScanner()
	n, m := in.ReadInt(), in.ReadInt()
	var alls []int
	add := make([]pair, 0)
	for i := 0; i < n; i++ {
		x, c := in.ReadInt(), in.ReadInt()
		add = append(add, pair{x, c})
		alls = append(alls, x)
	}

	query := make([]pair, 0)
	for i := 0; i < m; i++ {
		l, r := in.ReadInt(), in.ReadInt()
		query = append(query, pair{l, r})
		alls = append(alls, l)
		alls = append(alls, r)
	}

	sort.Ints(alls)
	alls = oj.Unique(alls)

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
