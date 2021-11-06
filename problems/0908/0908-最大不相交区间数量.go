package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	l int
	r int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var cnt int
	fmt.Scan(&cnt)
	pairs := make([]*pair, cnt)
	for i := 0; i < cnt; i++ {
		p := &pair{}
		fmt.Fscan(in, &p.l, &p.r)
		pairs[i] = p
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].r < pairs[j].r
	})
	r := int(-2e9)
	res := 0
	for _, v := range pairs {
		if v.l > r {
			r = v.r
			res++
		}
	}

	fmt.Println(res)
}
