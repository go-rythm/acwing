package main

import "fmt"

const N = 1000010

var (
	prime [N]int
	st    [N]bool
	cnt   int
	x     int
)

func getPrimes(x int) {
	for i := 2; i <= x; i++ {
		if !st[i] {
			prime[cnt] = i
			cnt++
		}
		for j := 0; prime[j] <= x/i; j++ {
			st[prime[j]*i] = true
			if i%prime[j] == 0 {
				break
			}
		}
	}
}

func main() {
	fmt.Scan(&x)
	getPrimes(x)
	fmt.Println(cnt)
}
