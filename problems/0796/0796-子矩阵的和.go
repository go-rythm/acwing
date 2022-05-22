package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-rythm/acwing/internal/oj"
)

const N = 1010

var a, s [N][N]int
var ot = bufio.NewWriterSize(os.Stdout, int(1e6))

func main() {
	defer ot.Flush()
	in := oj.NewScanner()
	n, m, q := in.ReadInt(), in.ReadInt(), in.ReadInt()
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] = in.ReadInt()
			s[i][j] = s[i][j-1] + s[i-1][j] - s[i-1][j-1] + a[i][j]
		}
	}
	for i := 0; i < q; i++ {
		r := in.ReadInts(4)
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		res := s[x2][y2] - s[x2][y1-1] - s[x1-1][y2] + s[x1-1][y1-1]
		fmt.Fprintln(ot, res)
	}
}
