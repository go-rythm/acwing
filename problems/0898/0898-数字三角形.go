package main

import (
	"fmt"
)

const N = 510

func main() {

	var (
		n    int
		a, f [N][N]int
	)

	// in, err := os.Open("in.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer in.Close()
	// os.Stdin = in

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		f[n][i] = a[n][i]
	}

	for i := n - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			f[i][j] = Max(f[i+1][j], f[i+1][j+1]) + a[i][j]
		}
	}

	fmt.Println(f[1][1])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
