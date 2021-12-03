package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 110

var (
	n, m int
	// 存储地图
	g [N][N]int
	// 存储每个点到起点的距离
	d [N][N]int
	// 模拟队列
	q [N * N]*point
)

type point struct {
	x, y int
}

func bfs(start *point) int {
	hh, tt := 0, 0
	q[0] = start

	// 初始化，置所有距离为-1，起点距离为0
	memset(&d, -1)
	d[0][0] = 0

	// 四个方向向量，表示一个点下一个走向
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, -1, 0, 1}

	for hh <= tt {
		p := q[hh]
		hh++
		for i := 0; i < 4; i++ {
			x := p.x + dx[i]
			y := p.y + dy[i]
			// 点的合法性校验
			if x >= 0 && x < n && y >= 0 && y < m && g[x][y] == 0 && d[x][y] == -1 {
				d[x][y] = d[p.x][p.y] + 1
				tt++
				q[tt] = &point{x, y}
			}
		}
	}

	return d[n-1][m-1]
}

func main() {
	s := NewScanner()
	n, m = s.ReadInt(), s.ReadInt()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			g[i][j] = s.ReadInt()
		}
	}
	dis := bfs(&point{0, 0})
	fmt.Println(dis)
}

func memset(arr *[N][N]int, value int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			arr[i][j] = value
		}
	}
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
