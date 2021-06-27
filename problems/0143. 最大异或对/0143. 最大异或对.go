package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const N = 100010

var (
	son [N * 31][2]int
	idx int

	scanner *bufio.Scanner
)

func insert(num int) {
	p := 0
	for i := 30; i >= 0; i-- {
		u := num >> i & 1
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
}

func query(num int) int {
	var res int
	p := 0
	for i := 30; i >= 0; i-- {
		u := num >> i & 1
		notu := u ^ 1
		if son[p][notu] != 0 {
			res += 1 << i
			p = son[p][notu]
			// } else if son[p][u] != 0 {
			// 	p = son[p][u]
			// } else {
			// 	break
		} else {
			p = son[p][u]
			fmt.Println(p)
		}
	}
	return res
}

func main() {
	var n int
	var max int
	fmt.Scan(&n)
	newScanner(os.Stdin)
	x := readLine()
	for _, v := range x {
		insert(v)
	}
	for _, v := range x {
		tmp := query(v)
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(max)
}

func newScanner(reader io.Reader) {
	scanner = bufio.NewScanner(reader)
	bs := make([]byte, 20000*1024)
	scanner.Buffer(bs, len(bs))
}

func readLine() []int {
	scanner.Scan()
	line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	res := make([]int, len(line))
	for i, l := range line {
		x, _ := strconv.Atoi(l)
		res[i] = x
	}
	return res
}
