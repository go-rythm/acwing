## [847\. 图中点的层次](https://www.acwing.com/problem/content/849/)

### 题目

给定一个 $n$ 个点 $m$ 条边的有向图，图中可能存在重边和自环。

所有边的长度都是 $1$，点的编号为 $1 \\sim n$。

请你求出 $1$ 号点到 $n$ 号点的最短距离，如果从 $1$ 号点无法走到 $n$ 号点，输出 $-1$。

### 输入格式

第一行包含两个整数 $n$ 和 $m$。

接下来 $m$ 行，每行包含两个整数 $a$ 和 $b$，表示存在一条从 $a$ 走到 $b$ 的长度为 $1$ 的边。

### 输出格式

输出一个整数，表示 $1$ 号点到 $n$ 号点的最短距离。

### 数据范围

$1 \\le n,m \\le 10^5$

### 输入样例：

```
4 5
1 2
2 3
3 4
1 3
1 4
```

### 输出样例：

```
1
```

### 题解

前置题目：0846

前置知识：队列

本题知识：搜索与图论-BFS

#### 题目分析

树的宽度优先遍历

```go
const N = 100010

var (
	h, e, ne, d, q [N]int
    st [N]bool
	idx            int
)

// 添加一条边a->b
func add(a int, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

// 初始化
func init() {
    idx = 0;
    for i := 0; i < len(h); i++ {
		h[i] = -1
	}
}

func bfs() {
    hh, tt := 0, 0
    q[0] = 1
    for hh <= tt {
		t := q[hh]
		hh++
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if !st[j] {
				st[j] = true
				tt++
				q[tt] = j
			}
		}
	}
}
```

