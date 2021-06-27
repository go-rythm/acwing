package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const N = 1000010

var (
	// a 保存输入
	// q 维护一个单调队列，保存着数组下标
	a, q [N]int
	head = 0
	tail = -1
	n    int
	k    int

	scanner *bufio.Scanner
)

func main() {
	fmt.Scan(&n, &k)
	newScanner(os.Stdin)
	x := readLine()
	for i := 0; i < n; i++ {
		a[i] = x[i]
	}

	for i := 0; i < n; i++ {
		// 队列非空且单调队列中第一个下标值已经超出了当前的窗口
		// 因为每次只移动一格，所以一条 if 语句就足够了，不需要使用循环
		if head <= tail && q[head] < i-k+1 {
			head++
		}
		// 队列非空且出现了更小的值
		for head <= tail && a[i] < a[q[tail]] {
			tail--
		}
		tail++
		q[tail] = i
		if i >= k-1 {
			fmt.Print(a[q[head]], " ")
		}
	}

	fmt.Println()
	head = 0
	tail = -1
	for i := 0; i < n; i++ {
		if head <= tail && q[head] < i-k+1 {
			head++
		}
		for head <= tail && a[i] > a[q[tail]] {
			tail--
		}
		tail++
		q[tail] = i
		if i >= k-1 {
			fmt.Print(a[q[head]], " ")
		}
	}
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
