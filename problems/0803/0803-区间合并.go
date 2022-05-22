package main

import (
	"fmt"
	"sort"

	"github.com/go-rythm/acwing/internal/oj"
)

type pair struct {
	left  int
	right int
}

func main() {
	s := oj.NewScanner()
	n := s.ReadInt()
	intervals := make([]pair, n)
	for i := 0; i < n; i++ {
		intervals[i].left, intervals[i].right = s.ReadInt(), s.ReadInt()
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
		if ed < v.left {
			if st != -2e9 {
				merged = append(merged, pair{st, ed})
			}
			st = v.left
			ed = v.right
		} else {
			ed = oj.Max(v.right, ed)
		}
	}
	if st != -2e9 {
		merged = append(merged, pair{st, ed})
	}
	fmt.Println(len(merged))
}
