package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := NewScanner()
	n, m := s.ReadInt(), s.ReadInt()
	q := s.ReadInts(n)
	for i := 0; i < m; i++ {
		x := s.ReadInt()

		var leftBound int
		var rightBound int
		l := 0
		r := n - 1
		for l+1 < r {
			mid := (l + r) >> 1
			if q[mid] >= x {
				r = mid
			} else {
				l = mid
			}
		}
		if q[l] == x {
			leftBound = l
		} else if q[r] == x {
			leftBound = r
		} else {
			fmt.Println("-1 -1")
			continue
		}

		l = 0
		r = n - 1
		for l+1 < r {
			mid := (l + r) >> 1
			if q[mid] > x {
				r = mid
			} else {
				l = mid
			}
		}
		if q[r] == x {
			rightBound = r
		} else if q[l] == x {
			rightBound = l
		} else {
			fmt.Println("-1 -1")
			continue
		}

		fmt.Printf("%d %d\n", leftBound, rightBound)
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
