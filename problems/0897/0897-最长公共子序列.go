package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-rythm/acwing/internal/oj"
)

const (
	N = 1010
)

var (
	r    = bufio.NewReader(os.Stdin)
	n, m int
	a, b string
	f    [N][N]int
)

func main() {
	fmt.Fscan(r, &n, &m, &a, &b)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = oj.Max(f[i-1][j], f[i][j-1])
			}
		}
	}

	fmt.Println(f[n][m])
}
