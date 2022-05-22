package main

import (
	"fmt"

	"github.com/go-rythm/acwing/internal/oj"
)

const (
	N = 20
	M = 1 << N
)

var (
	f [M][N]int
	w [N][N]int
)

func main() {
	// in := oj.In()
	// os.Stdin = in
	// defer in.Close()

	s := oj.NewScanner()
	n := s.ReadInt()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			w[i][j] = s.ReadInt()
		}
	}
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[0]); j++ {
			f[i][j] = int(1e9)
		}
	}
	f[1][0] = 0
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				for k := 0; k < n; k++ {
					if i>>k&1 == 1 {
						f[i][j] = oj.Min(f[i][j], f[i-(1<<j)][k]+w[k][j])
					}
				}
			}
		}
	}

	fmt.Println(f[(1<<n)-1][n-1])
}
