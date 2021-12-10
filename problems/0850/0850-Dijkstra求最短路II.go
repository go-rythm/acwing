package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	N   = 150010
	inf = 0x3f3f3f3f
)

var (
	h, e, ne, w, dist [N]int
	idx               int
	st                [N]bool
	n, m              int

	scanner *bufio.Scanner
)

type pair struct {
	vertex int
	dist   int
}

func add(x, y, c int) {
	w[idx] = c
	e[idx] = y
	ne[idx] = h[x]
	h[x] = idx
	idx++
}

func dijkstra() int {
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[1] = 0
	pq := PriorityQueue{
		&pair{
			vertex: 1,
			dist:   0,
		}}
	heap.Init(&pq)

	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*pair)
		if st[p.vertex] {
			continue
		}
		st[p.vertex] = true
		for i := h[p.vertex]; i != -1; i = ne[i] {
			v := e[i]
			if dist[v] > p.dist+w[i] {
				dist[v] = p.dist + w[i]
				heap.Push(&pq, &pair{
					vertex: v,
					dist:   dist[v],
				})
			}
		}
	}
	if dist[n] == inf {
		return -1
	}
	return dist[n]
}

func main() {
	newScanner(os.Stdin)
	for i := 0; i < len(h); i++ {
		h[i] = -1
	}
	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {
		a := readLine()
		add(a[0], a[1], a[2])
	}
	fmt.Println(dijkstra())
}

type PriorityQueue []*pair

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	if p[i].dist == p[j].dist {
		return p[i].vertex < p[j].vertex
	}
	return p[i].dist < p[j].dist
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(*pair))
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func newScanner(reader io.Reader) {
	scanner = bufio.NewScanner(reader)
	bs := make([]byte, 20000*1024)
	scanner.Buffer(bs, len(bs))
}

func readLine() []int {
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	res := make([]int, len(line))
	for i, l := range line {
		x, _ := strconv.Atoi(l)
		res[i] = x
	}
	return res
}
