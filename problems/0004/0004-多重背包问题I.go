package main

import (
	"fmt"
	"os"

	"github.com/luxcgo/go-acwing/internal/oj"
)

func main() {
	var (
		n, m    int
		v, w, s []int
		f       [][]int
	)

	in, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	os.Stdin = in

	fmt.Scan(&n, &m)
	v = make([]int, n+1)
	w = make([]int, n+1)
	s = make([]int, n+1)
	f = make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		fmt.Scan(&v[i], &w[i], &s[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= s[i] && k*v[i] <= j; k++ {
				f[i][j] = oj.Max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}
	fmt.Println(f[n][m])
}
