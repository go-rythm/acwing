package main

import "fmt"

func divide(x int) {
	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			c := 0
			for ; x%i == 0; x /= i {
				c++
			}
			fmt.Printf("%d %d\n", i, c)
		}
	}
	if x > 1 {
		fmt.Printf("%d %d\n", x, 1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		divide(x)
		fmt.Println()
	}
}
