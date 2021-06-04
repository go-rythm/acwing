package main

import "fmt"

const N = 100010

var (
	n int
	a [N]int
	s [N]int
)

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var max int
	for i, j := 0, 0; i < n; i++ {
		s[a[i]]++
		for s[a[i]] > 1 {
			s[a[j]]--
			j++
		}
		if i-j+1 > max {
			max = i - j + 1
		}
	}
	fmt.Println(max)
	for i, j := 0, 0; i < n; i++ {
		for j < i && check(i, j) {
			j++
		}
	}
}
