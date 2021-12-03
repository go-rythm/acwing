package main

import (
	"fmt"

	"github.com/luxcgo/go-acwing/internal/oj"
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
	s := oj.NewScanner()
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
