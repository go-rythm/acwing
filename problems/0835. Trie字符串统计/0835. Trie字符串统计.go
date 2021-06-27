package main

import "fmt"

const N = 100010

var (
	son [N][26]int
	idx int
	cnt [N]int
)

func insert(s string) {
	p := 0
	for _, c := range s {
		u := int(c - 'a')
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	cnt[p]++
}

func query(s string) int {
	p := 0
	for _, c := range s {
		u := int(c - 'a')
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	return cnt[p]
}

func main() {
	var n int
	fmt.Scan(&n)
	var res []int
	for i := 0; i < n; i++ {
		var op string
		var s string
		fmt.Scan(&op)
		fmt.Scan(&s)
		switch op {
		case "I":
			insert(s)
		case "Q":
			res = append(res, query((s)))
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
