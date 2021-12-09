package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N   = 510
	inf = 0x3f3f3f3f
)

var (
	g    [N][N]int
	dist [N]int
	st   [N]bool

	n, m int
)

func dijkstra() int {
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[1] = 0

	for i := 0; i < n; i++ {
		t := -1

		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		st[t] = true

		for j := 1; j <= n; j++ {
			dist[j] = Min(dist[j], dist[t]+g[t][j])
		}
	}

	if dist[n] == inf {
		return -1
	}
	return dist[n]
}

func main() {
	s := NewScanner()
	n, m = s.ReadInt(), s.ReadInt()
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			g[i][j] = inf
		}
	}

	for i := 0; i < m; i++ {
		x, y, z := s.ReadInt(), s.ReadInt(), s.ReadInt()
		g[x][y] = Min(g[x][y], z)
	}

	fmt.Println(dijkstra())
}

/*
 *        __         __
 *       / /_  ___  / /___  ___  __________
 *      / __ \/ _ \/ / __ \/ _ \/ ___/ ___/
 *     / / / /  __/ / /_/ /  __/ /  (__  )
 *    /_/ /_/\___/_/ .___/\___/_/  /____/
 *                /_/
 */

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

// Swap 交换两个int指针指向的值
func Swap(a, b *int) {
	*a, *b = *b, *a
}

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
