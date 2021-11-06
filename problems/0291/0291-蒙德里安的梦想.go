package main

import (
	"fmt"
)

const (
	N = 12
	M = 1 << N
)

func main() {
	var (
		st    [M]bool
		state = map[int][]int{}
		n, m  int
	)

	// in := oj.In()
	// os.Stdin = in
	// defer in.Close()

	for {
		fmt.Scan(&n, &m)
		if n|m == 0 {
			break
		}

		var f [N][M]int64

		for i := 0; i < 1<<n; i++ {
			cnt := 0
			isValid := true
			for j := 0; j < n; j++ {
				if i>>j&1 == 1 {
					if cnt&1 == 1 {
						isValid = false
						break
					}
					cnt = 0
				} else {
					cnt++
				}
			}
			if cnt&1 == 1 {
				isValid = false
			}
			st[i] = isValid
		}

		for i := 0; i < 1<<n; i++ {
			state[i] = []int{}
			for j := 0; j < 1<<n; j++ {
				if i&j == 0 && st[i|j] {
					state[i] = append(state[i], j)
				}
			}
		}

		f[0][0] = 1

		for i := 1; i <= m; i++ {
			for j := 0; j < 1<<n; j++ {
				for _, k := range state[j] {
					f[i][j] += f[i-1][k]
				}
			}
		}

		fmt.Println(f[m][0])
	}
}
