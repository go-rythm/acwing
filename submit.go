package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const N = 3e5 + 10

var a, s [N]int

type pair struct { // 定义操作
	p1 int
	p2 int
}

func main() {
	in := NewScanner()
	n, m := in.ReadInt(), in.ReadInt()
	var alls []int
	add := make([]pair, 0)
	for i := 0; i < n; i++ {
		x, c := in.ReadInt(), in.ReadInt()
		add = append(add, pair{x, c})
		alls = append(alls, x)
	}

	query := make([]pair, 0)
	for i := 0; i < m; i++ {
		l, r := in.ReadInt(), in.ReadInt()
		query = append(query, pair{l, r})
		alls = append(alls, l)
		alls = append(alls, r)
	}

	sort.Ints(alls)
	alls = Unique(alls)

	for _, v := range add {
		a[findIndex(alls, v.p1)] += v.p2
	}

	for i := 0; i < len(alls); i++ {
		s[i+1] = s[i] + a[i]
	}

	for _, v := range query {
		l := findIndex(alls, v.p1)
		r := findIndex(alls, v.p2)
		fmt.Println(s[r+1] - s[l])
	}
}

func findIndex(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l+1 < r {
		mid := (l + r) >> 1
		if arr[mid] < target {
			l = mid
		} else if arr[mid] > target {
			r = mid
		} else {
			return mid
		}
	}
	if arr[l] == target {
		return l
	}
	if arr[r] == target {
		return r
	}
	return -1
}

/*
 *        __         __
 *       / /_  ___  / /___  ___  __________
 *      / __ \/ _ \/ / __ \/ _ \/ ___/ ___/
 *     / / / /  __/ / /_/ /  __/ /  (__  )
 *    /_/ /_/\___/_/ .___/\___/_/  /____/
 *                /_/
 */

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func In() *os.File {
	f, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	return f
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := &Scanner{
		bufio.NewScanner(os.Stdin),
	}
	s.Split(bufio.ScanWords)
	s.Buffer(make([]byte, 1024), int(1e9))
	return s
}

func (s *Scanner) readString() string {
	s.Scan()
	return s.Text()
}

func (s *Scanner) readInt() int {
	i, err := strconv.Atoi(s.readString())
	if err != nil {
		panic(err)
	}
	return i
}

func (s *Scanner) ReadString() string {
	return s.readString()
}

func (s *Scanner) ReadInt() int {
	return s.readInt()
}

func (s *Scanner) ReadStrings(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = s.readString()
	}
	return a
}

func (s *Scanner) ReadInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = s.readInt()
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Memset(a []int, v int) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}

func Unique(a []int) []int {
	j := 0
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] != a[i-1] {
			a[j] = a[i]
			j++
		}
	}
	return a[:j]
}
