package main

import "fmt"

func lowbit(x uint) uint {
	return x & -x
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x uint
		fmt.Scan(&x)
		var count int
		for x != 0 {
			x -= lowbit(x)
			count++
		}
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(count)
	}
}
