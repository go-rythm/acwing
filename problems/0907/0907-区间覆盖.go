package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/luxcgo/go-acwing/internal/oj"
)

type pair struct {
	l int
	r int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var st, ed, n int
	fmt.Fscan(in, &st, &ed, &n)
	pairs := make([]*pair, n)
	for i := 0; i < n; i++ {
		p := &pair{}
		fmt.Fscan(in, &p.l, &p.r)
		pairs[i] = p
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].l < pairs[j].l
	})
	res := 0
	success := false
	for i := 0; i < n; i++ {
		r := int(-2e9)
		j := i
		for ; j < n && pairs[j].l <= st; j++ {
			r = oj.Max(r, pairs[j].r)
		}
		if r < st {
			break
		}
		res++
		if r >= ed {
			success = true
			break
		}
		st = r
		i = j - 1
	}
	if !success {
		res = -1
	}
	fmt.Println(res)
}
