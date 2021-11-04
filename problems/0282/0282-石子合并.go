package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
)

const (
	N   = 10
	inf = 100000000
)

func main() {
	var (
		n int
		s [N]int
		f [N][N]int
	)

	// in, err := os.Open("in.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer in.Close()
	// os.Stdin = in

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
		s[i] += s[i-1]
	}

	for len := 2; len <= n; len++ {
		for l := 1; l+len-1 <= n; l++ {
			r := l + len - 1
			f[l][r] = inf
			w := s[r] - s[l-1]
			for k := 0; k < r-l; k++ {
				f[l][r] = oj.Min(f[l][r], f[l][l+k]+f[l+k+1][r]+w)
			}
		}
	}
	fmt.Println(f[1][n])
}
