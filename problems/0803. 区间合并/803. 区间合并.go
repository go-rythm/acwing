package main

import (
	"fmt"
	"sort"
)

type pair struct {
	left  int
	right int
}

func main() {
	var n int
	fmt.Scan(&n)
	intervals := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&intervals[i].left, &intervals[i].right)
	}
	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i].left == intervals[j].left {
			return intervals[i].right < intervals[j].right
		}
		return intervals[i].left < intervals[j].left
	})
	var st, ed int = -2e9, -2e9
	merged := make([]pair, 0)
	for _, v := range intervals {
		if ed == -2e9 {
			st = v.left
			ed = v.right
		} else {
			if ed < v.left {
				merged = append(merged, pair{st, ed})
				st = v.left
				ed = v.right
			} else {
				ed = v.right
			}
		}
	}
	if st != -2e9 {
		merged = append(merged, pair{st, ed})
	}
	fmt.Println(len(merged))
}
