package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 310

var (
	n, m int
	h, f [N][N]int
	dx   = [4]int{-1, 0, 1, 0}
	dy   = [4]int{0, 1, 0, -1}
)

func dp(x, y int) int {
	if f[x][y] != -1 {
		return f[x][y]
	}
	f[x][y] = 1
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 1 && a <= n && b >= 1 && b <= m && h[a][b] < h[x][y] {
			f[x][y] = Max(f[x][y], dp(a, b)+1)
		}
	}
	return f[x][y]
}

func main() {
	s := NewScanner()
	n, m = s.ReadInt(), s.ReadInt()
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			h[i][j] = s.ReadInt()
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = -1
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			res = Max(res, dp(i, j))
		}
	}

	fmt.Println(res)
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
