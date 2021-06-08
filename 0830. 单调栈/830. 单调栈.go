package main

import "fmt"

const N = 100010

var stack [N]int
var tt int

func main() {
	var n int
	fmt.Scan(&n)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		for tt > 0 && stack[tt] >= x {
			tt--
		}
		if tt == 0 {
			res[i] = -1
		} else {
			res[i] = stack[tt]
		}
		tt++
		stack[tt] = x
	}

	for _, v := range res {
		fmt.Printf("%d ", v)
	}

}
