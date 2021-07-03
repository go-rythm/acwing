package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	sc *bufio.Scanner
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
	newScanner(os.Stdin)
	a := readLine()
	n, m = a[0], a[1]
	for i := 0; i < n; i++ {
		a = readLine()
		for j := 0; j < m; j++ {
			g[i][j] = a[j]
		}
	}
	dis := bfs(&point{0, 0})
	fmt.Println(dis)
}

func memset(arr *[N][N]int, value int) {
	// for i, v := range arr {
	// 	for j := range v {
	// 		arr[i][j] = value
	// 	}
	// }
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			arr[i][j] = value
		}
	}
}

func newScanner(reader io.Reader) {
	sc = bufio.NewScanner(reader)
	bs := make([]byte, 20000*1024)
	sc.Buffer(bs, len(bs))
}

func readLine() []int {
	sc.Scan()
	line := strings.Split(strings.TrimSpace(sc.Text()), " ")
	res := make([]int, len(line))
	for i, l := range line {
		x, _ := strconv.Atoi(l)
		res[i] = x
	}
	return res
}
